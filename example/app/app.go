// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/ChainSafe/chainbridge-core/topology"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/ChainSafe/chainbridge-core/chains/evm"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/contracts/bridge"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/events"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/evmclient"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/evmtransaction"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor/signAndSend"
	"github.com/ChainSafe/chainbridge-core/chains/evm/executor"
	"github.com/ChainSafe/chainbridge-core/chains/evm/listener"
	"github.com/ChainSafe/chainbridge-core/comm/elector"
	"github.com/ChainSafe/chainbridge-core/comm/p2p"
	"github.com/ChainSafe/chainbridge-core/config"
	"github.com/ChainSafe/chainbridge-core/config/chain"
	"github.com/ChainSafe/chainbridge-core/e2e/dummy"
	"github.com/ChainSafe/chainbridge-core/flags"
	"github.com/ChainSafe/chainbridge-core/lvldb"
	"github.com/ChainSafe/chainbridge-core/opentelemetry"
	"github.com/ChainSafe/chainbridge-core/relayer"
	"github.com/ChainSafe/chainbridge-core/store"
	"github.com/ChainSafe/chainbridge-core/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() error {
	configuration, err := config.GetConfig(viper.GetString(flags.ConfigFlagName))
	if err != nil {
		panic(err)
	}

	networkTopology, _ := topology.ProcessRawTopology(&topology.RawTopology{
		Peers: []topology.RawPeer{
			{PeerAddress: "/dns4/relayer2/tcp/9001/p2p/QmeTuMtdpPB7zKDgmobEwSvxodrf5aFVSmBXX3SQJVjJaT"},
			{PeerAddress: "/dns4/relayer3/tcp/9002/p2p/QmYAYuLUPNwYEBYJaKHcE7NKjUhiUV8txx2xDXHvcYa1xK"},
			{PeerAddress: "/dns4/relayer1/tcp/9000/p2p/QmcvEg7jGvuxdsUFRUiE4VdrL2P1Yeju5L83BsJvvXz7zX"},
		},
	})

	var allowedPeers peer.IDSlice
	for _, pAdrInfo := range networkTopology.Peers {
		allowedPeers = append(allowedPeers, pAdrInfo.ID)
	}

	db, err := lvldb.NewLvlDB(viper.GetString(flags.BlockstoreFlagName))
	if err != nil {
		panic(err)
	}
	blockstore := store.NewBlockStore(db)

	privBytes, err := ioutil.ReadFile(configuration.RelayerConfig.MpcConfig.KeystorePath)
	if err != nil {
		panic(err)
	}
	priv, err := crypto.UnmarshalPrivateKey(privBytes)
	if err != nil {
		panic(err)
	}
	host, err := p2p.NewHost(priv, networkTopology, configuration.RelayerConfig.MpcConfig.Port)
	if err != nil {
		panic(err)
	}
	comm := p2p.NewCommunication(host, "p2p/chainbridge", allowedPeers)
	electorFactory := elector.NewCoordinatorElectorFactory(host, configuration.RelayerConfig.BullyConfig)
	coordinator := tss.NewCoordinator(host, comm, electorFactory)
	keyshareStore := store.NewKeyshareStore(configuration.RelayerConfig.MpcConfig.KeysharePath)

	chains := []relayer.RelayedChain{}
	for _, chainConfig := range configuration.ChainConfigs {
		switch chainConfig["type"] {
		case "evm":
			{
				config, err := chain.NewEVMConfig(chainConfig)
				if err != nil {
					panic(err)
				}

				client, err := evmclient.NewEVMClient(config)
				if err != nil {
					panic(err)
				}

				bridgeAddress := common.HexToAddress(config.Bridge)
				dummyGasPricer := dummy.NewStaticGasPriceDeterminant(client, nil)
				t := signAndSend.NewSignAndSendTransactor(evmtransaction.NewTransaction, dummyGasPricer, client)
				bridgeContract := bridge.NewBridgeContract(client, bridgeAddress, t)

				depositHandler := listener.NewETHDepositHandler(*bridgeContract)
				depositHandler.RegisterDepositHandler(config.Erc20Handler, listener.Erc20DepositHandler)
				depositHandler.RegisterDepositHandler(config.Erc721Handler, listener.Erc721DepositHandler)
				depositHandler.RegisterDepositHandler(config.GenericHandler, listener.GenericDepositHandler)
				eventListener := events.NewListener(client)
				eventHandlers := make([]listener.EventHandler, 0)
				eventHandlers = append(eventHandlers, listener.NewDepositEventHandler(eventListener, depositHandler, bridgeAddress, *config.GeneralChainConfig.Id))
				eventHandlers = append(eventHandlers, listener.NewKeygenEventHandler(eventListener, coordinator, host, comm, keyshareStore, bridgeAddress, configuration.RelayerConfig.MpcConfig.Threshold))
				eventHandlers = append(eventHandlers, listener.NewRefreshEventHandler(nil, eventListener, coordinator, host, comm, keyshareStore, bridgeAddress, configuration.RelayerConfig.MpcConfig.Threshold))
				evmListener := listener.NewEVMListener(client, eventHandlers, blockstore, config)

				mh := executor.NewEVMMessageHandler(*bridgeContract)
				mh.RegisterMessageHandler(config.Erc20Handler, executor.ERC20MessageHandler)
				mh.RegisterMessageHandler(config.Erc721Handler, executor.ERC721MessageHandler)
				mh.RegisterMessageHandler(config.GenericHandler, executor.GenericMessageHandler)
				executor := executor.NewExecutor(host, comm, coordinator, mh, bridgeContract, keyshareStore)

				chain := evm.NewEVMChain(evmListener, executor, blockstore, config)

				chains = append(chains, chain)
			}
		default:
			panic(fmt.Errorf("type '%s' not recognized", chainConfig["type"]))
		}
	}

	r := relayer.NewRelayer(
		chains,
		&opentelemetry.ConsoleTelemetry{},
	)

	errChn := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go r.Start(ctx, errChn)

	sysErr := make(chan os.Signal, 1)
	signal.Notify(sysErr,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	select {
	case err := <-errChn:
		log.Error().Err(err).Msg("failed to listen and serve")
		return err
	case sig := <-sysErr:
		log.Info().Msgf("terminating got ` [%v] signal", sig)
		return nil
	}
}
