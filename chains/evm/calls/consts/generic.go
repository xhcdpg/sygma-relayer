package consts

const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionDepositerOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositerOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
const GenericHandlerBin = "0x60a060405234801561001057600080fd5b50604051610b02380380610b0283398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610a716100916000396000818160a801526106cc0152610a716000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063c54c2a1111610066578063c54c2a11146101a4578063cb624463146101cd578063de319d99146101f0578063e248cff214610205578063ec97d3b41461021857600080fd5b8063318c136e146100a35780637f79bea8146100e7578063a5c3a9851461011a578063aa50800b14610156578063b07e54bb14610184575b600080fd5b6100ca7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61010a6100f5366004610757565b60056020526000908152604090205460ff1681565b60405190151581526020016100de565b61013d610128366004610757565b60046020526000908152604090205460e01b81565b6040516001600160e01b031990911681526020016100de565b610176610164366004610757565b60036020526000908152604090205481565b6040519081526020016100de565b6101976101923660046107c2565b610238565b6040516100de919061084c565b6100ca6101b236600461087f565b6000602081905290815260409020546001600160a01b031681565b61013d6101db366004610757565b60026020526000908152604090205460e01b81565b6102036101fe3660046108b0565b610491565b005b610203610213366004610907565b610524565b610176610226366004610757565b60016020526000908152604090205481565b60606102426106c1565b600060606102528486018661087f565b9150846020856102628583610953565b9261026f93929190610979565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052508b815260208181526040808320546001600160a01b0316808452600390925290912054949550939250508115905061033b57828101602001516001600160a01b038916606082901c146103395760405162461bcd60e51b815260206004820152601f60248201527f696e636f7272656374206465706f736974657220696e2074686520646174610060448201526064015b60405180910390fd5b505b6001600160a01b03821660009081526005602052604090205460ff166103735760405162461bcd60e51b8152600401610330906109a3565b6001600160a01b03821660009081526002602052604090205460e01b6001600160e01b031981161561048357600081856040516020016103b49291906109ee565b6040516020818303038152906040529050600080856001600160a01b0316836040516103e09190610a1f565b6000604051808303816000865af19150503d806000811461041d576040519150601f19603f3d011682016040523d82523d6000602084013e610422565b606091505b5091509150816104745760405162461bcd60e51b815260206004820152601e60248201527f63616c6c20746f20636f6e747261637441646472657373206661696c656400006044820152606401610330565b97506104899650505050505050565b50505050505b949350505050565b6104996106c1565b60008581526020818152604080832080546001600160a01b0319166001600160a01b0398909816978817905595825260018082528683209790975560028152858220805463ffffffff1990811660e097881c1790915560038252868320949094556004815285822080549094169290941c91909117909155600590915220805460ff19169091179055565b61052c6106c1565b6000606061053c8385018561087f565b91508360208461054c8583610953565b9261055993929190610979565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525089815260208181526040808320546001600160a01b03168084526005909252909120549495509360ff1692506105d49150505760405162461bcd60e51b8152600401610330906109a3565b6001600160a01b03811660009081526004602052604090205460e01b6001600160e01b03198116156106b857600081846040516020016106159291906109ee565b60405160208183030381529060405290506000836001600160a01b0316826040516106409190610a1f565b6000604051808303816000865af19150503d806000811461067d576040519150601f19603f3d011682016040523d82523d6000602084013e610682565b606091505b50509050806106b5576040517f3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e9690600090a15b50505b50505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107395760405162461bcd60e51b815260206004820152601e60248201527f73656e646572206d7573742062652062726964676520636f6e747261637400006044820152606401610330565b565b80356001600160a01b038116811461075257600080fd5b919050565b60006020828403121561076957600080fd5b6107728261073b565b9392505050565b60008083601f84011261078b57600080fd5b50813567ffffffffffffffff8111156107a357600080fd5b6020830191508360208285010111156107bb57600080fd5b9250929050565b600080600080606085870312156107d857600080fd5b843593506107e86020860161073b565b9250604085013567ffffffffffffffff81111561080457600080fd5b61081087828801610779565b95989497509550505050565b60005b8381101561083757818101518382015260200161081f565b83811115610846576000848401525b50505050565b602081526000825180602084015261086b81604085016020870161081c565b601f01601f19169190910160400192915050565b60006020828403121561089157600080fd5b5035919050565b80356001600160e01b03198116811461075257600080fd5b600080600080600060a086880312156108c857600080fd5b853594506108d86020870161073b565b93506108e660408701610898565b9250606086013591506108fb60808701610898565b90509295509295909350565b60008060006040848603121561091c57600080fd5b83359250602084013567ffffffffffffffff81111561093a57600080fd5b61094686828701610779565b9497909650939450505050565b6000821982111561097457634e487b7160e01b600052601160045260246000fd5b500190565b6000808585111561098957600080fd5b8386111561099657600080fd5b5050820193919092039150565b6020808252602b908201527f70726f766964656420636f6e747261637441646472657373206973206e6f742060408201526a1dda1a5d195b1a5cdd195960aa1b606082015260800190565b6001600160e01b0319831681528151600090610a1181600485016020870161081c565b919091016004019392505050565b60008251610a3181846020870161081c565b919091019291505056fea2646970667358221220870d848d4fb0ac02f683c6ebf8a6371134093f98bee5544e25b85b7f8012eb0564736f6c634300080b0033"