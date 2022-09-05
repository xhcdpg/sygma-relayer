// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package accessControlSegregator

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/contracts"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"

	"github.com/ChainSafe/sygma-relayer/chains/evm/calls/consts"
)

type AccessControlSegregatorContract struct {
	contracts.Contract
}

func NewAccessControlSegregatorContract(
	client calls.ContractCallerDispatcher,
	address common.Address,
	transactor transactor.Transactor,
) *AccessControlSegregatorContract {
	a, _ := abi.JSON(strings.NewReader(consts.AccessControlSegregatorABI))
	b := common.FromHex(consts.AccessControlSegregatorBin)
	return &AccessControlSegregatorContract{contracts.NewContract(address, a, b, client, transactor)}
}
