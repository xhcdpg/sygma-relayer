// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package consts

const BridgeABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"accessControl\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAccessControl\",\"type\":\"address\"}],\"name\":\"AccessControlChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"handlerResponse\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EndKeygen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeHandler\",\"type\":\"address\"}],\"name\":\"FeeHandlerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"KeyRefresh\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"Retry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StartKeygen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_MPCAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_accessControl\",\"outputs\":[{\"internalType\":\"contractIAccessControlSegregator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_domainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeHandler\",\"outputs\":[{\"internalType\":\"contractIFeeHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositorOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"adminSetDepositNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"adminSetForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAccessControl\",\"type\":\"address\"}],\"name\":\"adminChangeAccessControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeHandler\",\"type\":\"address\"}],\"name\":\"adminChangeFeeHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBridge.Proposal\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBridge.Proposal[]\",\"name\":\"proposals\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"executeProposals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startKeygen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"MPCAddress\",\"type\":\"address\"}],\"name\":\"endKeygen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"refreshKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"retry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"isProposalExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBridge.Proposal[]\",\"name\":\"proposals\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
const BridgeBin = "0x6101606040523480156200001257600080fd5b506040516200301e3803806200301e833981016040819052620000359162000253565b604080518082018252600681526542726964676560d01b6020808301918252835180850190945260058452640332e312e360dc1b908401526000805460ff191690558151902060e08190527f0e23d0b508e2034d01a5c31f12e9d9bbb31708c5518057dde31201ab93b17cef6101008190524660a0529192917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6200011f8184846040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b6080523060c052610120525050505060ff821661014052600280546001600160a01b0319166001600160a01b038316179055620001656200015f6200016d565b620001b0565b5050620002a2565b600033601436108015906200019a57506001600160a01b03811660009081526005602052604090205460ff165b15620001ab575060131936013560601c5b919050565b620001ba62000206565b6000805460ff191660011790556040516001600160a01b03821681527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589060200160405180910390a150565b60005460ff1615620002515760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640160405180910390fd5b565b600080604083850312156200026757600080fd5b825160ff811681146200027957600080fd5b60208401519092506001600160a01b03811681146200029757600080fd5b809150509250929050565b60805160a05160c05160e051610100516101205161014051612d136200030b6000396000818161043d01528181610caa0152610dba01526000611ea301526000611ef201526000611ecd01526000611e2601526000611e5001526000611e7a0152612d136000f3fe6080604052600436106101b75760003560e01c80639ae0bf45116100ec578063d2e5fae91161008a578063f0ead51e11610064578063f0ead51e14610551578063f8c39e4414610571578063fe4648f4146105a1578063ffaac0eb146105c157600080fd5b8063d2e5fae9146104f1578063d823674414610511578063edc20c3c1461053157600080fd5b8063a546e8a1116100c6578063a546e8a114610471578063bd2a182014610491578063cb10f215146104b1578063d15ef64e146104d157600080fd5b80639ae0bf45146103eb5780639d33b6d41461040b5780639dd694f41461042b57600080fd5b80635c975abb1161015957806380ae1c281161013357806380ae1c281461036057806384db809f146103755780638b63aebf146103ab5780638c0c2631146103cb57600080fd5b80635c975abb146103145780636ba6db6b1461033857806373c45c981461034d57600080fd5b8063366b488511610195578063366b48851461026657806344e8e430146102865780634b0b919d146102a65780635a1ad87c146102f457600080fd5b8063059972d2146101bc57806308a64104146101fe5780631f5c64c114610244575b600080fd5b3480156101c857600080fd5b506000546101e19061010090046001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561020a57600080fd5b506102366102193660046122ee565b600660209081526000928352604080842090915290825290205481565b6040519081526020016101f5565b34801561025057600080fd5b5061026461025f3660046124d2565b6105d6565b005b34801561027257600080fd5b506102646102813660046125bc565b610afa565b34801561029257600080fd5b506002546101e1906001600160a01b031681565b3480156102b257600080fd5b506102dc6102c136600461260c565b6003602052600090815260409020546001600160401b031681565b6040516001600160401b0390911681526020016101f5565b34801561030057600080fd5b5061026461030f366004612656565b610b34565b34801561032057600080fd5b5060005460ff165b60405190151581526020016101f5565b34801561034457600080fd5b50610264610bff565b61026461035b3660046126bc565b610ca0565b34801561036c57600080fd5b50610264610fad565b34801561038157600080fd5b506101e1610390366004612745565b6004602052600090815260409020546001600160a01b031681565b3480156103b757600080fd5b506102646103c636600461275e565b610fd7565b3480156103d757600080fd5b506102646103e6366004612779565b61103d565b3480156103f757600080fd5b506103286104063660046122ee565b6110b9565b34801561041757600080fd5b5061026461042636600461275e565b611108565b34801561043757600080fd5b5061045f7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f5565b34801561047d57600080fd5b5061032861048c3660046124d2565b61116e565b34801561049d57600080fd5b506102646104ac3660046127ac565b6113c9565b3480156104bd57600080fd5b506102646104cc3660046127f9565b61140f565b3480156104dd57600080fd5b506102646104ec366004612843565b6114b3565b3480156104fd57600080fd5b5061026461050c36600461275e565b6114f6565b34801561051d57600080fd5b5061026461052c3660046125bc565b611628565b34801561053d57600080fd5b5061026461054c36600461287a565b61166f565b34801561055d57600080fd5b5061026461056c3660046128a4565b611738565b34801561057d57600080fd5b5061032861058c36600461275e565b60056020526000908152604090205460ff1681565b3480156105ad57600080fd5b506001546101e1906001600160a01b031681565b3480156105cd57600080fd5b50610264611b2e565b6105de611ba4565b600083511161063e5760405162461bcd60e51b815260206004820152602160248201527f50726f706f73616c732063616e277420626520616e20656d70747920617272616044820152607960f81b60648201526084015b60405180910390fd5b61064983838361116e565b61068f5760405162461bcd60e51b815260206004820152601760248201527624b73b30b634b210383937b837b9b0b61039b4b3b732b960491b6044820152606401610635565b60005b8351811015610af4576106e88482815181106106b0576106b06128f2565b6020026020010151600001518583815181106106ce576106ce6128f2565b6020026020010151602001516001600160401b03166110b9565b156106f257610ae2565b60006004600086848151811061070a5761070a6128f2565b602002602001015160400151815260200190815260200160002060009054906101000a90046001600160a01b0316905060008186848151811061074f5761074f6128f2565b60200260200101516060015160405160200161076c929190612934565b604051602081830303815290604052805190602001209050600082905061010087858151811061079e5761079e6128f2565b6020026020010151602001516107b49190612982565b6001600160401b03166001901b600660008987815181106107d7576107d76128f2565b60200260200101516000015160ff1660ff16815260200190815260200160002060006101008a888151811061080e5761080e6128f2565b60200260200101516020015161082491906129be565b6001600160401b0316815260200190815260200160002060008282541792505081905550806001600160a01b031663e248cff2888681518110610869576108696128f2565b602002602001015160400151898781518110610887576108876128f2565b6020026020010151606001516040518363ffffffff1660e01b81526004016108b0929190612a10565b600060405180830381600087803b1580156108ca57600080fd5b505af19250505080156108db575060015b610a4a573d808015610909576040519150601f19603f3d011682016040523d82523d6000602084013e61090e565b606091505b507f19f774a63ee465292252a9981ae52051acc13da671c698ac4b5bf25b38c5b6fc81898781518110610943576109436128f2565b6020026020010151600001518a8881518110610961576109616128f2565b60200260200101516020015160405161097c93929190612a29565b60405180910390a1610100888681518110610999576109996128f2565b6020026020010151602001516109af9190612982565b6001600160401b03166001901b19600660008a88815181106109d3576109d36128f2565b60200260200101516000015160ff1660ff16815260200190815260200160002060006101008b8981518110610a0a57610a0a6128f2565b602002602001015160200151610a2091906129be565b6001600160401b0316815260208101919091526040016000208054909116905550610ae292505050565b7f6018c584b8d99bafeda249b2429f5907d830e792222070c1b3a94aa76ee71677878581518110610a7d57610a7d6128f2565b602002602001015160000151888681518110610a9b57610a9b6128f2565b60200260200101516020015184604051610ad69392919060ff9390931683526001600160401b03919091166020830152604082015260600190565b60405180910390a15050505b80610aec81612a5e565b915050610692565b50505050565b7f9069464c059b9a90135a3fdf2c47855263346b912894ad7562d989532c3fad4c81604051610b299190612a79565b60405180910390a150565b610b516000356001600160e01b031916610b4c611bea565b611c2b565b60008581526004602081905260409182902080546001600160a01b0319166001600160a01b038a8116918217909255925163de319d9960e01b8152918201889052861660248201526001600160e01b03198086166044830152606482018590528316608482015287919063de319d999060a401600060405180830381600087803b158015610bde57600080fd5b505af1158015610bf2573d6000803e3d6000fd5b5050505050505050505050565b610c176000356001600160e01b031916610b4c611bea565b60005461010090046001600160a01b031615610c755760405162461bcd60e51b815260206004820152601a60248201527f4d5043206164647265737320697320616c7265616479207365740000000000006044820152606401610635565b6040517f24e723a5c27b62883404028b8dee9965934de6a46828cda2ff63bf9a5e65ce4390600090a1565b610ca8611ba4565b7f000000000000000000000000000000000000000000000000000000000000000060ff168660ff161415610d1e5760405162461bcd60e51b815260206004820152601f60248201527f43616e2774206465706f73697420746f2063757272656e7420646f6d61696e006044820152606401610635565b6000610d28611bea565b6001549091506001600160a01b0316610d8e573415610d895760405162461bcd60e51b815260206004820152601d60248201527f6e6f2046656548616e646c65722c206d73672e76616c756520213d20300000006044820152606401610635565b610e21565b600154604051632530706560e01b81526001600160a01b03909116906325307065903490610dee9085907f0000000000000000000000000000000000000000000000000000000000000000908d908d908d908d908d908d90600401612ab5565b6000604051808303818588803b158015610e0757600080fd5b505af1158015610e1b573d6000803e3d6000fd5b50505050505b6000868152600460205260409020546001600160a01b031680610e865760405162461bcd60e51b815260206004820181905260248201527f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726044820152606401610635565b60ff8816600090815260036020526040812080548290610eae906001600160401b0316612b11565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060008290506000816001600160a01b031663b07e54bb8b878c8c6040518563ffffffff1660e01b8152600401610f0c9493929190612b38565b6000604051808303816000875af1158015610f2b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f539190810190612b6d565b9050846001600160a01b03167f17bc3181e17a9620a479c24e6c606e474ba84fc036877b768926872e8cd0e11f8c8c868d8d87604051610f9896959493929190612be3565b60405180910390a25050505050505050505050565b610fc56000356001600160e01b031916610b4c611bea565b610fd5610fd0611bea565b611d08565b565b610fef6000356001600160e01b031916610b4c611bea565b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f729170bd142e4965055b26a285faeedf03baf2b915bfc5a7c75d24b45815ff2c90602001610b29565b6110556000356001600160e01b031916610b4c611bea565b6040516307b7ed9960e01b81526001600160a01b0382811660048301528391908216906307b7ed99906024015b600060405180830381600087803b15801561109c57600080fd5b505af11580156110b0573d6000803e3d6000fd5b50505050505050565b60006110c761010083612c34565b60ff84166000908152600660205260408120600190921b91906110ec61010086612c48565b8152602001908152602001600020541660001415905092915050565b6111206000356001600160e01b031916610b4c611bea565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527f497acaa34ac19c2a2a579ad43eca493b4fea820459e254e9383e7dda216b0f0490602001610b29565b60008084516001600160401b0381111561118a5761118a612318565b6040519080825280602002602001820160405280156111b3578160200160208202803683370190505b50905060005b85518110156112e5577fcc13634e956dd3d4ec8d808ee8bf294e1cd05a38f63fe7f234b079a0a4c36a708682815181106111f5576111f56128f2565b602002602001015160000151878381518110611213576112136128f2565b602002602001015160200151888481518110611231576112316128f2565b60200260200101516040015189858151811061124f5761124f6128f2565b602002602001015160600151805190602001206040516020016112a095949392919094855260ff9390931660208501526001600160401b039190911660408401526060830152608082015260a00190565b604051602081830303815290604052805190602001208282815181106112c8576112c86128f2565b6020908102919091010152806112dd81612a5e565b9150506111b9565b5060006113a985858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040516113a392507f989d14110ba109ccad392cc18511d1f6ae3a85165c5960e49d72c2c67682fde59150611358908790602001612c5c565b60405160208183030381529060405280519060200120604051602001611388929190918252602082015260400190565b60405160208183030381529060405280519060200120611d56565b90611daa565b60005461010090046001600160a01b039081169116149695505050505050565b6113e16000356001600160e01b031916610b4c611bea565b60405163025a3c9960e21b815282906001600160a01b03821690630968f26490611082908590600401612a79565b6114276000356001600160e01b031916610b4c611bea565b60008281526004602081905260409182902080546001600160a01b0319166001600160a01b038781169182179092559251635c7d1b9b60e11b81529182018590528316602482015284919063b8fa373690604401600060405180830381600087803b15801561149557600080fd5b505af11580156114a9573d6000803e3d6000fd5b5050505050505050565b6114cb6000356001600160e01b031916610b4c611bea565b6001600160a01b03919091166000908152600560205260409020805460ff1916911515919091179055565b61150e6000356001600160e01b031916610b4c611bea565b6001600160a01b03811661156e5760405162461bcd60e51b815260206004820152602160248201527f4d504320616464726573732063616e2774206265206e756c6c2d6164647265736044820152607360f81b6064820152608401610635565b60005461010090046001600160a01b0316156115cc5760405162461bcd60e51b815260206004820152601c60248201527f4d504320616464726573732063616e27742062652075706461746564000000006044820152606401610635565b60008054610100600160a81b0319166101006001600160a01b038416021790556115fc6115f7611bea565b611dce565b6040517f4187686ceef7b541a1f224d48d4cded8f2c535e0e58ac0f0514071b1de3dad5790600090a150565b6116406000356001600160e01b031916610b4c611bea565b7fe78d813a9260522f81d6c761e311727b2e19008daadd2b9e174be86bc4f06a4b81604051610b299190612a79565b6116876000356001600160e01b031916610b4c611bea565b60ff82166000908152600360205260409020546001600160401b03908116908216116117045760405162461bcd60e51b815260206004820152602660248201527f446f6573206e6f7420616c6c6f772064656372656d656e7473206f6620746865604482015265206e6f6e636560d01b6064820152608401610635565b60ff919091166000908152600360205260409020805467ffffffffffffffff19166001600160401b03909216919091179055565b611740611ba4565b604080516001808252818301909252600091816020015b604080516080810182526000808252602080830182905292820152606080820152825260001990920191018161175757905050905083816000815181106117a0576117a06128f2565b60200260200101819052506117e0816000815181106117c1576117c16128f2565b602002602001015160000151826000815181106106ce576106ce6128f2565b1515600114156118475760405162461bcd60e51b815260206004820152602c60248201527f4465706f73697420776974682070726f7669646564206e6f6e636520616c726560448201526b18591e48195e1958dd5d195960a21b6064820152608401610635565b61185281848461116e565b6118985760405162461bcd60e51b815260206004820152601760248201527624b73b30b634b210383937b837b9b0b61039b4b3b732b960491b6044820152606401610635565b600060046000836000815181106118b1576118b16128f2565b602002602001015160400151815260200190815260200160002060009054906101000a90046001600160a01b03169050600081836000815181106118f7576118f76128f2565b602002602001015160600151604051602001611914929190612934565b604051602081830303815290604052805190602001209050600082905061010084600081518110611947576119476128f2565b60200260200101516020015161195d9190612982565b6001600160401b03166001901b6006600086600081518110611981576119816128f2565b60200260200101516000015160ff1660ff1681526020019081526020016000206000610100876000815181106119b9576119b96128f2565b6020026020010151602001516119cf91906129be565b6001600160401b0316815260200190815260200160002060008282541792505081905550806001600160a01b031663e248cff285600081518110611a1557611a156128f2565b60200260200101516040015186600081518110611a3457611a346128f2565b6020026020010151606001516040518363ffffffff1660e01b8152600401611a5d929190612a10565b600060405180830381600087803b158015611a7757600080fd5b505af1158015611a8b573d6000803e3d6000fd5b505050507f6018c584b8d99bafeda249b2429f5907d830e792222070c1b3a94aa76ee7167784600081518110611ac357611ac36128f2565b60200260200101516000015185600081518110611ae257611ae26128f2565b60200260200101516020015184604051611b1d9392919060ff9390931683526001600160401b03919091166020830152604082015260600190565b60405180910390a150505050505050565b611b466000356001600160e01b031916610b4c611bea565b60005461010090046001600160a01b0316611b995760405162461bcd60e51b8152602060048201526013602482015272135410c81859191c995cdcc81b9bdd081cd95d606a1b6044820152606401610635565b610fd56115f7611bea565b60005460ff1615610fd55760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610635565b60003360143610801590611c1657506001600160a01b03811660009081526005602052604090205460ff165b15611c26575060131936013560601c5b919050565b600254604051631c72548760e21b81526001600160e01b0319841660048201526001600160a01b038381166024830152909116906371c9521c90604401602060405180830381865afa158015611c85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ca99190612c92565b611d045760405162461bcd60e51b815260206004820152602660248201527f73656e64657220646f65736e277420686176652061636365737320746f2066756044820152653731ba34b7b760d11b6064820152608401610635565b5050565b611d10611ba4565b6000805460ff191660011790556040516001600160a01b03821681527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001610b29565b6000611da4611d63611e19565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b92915050565b6000806000611db98585611f40565b91509150611dc681611fb0565b509392505050565b611dd661216e565b6000805460ff191690556040516001600160a01b03821681527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa90602001610b29565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015611e7257507f000000000000000000000000000000000000000000000000000000000000000046145b15611e9c57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b600080825160411415611f775760208301516040840151606085015160001a611f6b878285856121b7565b94509450505050611fa9565b825160401415611fa15760208301516040840151611f968683836122a4565b935093505050611fa9565b506000905060025b9250929050565b6000816004811115611fc457611fc4612caf565b1415611fcd5750565b6001816004811115611fe157611fe1612caf565b141561202f5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610635565b600281600481111561204357612043612caf565b14156120915760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610635565b60038160048111156120a5576120a5612caf565b14156120fe5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610635565b600481600481111561211257612112612caf565b141561216b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610635565b50565b60005460ff16610fd55760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610635565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156121ee575060009050600361229b565b8460ff16601b1415801561220657508460ff16601c14155b15612217575060009050600461229b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561226b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166122945760006001925092505061229b565b9150600090505b94509492505050565b6000806001600160ff1b038316816122c160ff86901c601b612cc5565b90506122cf878288856121b7565b935093505050935093915050565b803560ff81168114611c2657600080fd5b6000806040838503121561230157600080fd5b61230a836122dd565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561235657612356612318565b604052919050565b80356001600160401b0381168114611c2657600080fd5b60006001600160401b0382111561238e5761238e612318565b50601f01601f191660200190565b60006123af6123aa84612375565b61232e565b90508281528383830111156123c357600080fd5b828260208301376000602084830101529392505050565b600082601f8301126123eb57600080fd5b6123fa8383356020850161239c565b9392505050565b60006080828403121561241357600080fd5b604051608081016001600160401b03828210818311171561243657612436612318565b81604052829350612446856122dd565b83526124546020860161235e565b602084015260408501356040840152606085013591508082111561247757600080fd5b50612484858286016123da565b6060830152505092915050565b60008083601f8401126124a357600080fd5b5081356001600160401b038111156124ba57600080fd5b602083019150836020828501011115611fa957600080fd5b6000806000604084860312156124e757600080fd5b83356001600160401b03808211156124fe57600080fd5b818601915086601f83011261251257600080fd5b813560208282111561252657612526612318565b8160051b61253582820161232e565b928352848101820192828101908b85111561254f57600080fd5b83870192505b8483101561258b5782358681111561256d5760008081fd5b61257b8d86838b0101612401565b8352509183019190830190612555565b98505050870135925050808211156125a257600080fd5b506125af86828701612491565b9497909650939450505050565b6000602082840312156125ce57600080fd5b81356001600160401b038111156125e457600080fd5b8201601f810184136125f557600080fd5b6126048482356020840161239c565b949350505050565b60006020828403121561261e57600080fd5b6123fa826122dd565b80356001600160a01b0381168114611c2657600080fd5b80356001600160e01b031981168114611c2657600080fd5b60008060008060008060c0878903121561266f57600080fd5b61267887612627565b95506020870135945061268d60408801612627565b935061269b6060880161263e565b9250608087013591506126b060a0880161263e565b90509295509295509295565b600080600080600080608087890312156126d557600080fd5b6126de876122dd565b95506020870135945060408701356001600160401b038082111561270157600080fd5b61270d8a838b01612491565b9096509450606089013591508082111561272657600080fd5b5061273389828a01612491565b979a9699509497509295939492505050565b60006020828403121561275757600080fd5b5035919050565b60006020828403121561277057600080fd5b6123fa82612627565b6000806040838503121561278c57600080fd5b61279583612627565b91506127a360208401612627565b90509250929050565b600080604083850312156127bf57600080fd5b6127c883612627565b915060208301356001600160401b038111156127e357600080fd5b6127ef858286016123da565b9150509250929050565b60008060006060848603121561280e57600080fd5b61281784612627565b92506020840135915061282c60408501612627565b90509250925092565b801515811461216b57600080fd5b6000806040838503121561285657600080fd5b61285f83612627565b9150602083013561286f81612835565b809150509250929050565b6000806040838503121561288d57600080fd5b612896836122dd565b91506127a36020840161235e565b6000806000604084860312156128b957600080fd5b83356001600160401b03808211156128d057600080fd5b6128dc87838801612401565b945060208601359150808211156125a257600080fd5b634e487b7160e01b600052603260045260246000fd5b60005b8381101561292357818101518382015260200161290b565b83811115610af45750506000910152565b6bffffffffffffffffffffffff198360601b1681526000825161295e816014850160208701612908565b919091016014019392505050565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b038084168061299c5761299c61296c565b92169190910692915050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03808416806129d8576129d861296c565b92169190910492915050565b600081518084526129fc816020860160208601612908565b601f01601f19169290920160200192915050565b82815260406020820152600061260460408301846129e4565b606081526000612a3c60608301866129e4565b905060ff841660208301526001600160401b0383166040830152949350505050565b6000600019821415612a7257612a726129a8565b5060010190565b6020815260006123fa60208301846129e4565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b038916815260ff8816602082015260ff8716604082015285606082015260c060808201526000612aef60c083018688612a8c565b82810360a0840152612b02818587612a8c565b9b9a5050505050505050505050565b60006001600160401b0380831681811415612b2e57612b2e6129a8565b6001019392505050565b8481526001600160a01b0384166020820152606060408201819052600090612b639083018486612a8c565b9695505050505050565b600060208284031215612b7f57600080fd5b81516001600160401b03811115612b9557600080fd5b8201601f81018413612ba657600080fd5b8051612bb46123aa82612375565b818152856020838501011115612bc957600080fd5b612bda826020830160208601612908565b95945050505050565b60ff871681528560208201526001600160401b038516604082015260a060608201526000612c1560a083018587612a8c565b8281036080840152612c2781856129e4565b9998505050505050505050565b600082612c4357612c4361296c565b500690565b600082612c5757612c5761296c565b500490565b815160009082906020808601845b83811015612c8657815185529382019390820190600101612c6a565b50929695505050505050565b600060208284031215612ca457600080fd5b81516123fa81612835565b634e487b7160e01b600052602160045260246000fd5b60008219821115612cd857612cd86129a8565b50019056fea26469706673582212208e6627a62e6de505ca2ec55e22d398b3e08080b65d6c2dabe265f850e9bf0cc964736f6c634300080b0033"
