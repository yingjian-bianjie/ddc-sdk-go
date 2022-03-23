// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChargeMetaData contains all meta data concerning the Charge contract.
var ChargeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"}],\"name\":\"DelDDC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"DelFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accAddr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"}],\"name\":\"delDDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"delFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"queryFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"selfRecharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorityProxyAddress\",\"type\":\"address\"}],\"name\":\"setAuthorityProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"settlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50600054610100900460ff168061002e575060005460ff16155b6100955760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff161580156100b7576000805461ffff19166101011790555b80156100c9576000805461ff00191690555b506080516120ca6100fa600039600081816104c3015281816105030152818161058c01526105cc01526120ca6000f3fe6080604052600436106100fd5760003560e01c8063715018a611610095578063c5837d8211610064578063c5837d821461026d578063c9c45a0f1461028d578063d213fe45146102c2578063ef18e3c9146102e2578063f2fde38b1461030257600080fd5b8063715018a6146101fb5780638129fc1c146102105780638a27a80d146102255780638da5cb5b1461024557600080fd5b80633659cfe6116100d15780633659cfe6146101885780634f1ef286146101a857806363569189146101bb57806370a08231146101db57600080fd5b80620b8f7014610102578063093f28e01461012457806318160ddd1461014457806336351c7c14610168575b600080fd5b34801561010e57600080fd5b5061012261011d366004611a09565b610322565b005b34801561013057600080fd5b5061012261013f366004611a3c565b6103a5565b34801561015057600080fd5b5060c9545b6040519081526020015b60405180910390f35b34801561017457600080fd5b50610122610183366004611a6f565b610441565b34801561019457600080fd5b506101226101a3366004611a09565b6104b8565b6101226101b6366004611b08565b610581565b3480156101c757600080fd5b506101226101d6366004611b99565b61063b565b3480156101e757600080fd5b506101556101f6366004611a09565b6106fb565b34801561020757600080fd5b5061012261073f565b34801561021c57600080fd5b50610122610775565b34801561023157600080fd5b50610122610240366004611a6f565b6107f0565b34801561025157600080fd5b506033546040516001600160a01b03909116815260200161015f565b34801561027957600080fd5b50610122610288366004611a09565b6108de565b34801561029957600080fd5b506102ad6102a8366004611a3c565b610980565b60405163ffffffff909116815260200161015f565b3480156102ce57600080fd5b506101226102dd366004611be9565b610a5d565b3480156102ee57600080fd5b506101226102fd366004611c02565b610ac8565b34801561030e57600080fd5b5061012261031d366004611a09565b610bc5565b6001600160a01b0381166103515760405162461bcd60e51b815260040161034890611c3e565b60405180910390fd5b610359610c5d565b6001600160a01b038116600081815260ca6020526040808220600101805460ff19169055517f0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e89190a250565b6001600160a01b0382166103cb5760405162461bcd60e51b815260040161034890611c3e565b6103d3610c5d565b6001600160a01b038216600081815260ca602090815260408083206001600160e01b0319861680855290835292819020805463ffffffff19169055519182527f2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d910160405180910390a25050565b8061045e5760405162461bcd60e51b815260040161034890611c6b565b6104683383610d2a565b610473338383610e44565b6040518181526001600160a01b0383169033907f4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1906020015b60405180910390a35050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105015760405162461bcd60e51b815260040161034890611ca0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610533610f36565b6001600160a01b0316146105595760405162461bcd60e51b815260040161034890611cec565b61056281610f64565b6040805160008082526020820190925261057e91839190610f8e565b50565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105ca5760405162461bcd60e51b815260040161034890611ca0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105fc610f36565b6001600160a01b0316146106225760405162461bcd60e51b815260040161034890611cec565b61062b82610f64565b61063782826001610f8e565b5050565b6001600160a01b0383166106615760405162461bcd60e51b815260040161034890611c3e565b610669610c5d565b6001600160a01b038316600081815260ca602090815260408083206001808201805460ff191690911790556001600160e01b0319871680855290835292819020805463ffffffff191663ffffffff87169081179091558151938452918301919091527f929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc910160405180910390a2505050565b60006001600160a01b0382166107235760405162461bcd60e51b815260040161034890611c3e565b506001600160a01b0316600090815260cb602052604090205490565b6033546001600160a01b031633146107695760405162461bcd60e51b815260040161034890611d38565b61077360006110d9565b565b600054610100900460ff168061078e575060005460ff16155b6107aa5760405162461bcd60e51b815260040161034890611d6d565b600054610100900460ff161580156107cc576000805461ffff19166101011790555b6107d461112b565b6107dc611192565b801561057e576000805461ff001916905550565b8061080d5760405162461bcd60e51b815260040161034890611c6b565b6001600160a01b0382163b1515801561084157506001600160a01b038216600090815260ca602052604090206001015460ff165b61088d5760405162461bcd60e51b815260206004820152601860248201527f6368617267653a206e6f742044444320636f6e747261637400000000000000006044820152606401610348565b610895610c5d565b6108a182335b83610e44565b6040518181526001600160a01b0383169033907fca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be906020016104ac565b6033546001600160a01b031633146109085760405162461bcd60e51b815260040161034890611d38565b6001600160a01b03811661095e5760405162461bcd60e51b815260206004820152601d60248201527f6368617267653a206175746820746865207a65726f20616464726573730000006044820152606401610348565b60cc80546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160a01b0383166109a85760405162461bcd60e51b815260040161034890611c3e565b6001600160a01b038316600090815260ca602052604090206001015460ff16610a215760405162461bcd60e51b815260206004820152602560248201527f6368617267653a6464632070726f787920636f6e747261637420756e617661696044820152646c61626c6560d81b6064820152608401610348565b506001600160a01b038216600090815260ca602090815260408083206001600160e01b03198516845290915290205463ffffffff165b92915050565b80610a7a5760405162461bcd60e51b815260040161034890611c6b565b610a82610c5d565b610a8d60003361089b565b60405181815233906000907f4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb19060200160405180910390a350565b6001600160a01b038316610aee5760405162461bcd60e51b815260040161034890611c3e565b80610b325760405162461bcd60e51b815260206004820152601460248201527318da185c99d94e9a5b9d985b1a5908191918d25960621b6044820152606401610348565b336000610b3f8285610980565b905063ffffffff811615610b5e57610b5e85838363ffffffff16610e44565b604080516001600160e01b03198616815263ffffffff831660208201529081018490526001600160a01b0380841691908716907f750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab9060600160405180910390a35050505050565b6033546001600160a01b03163314610bef5760405162461bcd60e51b815260040161034890611d38565b6001600160a01b038116610c545760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610348565b61057e816110d9565b60cc546001600160a01b031663ed5cad643360006040518363ffffffff1660e01b8152600401610c8e929190611dd1565b602060405180830381865afa158015610cab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ccf9190611e0c565b6107735760405162461bcd60e51b815260206004820152602660248201527f4444433732313a6e6f742061206f70657261746f7220726f6c65206f722064696044820152651cd8589b195960d21b6064820152608401610348565b6001600160a01b038116610d8c5760405162461bcd60e51b8152602060048201526024808201527f6368617267653a20726563686172676520746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610348565b806001600160a01b0316826001600160a01b03161415610dee5760405162461bcd60e51b815260206004820181905260248201527f6368617267653a206e6f207265636861726765206973206e65636573736172796044820152606401610348565b610df882826111f9565b6106375760405162461bcd60e51b815260206004820152601e60248201527f6368617267653a206e6f207265636861726765207065726d697373696f6e00006044820152606401610348565b6001600160a01b03831615610eec5780610e5d846106fb565b1015610eb95760405162461bcd60e51b815260206004820152602560248201527f6368617267653a206163636f756e742062616c616e6365206973206e6f7420656044820152640dcdeeaced60db1b6064820152608401610348565b6001600160a01b038316600090815260cb602052604081208054839290610ee1908490611e44565b90915550610f049050565b8060c96000828254610efe9190611e5b565b90915550505b6001600160a01b038216600090815260cb602052604081208054839290610f2c908490611e5b565b9091555050505050565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6033546001600160a01b0316331461057e5760405162461bcd60e51b815260040161034890611d38565b6000610f98610f36565b9050610fa38461167b565b600083511180610fb05750815b15610fc157610fbf8484611720565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff166110d257805460ff191660011781556040516001600160a01b038316602482015261104090869060440160408051601f198184030181529190526020810180516001600160e01b0316631b2ce7f360e11b179052611720565b50805460ff19168155611051610f36565b6001600160a01b0316826001600160a01b0316146110c95760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201526e75727468657220757067726164657360881b6064820152608401610348565b6110d285611802565b5050505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680611144575060005460ff16155b6111605760405162461bcd60e51b815260040161034890611d6d565b600054610100900460ff16158015611182576000805461ffff19166101011790555b61118a611842565b6107dc6118ac565b600054610100900460ff16806111ab575060005460ff16155b6111c75760405162461bcd60e51b815260040161034890611d6d565b600054610100900460ff161580156111e9576000805461ffff19166101011790555b6111f1611842565b6107dc611842565b60006112396040805160e0810182526060808252602082015290810160008152606060208201526040016000815260200160008152602001606081525090565b60cc5460405163fbcbc0f160e01b81526001600160a01b0386811660048301529091169063fbcbc0f190602401600060405180830381865afa158015611283573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112ab9190810190611f0e565b5092935090918560408101606082016080830160a084018560018111156112d4576112d4611dbb565b60018111156112e5576112e5611dbb565b90528560018111156112f9576112f9611dbb565b600181111561130a5761130a611dbb565b905285905285600281111561132157611321611dbb565b600281111561133257611332611dbb565b9052949094525060019250611345915050565b8160800151600181111561135b5761135b611dbb565b14801561137d575060018160a00151600181111561137b5761137b611dbb565b145b6113c95760405162461bcd60e51b815260206004820152601960248201527f6368617267653a206163636f756e742069732066726f7a656e000000000000006044820152606401610348565b6114076040805160e0810182526060808252602082015290810160008152606060208201526040016000815260200160008152602001606081525090565b60cc5460405163fbcbc0f160e01b81526001600160a01b0386811660048301529091169063fbcbc0f190602401600060405180830381865afa158015611451573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114799190810190611f0e565b5092935090918560408101606082016080830160a084018560018111156114a2576114a2611dbb565b60018111156114b3576114b3611dbb565b90528560018111156114c7576114c7611dbb565b60018111156114d8576114d8611dbb565b90528590528560028111156114ef576114ef611dbb565b600281111561150057611500611dbb565b9052949094525060019250611513915050565b8160800151600181111561152957611529611dbb565b14801561154b575060018160a00151600181111561154957611549611dbb565b145b61158e5760405162461bcd60e51b815260206004820152601460248201527331b430b933b29d103a379034b990333937bd32b760611b6044820152606401610348565b6002826040015160028111156115a6576115a6611dbb565b14156115f45760405162461bcd60e51b815260206004820152601e60248201527f6368617267653a206e6f207265636861726765207065726d697373696f6e00006044820152606401610348565b60008260400151600281111561160c5761160c611dbb565b14806116235750606081015182516116239161190c565b8061167257506060808201519083015161163c9161190c565b80156116505750805182516116509161190c565b8015611672575060028160400151600281111561166f5761166f611dbb565b14155b95945050505050565b803b6116df5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610348565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b61177f5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610348565b600080846001600160a01b03168460405161179a9190611fed565b600060405180830381855af49150503d80600081146117d5576040519150601f19603f3d011682016040523d82523d6000602084013e6117da565b606091505b5091509150611672828260405180606001604052806027815260200161206e602791396119ad565b61180b8161167b565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600054610100900460ff168061185b575060005460ff16155b6118775760405162461bcd60e51b815260040161034890611d6d565b600054610100900460ff161580156107dc576000805461ffff1916610101179055801561057e576000805461ff001916905550565b600054610100900460ff16806118c5575060005460ff16155b6118e15760405162461bcd60e51b815260040161034890611d6d565b600054610100900460ff16158015611903576000805461ffff19166101011790555b6107dc336110d9565b80518251600091849184911461192757600092505050610a57565b815160005b818110156119a05782818151811061194657611946612009565b602001015160f81c60f81b6001600160f81b03191684828151811061196d5761196d612009565b01602001516001600160f81b0319161461198e576000945050505050610a57565b806119988161201f565b91505061192c565b5060019695505050505050565b606083156119bc5750816119e6565b8251156119cc5782518084602001fd5b8160405162461bcd60e51b8152600401610348919061203a565b9392505050565b80356001600160a01b0381168114611a0457600080fd5b919050565b600060208284031215611a1b57600080fd5b6119e6826119ed565b80356001600160e01b031981168114611a0457600080fd5b60008060408385031215611a4f57600080fd5b611a58836119ed565b9150611a6660208401611a24565b90509250929050565b60008060408385031215611a8257600080fd5b611a8b836119ed565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611ad857611ad8611a99565b604052919050565b600067ffffffffffffffff821115611afa57611afa611a99565b50601f01601f191660200190565b60008060408385031215611b1b57600080fd5b611b24836119ed565b9150602083013567ffffffffffffffff811115611b4057600080fd5b8301601f81018513611b5157600080fd5b8035611b64611b5f82611ae0565b611aaf565b818152866020838501011115611b7957600080fd5b816020840160208301376000602083830101528093505050509250929050565b600080600060608486031215611bae57600080fd5b611bb7846119ed565b9250611bc560208501611a24565b9150604084013563ffffffff81168114611bde57600080fd5b809150509250925092565b600060208284031215611bfb57600080fd5b5035919050565b600080600060608486031215611c1757600080fd5b611c20846119ed565b9250611c2e60208501611a24565b9150604084013590509250925092565b6020808252601390820152726368617267653a7a65726f206164647265737360681b604082015260600190565b6020808252818101527f6368617267653a206e6f207472616e73666572206973206e6563657373617279604082015260600190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b03831681526040810160038310611dff57634e487b7160e01b600052602160045260246000fd5b8260208301529392505050565b600060208284031215611e1e57600080fd5b815180151581146119e657600080fd5b634e487b7160e01b600052601160045260246000fd5b600082821015611e5657611e56611e2e565b500390565b60008219821115611e6e57611e6e611e2e565b500190565b60005b83811015611e8e578181015183820152602001611e76565b83811115611e9d576000848401525b50505050565b600082601f830112611eb457600080fd5b8151611ec2611b5f82611ae0565b818152846020838601011115611ed757600080fd5b611ee8826020830160208701611e73565b949350505050565b805160038110611a0457600080fd5b805160028110611a0457600080fd5b600080600080600080600060e0888a031215611f2957600080fd5b875167ffffffffffffffff80821115611f4157600080fd5b611f4d8b838c01611ea3565b985060208a0151915080821115611f6357600080fd5b611f6f8b838c01611ea3565b9750611f7d60408b01611ef0565b965060608a0151915080821115611f9357600080fd5b611f9f8b838c01611ea3565b9550611fad60808b01611eff565b9450611fbb60a08b01611eff565b935060c08a0151915080821115611fd157600080fd5b50611fde8a828b01611ea3565b91505092959891949750929550565b60008251611fff818460208701611e73565b9190910192915050565b634e487b7160e01b600052603260045260246000fd5b600060001982141561203357612033611e2e565b5060010190565b6020815260008251806020840152612059816040850160208701611e73565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220df6cfa1c141a84dd0e3e7ea25be18b487edeb9467c1cd1ac9e5cf96b8da4033764736f6c634300080b0033",
}

// ChargeABI is the input ABI used to generate the binding from.
// Deprecated: Use ChargeMetaData.ABI instead.
var ChargeABI = ChargeMetaData.ABI

// ChargeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChargeMetaData.Bin instead.
var ChargeBin = ChargeMetaData.Bin

// DeployCharge deploys a new Ethereum contract, binding an instance of Charge to it.
func DeployCharge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Charge, error) {
	parsed, err := ChargeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChargeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Charge{ChargeCaller: ChargeCaller{contract: contract}, ChargeTransactor: ChargeTransactor{contract: contract}, ChargeFilterer: ChargeFilterer{contract: contract}}, nil
}

// Charge is an auto generated Go binding around an Ethereum contract.
type Charge struct {
	ChargeCaller     // Read-only binding to the contract
	ChargeTransactor // Write-only binding to the contract
	ChargeFilterer   // Log filterer for contract events
}

// ChargeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChargeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChargeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChargeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChargeSession struct {
	Contract     *Charge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChargeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChargeCallerSession struct {
	Contract *ChargeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChargeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChargeTransactorSession struct {
	Contract     *ChargeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChargeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChargeRaw struct {
	Contract *Charge // Generic contract binding to access the raw methods on
}

// ChargeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChargeCallerRaw struct {
	Contract *ChargeCaller // Generic read-only contract binding to access the raw methods on
}

// ChargeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChargeTransactorRaw struct {
	Contract *ChargeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCharge creates a new instance of Charge, bound to a specific deployed contract.
func NewCharge(address common.Address, backend bind.ContractBackend) (*Charge, error) {
	contract, err := bindCharge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Charge{ChargeCaller: ChargeCaller{contract: contract}, ChargeTransactor: ChargeTransactor{contract: contract}, ChargeFilterer: ChargeFilterer{contract: contract}}, nil
}

// NewChargeCaller creates a new read-only instance of Charge, bound to a specific deployed contract.
func NewChargeCaller(address common.Address, caller bind.ContractCaller) (*ChargeCaller, error) {
	contract, err := bindCharge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChargeCaller{contract: contract}, nil
}

// NewChargeTransactor creates a new write-only instance of Charge, bound to a specific deployed contract.
func NewChargeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChargeTransactor, error) {
	contract, err := bindCharge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChargeTransactor{contract: contract}, nil
}

// NewChargeFilterer creates a new log filterer instance of Charge, bound to a specific deployed contract.
func NewChargeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChargeFilterer, error) {
	contract, err := bindCharge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChargeFilterer{contract: contract}, nil
}

// bindCharge binds a generic wrapper to an already deployed contract.
func bindCharge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChargeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Charge *ChargeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Charge.Contract.ChargeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Charge *ChargeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.Contract.ChargeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Charge *ChargeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Charge.Contract.ChargeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Charge *ChargeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Charge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Charge *ChargeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Charge *ChargeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Charge.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_Charge *ChargeCaller) BalanceOf(opts *bind.CallOpts, accAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "balanceOf", accAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_Charge *ChargeSession) BalanceOf(accAddr common.Address) (*big.Int, error) {
	return _Charge.Contract.BalanceOf(&_Charge.CallOpts, accAddr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_Charge *ChargeCallerSession) BalanceOf(accAddr common.Address) (*big.Int, error) {
	return _Charge.Contract.BalanceOf(&_Charge.CallOpts, accAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Charge *ChargeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Charge *ChargeSession) Owner() (common.Address, error) {
	return _Charge.Contract.Owner(&_Charge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Charge *ChargeCallerSession) Owner() (common.Address, error) {
	return _Charge.Contract.Owner(&_Charge.CallOpts)
}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_Charge *ChargeCaller) QueryFee(opts *bind.CallOpts, ddcAddr common.Address, sig [4]byte) (uint32, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "queryFee", ddcAddr, sig)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_Charge *ChargeSession) QueryFee(ddcAddr common.Address, sig [4]byte) (uint32, error) {
	return _Charge.Contract.QueryFee(&_Charge.CallOpts, ddcAddr, sig)
}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_Charge *ChargeCallerSession) QueryFee(ddcAddr common.Address, sig [4]byte) (uint32, error) {
	return _Charge.Contract.QueryFee(&_Charge.CallOpts, ddcAddr, sig)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Charge *ChargeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Charge *ChargeSession) TotalSupply() (*big.Int, error) {
	return _Charge.Contract.TotalSupply(&_Charge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Charge *ChargeCallerSession) TotalSupply() (*big.Int, error) {
	return _Charge.Contract.TotalSupply(&_Charge.CallOpts)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_Charge *ChargeTransactor) DelDDC(opts *bind.TransactOpts, ddcAddr common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "delDDC", ddcAddr)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_Charge *ChargeSession) DelDDC(ddcAddr common.Address) (*types.Transaction, error) {
	return _Charge.Contract.DelDDC(&_Charge.TransactOpts, ddcAddr)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_Charge *ChargeTransactorSession) DelDDC(ddcAddr common.Address) (*types.Transaction, error) {
	return _Charge.Contract.DelDDC(&_Charge.TransactOpts, ddcAddr)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_Charge *ChargeTransactor) DelFee(opts *bind.TransactOpts, ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "delFee", ddcAddr, sig)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_Charge *ChargeSession) DelFee(ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _Charge.Contract.DelFee(&_Charge.TransactOpts, ddcAddr, sig)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_Charge *ChargeTransactorSession) DelFee(ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _Charge.Contract.DelFee(&_Charge.TransactOpts, ddcAddr, sig)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Charge *ChargeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Charge *ChargeSession) Initialize() (*types.Transaction, error) {
	return _Charge.Contract.Initialize(&_Charge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Charge *ChargeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Charge.Contract.Initialize(&_Charge.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_Charge *ChargeTransactor) Pay(opts *bind.TransactOpts, payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "pay", payer, sig, ddcId)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_Charge *ChargeSession) Pay(payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Pay(&_Charge.TransactOpts, payer, sig, ddcId)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_Charge *ChargeTransactorSession) Pay(payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Pay(&_Charge.TransactOpts, payer, sig, ddcId)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_Charge *ChargeTransactor) Recharge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "recharge", to, amount)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_Charge *ChargeSession) Recharge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Recharge(&_Charge.TransactOpts, to, amount)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_Charge *ChargeTransactorSession) Recharge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Recharge(&_Charge.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Charge *ChargeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Charge *ChargeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Charge.Contract.RenounceOwnership(&_Charge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Charge *ChargeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Charge.Contract.RenounceOwnership(&_Charge.TransactOpts)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_Charge *ChargeTransactor) SelfRecharge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "selfRecharge", amount)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_Charge *ChargeSession) SelfRecharge(amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.SelfRecharge(&_Charge.TransactOpts, amount)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_Charge *ChargeTransactorSession) SelfRecharge(amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.SelfRecharge(&_Charge.TransactOpts, amount)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_Charge *ChargeTransactor) SetAuthorityProxyAddress(opts *bind.TransactOpts, authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "setAuthorityProxyAddress", authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_Charge *ChargeSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _Charge.Contract.SetAuthorityProxyAddress(&_Charge.TransactOpts, authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_Charge *ChargeTransactorSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _Charge.Contract.SetAuthorityProxyAddress(&_Charge.TransactOpts, authorityProxyAddress)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_Charge *ChargeTransactor) SetFee(opts *bind.TransactOpts, ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "setFee", ddcAddr, sig, amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_Charge *ChargeSession) SetFee(ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _Charge.Contract.SetFee(&_Charge.TransactOpts, ddcAddr, sig, amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_Charge *ChargeTransactorSession) SetFee(ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _Charge.Contract.SetFee(&_Charge.TransactOpts, ddcAddr, sig, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_Charge *ChargeTransactor) Settlement(opts *bind.TransactOpts, ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "settlement", ddcAddr, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_Charge *ChargeSession) Settlement(ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Settlement(&_Charge.TransactOpts, ddcAddr, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_Charge *ChargeTransactorSession) Settlement(ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Settlement(&_Charge.TransactOpts, ddcAddr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Charge *ChargeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Charge *ChargeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Charge.Contract.TransferOwnership(&_Charge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Charge *ChargeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Charge.Contract.TransferOwnership(&_Charge.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Charge *ChargeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Charge *ChargeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeTo(&_Charge.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Charge *ChargeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeTo(&_Charge.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Charge *ChargeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Charge *ChargeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeToAndCall(&_Charge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Charge *ChargeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeToAndCall(&_Charge.TransactOpts, newImplementation, data)
}

// ChargeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Charge contract.
type ChargeAdminChangedIterator struct {
	Event *ChargeAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeAdminChanged represents a AdminChanged event raised by the Charge contract.
type ChargeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Charge *ChargeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ChargeAdminChangedIterator, error) {

	logs, sub, err := _Charge.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ChargeAdminChangedIterator{contract: _Charge.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Charge *ChargeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ChargeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Charge.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeAdminChanged)
				if err := _Charge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Charge *ChargeFilterer) ParseAdminChanged(log types.Log) (*ChargeAdminChanged, error) {
	event := new(ChargeAdminChanged)
	if err := _Charge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Charge contract.
type ChargeBeaconUpgradedIterator struct {
	Event *ChargeBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeBeaconUpgraded represents a BeaconUpgraded event raised by the Charge contract.
type ChargeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Charge *ChargeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ChargeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ChargeBeaconUpgradedIterator{contract: _Charge.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Charge *ChargeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ChargeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeBeaconUpgraded)
				if err := _Charge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Charge *ChargeFilterer) ParseBeaconUpgraded(log types.Log) (*ChargeBeaconUpgraded, error) {
	event := new(ChargeBeaconUpgraded)
	if err := _Charge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeDelDDCIterator is returned from FilterDelDDC and is used to iterate over the raw logs and unpacked data for DelDDC events raised by the Charge contract.
type ChargeDelDDCIterator struct {
	Event *ChargeDelDDC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeDelDDCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeDelDDC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeDelDDC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeDelDDCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeDelDDCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeDelDDC represents a DelDDC event raised by the Charge contract.
type ChargeDelDDC struct {
	DdcAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelDDC is a free log retrieval operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_Charge *ChargeFilterer) FilterDelDDC(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeDelDDCIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "DelDDC", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeDelDDCIterator{contract: _Charge.contract, event: "DelDDC", logs: logs, sub: sub}, nil
}

// WatchDelDDC is a free log subscription operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_Charge *ChargeFilterer) WatchDelDDC(opts *bind.WatchOpts, sink chan<- *ChargeDelDDC, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "DelDDC", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeDelDDC)
				if err := _Charge.contract.UnpackLog(event, "DelDDC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelDDC is a log parse operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_Charge *ChargeFilterer) ParseDelDDC(log types.Log) (*ChargeDelDDC, error) {
	event := new(ChargeDelDDC)
	if err := _Charge.contract.UnpackLog(event, "DelDDC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeDelFeeIterator is returned from FilterDelFee and is used to iterate over the raw logs and unpacked data for DelFee events raised by the Charge contract.
type ChargeDelFeeIterator struct {
	Event *ChargeDelFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeDelFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeDelFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeDelFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeDelFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeDelFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeDelFee represents a DelFee event raised by the Charge contract.
type ChargeDelFee struct {
	DdcAddr common.Address
	Sig     [4]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelFee is a free log retrieval operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_Charge *ChargeFilterer) FilterDelFee(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeDelFeeIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "DelFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeDelFeeIterator{contract: _Charge.contract, event: "DelFee", logs: logs, sub: sub}, nil
}

// WatchDelFee is a free log subscription operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_Charge *ChargeFilterer) WatchDelFee(opts *bind.WatchOpts, sink chan<- *ChargeDelFee, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "DelFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeDelFee)
				if err := _Charge.contract.UnpackLog(event, "DelFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelFee is a log parse operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_Charge *ChargeFilterer) ParseDelFee(log types.Log) (*ChargeDelFee, error) {
	event := new(ChargeDelFee)
	if err := _Charge.contract.UnpackLog(event, "DelFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Charge contract.
type ChargeOwnershipTransferredIterator struct {
	Event *ChargeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeOwnershipTransferred represents a OwnershipTransferred event raised by the Charge contract.
type ChargeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Charge *ChargeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChargeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChargeOwnershipTransferredIterator{contract: _Charge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Charge *ChargeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChargeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeOwnershipTransferred)
				if err := _Charge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Charge *ChargeFilterer) ParseOwnershipTransferred(log types.Log) (*ChargeOwnershipTransferred, error) {
	event := new(ChargeOwnershipTransferred)
	if err := _Charge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargePayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the Charge contract.
type ChargePayIterator struct {
	Event *ChargePay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargePayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargePay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargePay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargePayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargePayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargePay represents a Pay event raised by the Charge contract.
type ChargePay struct {
	Payer  common.Address
	Payee  common.Address
	Sig    [4]byte
	Amount uint32
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_Charge *ChargeFilterer) FilterPay(opts *bind.FilterOpts, payer []common.Address, payee []common.Address) (*ChargePayIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Pay", payerRule, payeeRule)
	if err != nil {
		return nil, err
	}
	return &ChargePayIterator{contract: _Charge.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_Charge *ChargeFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *ChargePay, payer []common.Address, payee []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Pay", payerRule, payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargePay)
				if err := _Charge.contract.UnpackLog(event, "Pay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePay is a log parse operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_Charge *ChargeFilterer) ParsePay(log types.Log) (*ChargePay, error) {
	event := new(ChargePay)
	if err := _Charge.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeRechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the Charge contract.
type ChargeRechargeIterator struct {
	Event *ChargeRecharge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeRecharge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeRecharge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeRecharge represents a Recharge event raised by the Charge contract.
type ChargeRecharge struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_Charge *ChargeFilterer) FilterRecharge(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChargeRechargeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Recharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChargeRechargeIterator{contract: _Charge.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_Charge *ChargeFilterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *ChargeRecharge, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Recharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeRecharge)
				if err := _Charge.contract.UnpackLog(event, "Recharge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRecharge is a log parse operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_Charge *ChargeFilterer) ParseRecharge(log types.Log) (*ChargeRecharge, error) {
	event := new(ChargeRecharge)
	if err := _Charge.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the Charge contract.
type ChargeSetFeeIterator struct {
	Event *ChargeSetFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeSetFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeSetFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeSetFee represents a SetFee event raised by the Charge contract.
type ChargeSetFee struct {
	DdcAddr common.Address
	Sig     [4]byte
	Amount  uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_Charge *ChargeFilterer) FilterSetFee(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeSetFeeIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "SetFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeSetFeeIterator{contract: _Charge.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_Charge *ChargeFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *ChargeSetFee, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "SetFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeSetFee)
				if err := _Charge.contract.UnpackLog(event, "SetFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFee is a log parse operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_Charge *ChargeFilterer) ParseSetFee(log types.Log) (*ChargeSetFee, error) {
	event := new(ChargeSetFee)
	if err := _Charge.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the Charge contract.
type ChargeSettlementIterator struct {
	Event *ChargeSettlement // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeSettlement)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeSettlement)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeSettlement represents a Settlement event raised by the Charge contract.
type ChargeSettlement struct {
	AccAddr common.Address
	DdcAddr common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_Charge *ChargeFilterer) FilterSettlement(opts *bind.FilterOpts, accAddr []common.Address, ddcAddr []common.Address) (*ChargeSettlementIterator, error) {

	var accAddrRule []interface{}
	for _, accAddrItem := range accAddr {
		accAddrRule = append(accAddrRule, accAddrItem)
	}
	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Settlement", accAddrRule, ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeSettlementIterator{contract: _Charge.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_Charge *ChargeFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *ChargeSettlement, accAddr []common.Address, ddcAddr []common.Address) (event.Subscription, error) {

	var accAddrRule []interface{}
	for _, accAddrItem := range accAddr {
		accAddrRule = append(accAddrRule, accAddrItem)
	}
	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Settlement", accAddrRule, ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeSettlement)
				if err := _Charge.contract.UnpackLog(event, "Settlement", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettlement is a log parse operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_Charge *ChargeFilterer) ParseSettlement(log types.Log) (*ChargeSettlement, error) {
	event := new(ChargeSettlement)
	if err := _Charge.contract.UnpackLog(event, "Settlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Charge contract.
type ChargeUpgradedIterator struct {
	Event *ChargeUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeUpgraded represents a Upgraded event raised by the Charge contract.
type ChargeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Charge *ChargeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChargeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChargeUpgradedIterator{contract: _Charge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Charge *ChargeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChargeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeUpgraded)
				if err := _Charge.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Charge *ChargeFilterer) ParseUpgraded(log types.Log) (*ChargeUpgraded, error) {
	event := new(ChargeUpgraded)
	if err := _Charge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
