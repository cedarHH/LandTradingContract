// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// LandRegistryLand is an auto generated low-level Go binding around an user-defined struct.
type LandRegistryLand struct {
	Owner         common.Address
	Location      string
	Area          *big.Int
	IsVerified    bool
	DetailsHash   string
	ReportHash    string
	DocumentsHash string
	IsVaild       bool
}

// LandRegistryUser is an auto generated low-level Go binding around an user-defined struct.
type LandRegistryUser struct {
	Username          string
	LandIdList        []*big.Int
	TransactionIdList []*big.Int
	IsVaild           bool
}

// LandTransactionTransaction is an auto generated low-level Go binding around an user-defined struct.
type LandTransactionTransaction struct {
	LandId    *big.Int
	From      common.Address
	To        common.Address
	Timestamp *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"LandRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"area\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"surveyorsAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestam\",\"type\":\"uint256\"}],\"name\":\"LandSurveying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"detailsHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reportHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"documentsHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notariesAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LandVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_area\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reportHash\",\"type\":\"string\"}],\"name\":\"LandSurveyingArea\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"addNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_surveyor\",\"type\":\"address\"}],\"name\":\"addSurveyor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"}],\"name\":\"getLandTransactionHistory\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"getTransactionDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLandTransaction.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lands\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"area\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"detailsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reportHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"documentsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"landsTransactionIdList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"notaries\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"}],\"name\":\"queryLand\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"area\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"detailsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reportHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"documentsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"internalType\":\"structLandRegistry.Land\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"landIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"transactionIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"internalType\":\"structLandRegistry.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_detailsHash\",\"type\":\"string\"}],\"name\":\"registerLand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_userName\",\"type\":\"string\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"surveyors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMapping\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_documentsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isVerified\",\"type\":\"bool\"}],\"name\":\"verifyLand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f6006553480156012575f80fd5b506040516120a03803806120a0833981016040819052602f916060565b5f80546001600160a01b03199081163317909155600180546001600160a01b0390931692909116919091179055608b565b5f60208284031215606f575f80fd5b81516001600160a01b03811681146084575f80fd5b9392505050565b612008806100985f395ff3fe608060405234801561000f575f80fd5b5060043610610131575f3560e01c80637dc0d1d0116100b45780639de8e3ea116100795780639de8e3ea146103bc578063a45c2c4c146103dc578063ad812a52146103ef578063e261f1e514610402578063e756769614610429578063f5b73cbc14610449575f80fd5b80637dc0d1d0146102d957806387b6963f146103045780638da5cb5b146103265780639437922a146103385780639ace38c21461034b575f80fd5b806314e887e8116100fa57806314e887e81461027457806349695ecc1461028b5780634fb2e45d1461029e5780635307d396146102b157806370435c99146102c6575f80fd5b8062e168f0146101355780630588565f1461016c5780630ea126f9146101815780630fa683d314610194578063118b953514610253575b5f80fd5b61015761014336600461184b565b60056020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61017f61017a36600461184b565b61045c565b005b61017f61018f36600461190a565b610494565b6102126101a2366004611955565b60408051608080820183525f8083526020808401829052838501829052606093840182905294815260078552839020835191820184528054825260018101546001600160a01b03908116958301959095526002810154909416928101929092526003909201549181019190915290565b60408051825181526020808401516001600160a01b039081169183019190915283830151169181019190915260609182015191810191909152608001610163565b61026661026136600461184b565b6105b8565b60405161016392919061199a565b61027d60095481565b604051908152602001610163565b61017f6102993660046119bd565b61065c565b61017f6102ac36600461184b565b6107d9565b6102b9610859565b6040516101639190611a52565b61017f6102d4366004611ac1565b6109ef565b6001546102ec906001600160a01b031681565b6040516001600160a01b039091168152602001610163565b61015761031236600461184b565b60046020525f908152604090205460ff1681565b5f546102ec906001600160a01b031681565b61017f610346366004611af4565b610cab565b61038c610359366004611955565b60076020525f9081526040902080546001820154600283015460039093015491926001600160a01b039182169291169084565b60405161016394939291909384526001600160a01b03928316602085015291166040830152606082015260800190565b6103cf6103ca366004611955565b610ef1565b6040516101639190611b61565b61017f6103ea366004611ba3565b610f50565b61017f6103fd36600461184b565b6110a0565b610415610410366004611955565b6110d8565b604051610163989796959493929190611bd9565b61043c610437366004611955565b611345565b6040516101639190611c5d565b61027d610457366004611d29565b61168e565b5f546001600160a01b03163314610471575f80fd5b6001600160a01b03165f908152600560205260409020805460ff19166001179055565b5f546001600160a01b031633146104a9575f80fd5b6001600160a01b0382165f9081526002602052604090206003015460ff161561050a5760405162461bcd60e51b815260206004820152600e60248201526d757365722069732065786973747360901b60448201526064015b60405180910390fd5b6040805160808101825282815281515f80825260208083018552808401929092528351818152808301855283850152600160608401526001600160a01b0386168152600290915291909120815181906105639082611dcc565b50602082810151805161057c92600185019201906117d3565b50604082015180516105989160028401916020909101906117d3565b50606091909101516003909101805460ff19169115159190911790555050565b60026020525f90815260409020805481906105d290611d49565b80601f01602080910402602001604051908101604052809291908181526020018280546105fe90611d49565b80156106495780601f1061062057610100808354040283529160200191610649565b820191905f5260205f20905b81548152906001019060200180831161062c57829003601f168201915b5050506003909301549192505060ff1682565b335f9081526005602052604090205460ff166106ab5760405162461bcd60e51b815260206004820152600e60248201526d4e6f742061206e6f74617269657360901b6044820152606401610501565b5f8381526003602052604090206007015460ff166107005760405162461bcd60e51b81526020600482015260126024820152716c616e64206973206e6f742065786973747360701b6044820152606401610501565b6001546001600160a01b031633146107515760405162461bcd60e51b815260206004820152601460248201527313db9b1e481bdc9858db194818d85b8818d85b1b60621b6044820152606401610501565b5f83815260036020526040902060060161076b8382611dcc565b505f83815260036020819052604091829020908101805460ff19168415151790559051339185917f378e3aa8b6938b8ea84d87fd3355efdb97062e22ca268d202cb08acb23e1f0d9916107cc91600482019160050190889088904290611f06565b60405180910390a3505050565b5f546001600160a01b031633146107ee575f80fd5b6001600160a01b038116610800575f80fd5b5f80546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35f80546001600160a01b0319166001600160a01b0392909216919091179055565b61088560405180608001604052806060815260200160608152602001606081526020015f151581525090565b335f90815260026020526040908190208151608081019092528054829082906108ad90611d49565b80601f01602080910402602001604051908101604052809291908181526020018280546108d990611d49565b80156109245780601f106108fb57610100808354040283529160200191610924565b820191905f5260205f20905b81548152906001019060200180831161090757829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561097a57602002820191905f5260205f20905b815481526020019060010190808311610966575b50505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156109d057602002820191905f5260205f20905b8154815260200190600101908083116109bc575b50505091835250506003919091015460ff161515602090910152919050565b826109f9816116b9565b610a355760405162461bcd60e51b815260206004820152600d60248201526c1b9bdd081e5bdd5c881b185b99609a1b6044820152606401610501565b6001600160a01b0383163303610a8d5760405162461bcd60e51b815260206004820181905260248201527f43616e6e6f74206265207472616e7366657272656420746f206f6e6573656c666044820152606401610501565b5f848152600360208190526040909120015460ff16610aee5760405162461bcd60e51b815260206004820152601960248201527f74686973206c616e64206973206e6f74205665726966696564000000000000006044820152606401610501565b6001600160a01b0383165f9081526002602052604090206003015460ff16610b4f5760405162461bcd60e51b8152602060048201526014602482015273746f55736572206973206e6f742065786973747360601b6044820152606401610501565b6001600160a01b0383165f90815260026020908152604082206001908101805491820181558352912001849055610b8584611714565b6001600160a01b038381165f81815260026020818152604080842083018054600180820183559186528386200189905533808652828620850180548084018255908752848720018a905582516080810184528c815280850182815281850189815242606084018181528e8b5260078952998790209351845591519483018054958c166001600160a01b031996871617905551968201805497909a16969093169590951790975593516003909301929092558151938452830191909152869185917f4bb86a8504e257f90fe72dcaa3a3d0dfb04b72836b57aa6882b7bc2a1b82a4d1910160405180910390a460098054905f610c7f83611f67565b9091555050505f9283526008602090815260408420805460018101825590855293209092019190915550565b335f9081526002602052604090206003015460ff16610cfa5760405162461bcd60e51b815260206004820152600b60248201526a34b9903737ba103ab9b2b960a91b6044820152606401610501565b5f8381526003602052604090206007015460ff1615610d4c5760405162461bcd60e51b815260206004820152600e60248201526d6c616e642069732065786973747360901b6044820152606401610501565b604080516101008101825233815260208082018581525f838501819052606084018190528451808401865281815260808501528451808401865281815260a08501528451808401865281815260c0850152600160e08501819052888252600390935293909320825181546001600160a01b0319166001600160a01b03909116178155925191929190820190610de19082611dcc565b5060408201516002820155606082015160038201805460ff191691151591909117905560808201516004820190610e189082611dcc565b5060a08201516005820190610e2d9082611dcc565b5060c08201516006820190610e429082611dcc565b5060e091909101516007909101805460ff19169115159190911790555f838152600360205260409020600401610e788282611dcc565b50335f818152600260209081526040808320600190810180549182018155845291909220018590555184907f9c4ad74928fec1ee86181680df0e76bf575b79ea2a0347f7ed270884bdd26ad790610ed0908690611f7f565b60405180910390a360068054905f610ee783611f67565b9190505550505050565b5f81815260086020908152604091829020805483518184028101840190945280845260609392830182828015610f4457602002820191905f5260205f20905b815481526020019060010190808311610f30575b50505050509050919050565b335f9081526004602052604090205460ff16610f9f5760405162461bcd60e51b815260206004820152600e60248201526d2737ba10309039bab93b32bcb7b960911b6044820152606401610501565b5f8381526003602052604090206007015460ff16610ff45760405162461bcd60e51b81526020600482015260126024820152716c616e64206973206e6f742065786973747360701b6044820152606401610501565b6001546001600160a01b031633146110455760405162461bcd60e51b815260206004820152601460248201527313db9b1e481bdc9858db194818d85b8818d85b1b60621b6044820152606401610501565b5f838152600360205260409020600281018390556005016110668282611dcc565b5060408051838152426020820152339185917f9f207f53ad48e112fa882ba14907e3301685086486bac5d73bab867a7c9a3c1191016107cc565b5f546001600160a01b031633146110b5575f80fd5b6001600160a01b03165f908152600460205260409020805460ff19166001179055565b60036020525f9081526040902080546001820180546001600160a01b03909216929161110390611d49565b80601f016020809104026020016040519081016040528092919081815260200182805461112f90611d49565b801561117a5780601f106111515761010080835404028352916020019161117a565b820191905f5260205f20905b81548152906001019060200180831161115d57829003601f168201915b50505050600283015460038401546004850180549495929460ff9092169350906111a390611d49565b80601f01602080910402602001604051908101604052809291908181526020018280546111cf90611d49565b801561121a5780601f106111f15761010080835404028352916020019161121a565b820191905f5260205f20905b8154815290600101906020018083116111fd57829003601f168201915b50505050509080600501805461122f90611d49565b80601f016020809104026020016040519081016040528092919081815260200182805461125b90611d49565b80156112a65780601f1061127d576101008083540402835291602001916112a6565b820191905f5260205f20905b81548152906001019060200180831161128957829003601f168201915b5050505050908060060180546112bb90611d49565b80601f01602080910402602001604051908101604052809291908181526020018280546112e790611d49565b80156113325780601f1061130957610100808354040283529160200191611332565b820191905f5260205f20905b81548152906001019060200180831161131557829003601f168201915b5050506007909301549192505060ff1688565b6113966040518061010001604052805f6001600160a01b03168152602001606081526020015f81526020015f151581526020016060815260200160608152602001606081526020015f151581525090565b5f828152600360208190526040909120015460ff166113ee5760405162461bcd60e51b81526020600482015260146024820152731b185b99081a5cc81b9bdd081d995c9a599a595960621b6044820152606401610501565b5f828152600360209081526040918290208251610100810190935280546001600160a01b03168352600181018054919284019161142a90611d49565b80601f016020809104026020016040519081016040528092919081815260200182805461145690611d49565b80156114a15780601f10611478576101008083540402835291602001916114a1565b820191905f5260205f20905b81548152906001019060200180831161148457829003601f168201915b505050918352505060028201546020820152600382015460ff16151560408201526004820180546060909201916114d790611d49565b80601f016020809104026020016040519081016040528092919081815260200182805461150390611d49565b801561154e5780601f106115255761010080835404028352916020019161154e565b820191905f5260205f20905b81548152906001019060200180831161153157829003601f168201915b5050505050815260200160058201805461156790611d49565b80601f016020809104026020016040519081016040528092919081815260200182805461159390611d49565b80156115de5780601f106115b5576101008083540402835291602001916115de565b820191905f5260205f20905b8154815290600101906020018083116115c157829003601f168201915b505050505081526020016006820180546115f790611d49565b80601f016020809104026020016040519081016040528092919081815260200182805461162390611d49565b801561166e5780601f106116455761010080835404028352916020019161166e565b820191905f5260205f20905b81548152906001019060200180831161165157829003601f168201915b50505091835250506007919091015460ff16151560209091015292915050565b6008602052815f5260405f2081815481106116a7575f80fd5b905f5260205f20015f91509150505481565b335f908152600260205260408120815b600182015481101561170b57838260010182815481106116eb576116eb611f91565b905f5260205f20015403611703575060019392505050565b6001016116c9565b505f9392505050565b335f908152600260205260408120905b60018201548110156117ce578282600101828154811061174657611746611f91565b905f5260205f200154036117c65760018083018054909161176691611fa5565b8154811061177657611776611f91565b905f5260205f20015482600101828154811061179457611794611f91565b5f91825260209091200155600182018054806117b2576117b2611fbe565b600190038181905f5260205f20015f905590555b600101611724565b505050565b828054828255905f5260205f2090810192821561180c579160200282015b8281111561180c5782518255916020019190600101906117f1565b5061181892915061181c565b5090565b5b80821115611818575f815560010161181d565b80356001600160a01b0381168114611846575f80fd5b919050565b5f6020828403121561185b575f80fd5b61186482611830565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261188e575f80fd5b813567ffffffffffffffff8111156118a8576118a861186b565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156118d7576118d761186b565b6040528181528382016020018510156118ee575f80fd5b816020850160208301375f918101602001919091529392505050565b5f806040838503121561191b575f80fd5b61192483611830565b9150602083013567ffffffffffffffff81111561193f575f80fd5b61194b8582860161187f565b9150509250929050565b5f60208284031215611965575f80fd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b604081525f6119ac604083018561196c565b905082151560208301529392505050565b5f805f606084860312156119cf575f80fd5b83359250602084013567ffffffffffffffff8111156119ec575f80fd5b6119f88682870161187f565b92505060408401358015158114611a0d575f80fd5b809150509250925092565b5f8151808452602084019350602083015f5b82811015611a48578151865260209586019590910190600101611a2a565b5093949350505050565b602081525f825160806020840152611a6d60a084018261196c565b90506020840151601f19848303016040850152611a8a8282611a18565b9150506040840151601f19848303016060850152611aa88282611a18565b9150506060840151151560808401528091505092915050565b5f805f60608486031215611ad3575f80fd5b83359250611ae360208501611830565b929592945050506040919091013590565b5f805f60608486031215611b06575f80fd5b83359250602084013567ffffffffffffffff811115611b23575f80fd5b611b2f8682870161187f565b925050604084013567ffffffffffffffff811115611b4b575f80fd5b611b578682870161187f565b9150509250925092565b602080825282518282018190525f918401906040840190835b81811015611b98578351835260209384019390920191600101611b7a565b509095945050505050565b5f805f60608486031215611bb5575f80fd5b8335925060208401359150604084013567ffffffffffffffff811115611b4b575f80fd5b6001600160a01b0389168152610100602082018190525f90611bfd9083018a61196c565b88604084015287151560608401528281036080840152611c1d818861196c565b905082810360a0840152611c31818761196c565b905082810360c0840152611c45818661196c565b91505082151560e08301529998505050505050505050565b60208152611c776020820183516001600160a01b03169052565b5f60208301516101006040840152611c9361012084018261196c565b9050604084015160608401526060840151611cb2608085018215159052565b506080840151838203601f190160a0850152611cce828261196c565b91505060a0840151601f198483030160c0850152611cec828261196c565b91505060c0840151601f198483030160e0850152611d0a828261196c565b91505060e0840151611d2161010085018215159052565b509392505050565b5f8060408385031215611d3a575f80fd5b50508035926020909101359150565b600181811c90821680611d5d57607f821691505b602082108103611d7b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156117ce57805f5260205f20601f840160051c81016020851015611da65750805b601f840160051c820191505b81811015611dc5575f8155600101611db2565b5050505050565b815167ffffffffffffffff811115611de657611de661186b565b611dfa81611df48454611d49565b84611d81565b6020601f821160018114611e2c575f8315611e155750848201515b5f19600385901b1c1916600184901b178455611dc5565b5f84815260208120601f198516915b82811015611e5b5787850151825560209485019460019092019101611e3b565b5084821015611e7857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f8154611e9381611d49565b808552600182168015611ead5760018114611ec957611efd565b60ff1983166020870152602082151560051b8701019350611efd565b845f5260205f205f5b83811015611ef45781546020828a010152600182019150602081019050611ed2565b87016020019450505b50505092915050565b60a081525f611f1860a0830188611e87565b8281036020840152611f2a8188611e87565b90508281036040840152611f3e818761196c565b94151560608401525050608001529392505050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201611f7857611f78611f53565b5060010190565b602081525f611864602083018461196c565b634e487b7160e01b5f52603260045260245ffd5b81810381811115611fb857611fb8611f53565b92915050565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220682753d56e8866c05f98a87f41d60c645cd1aae9124680a104fac3becea8133764736f6c634300081a0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetLandTransactionHistory is a free data retrieval call binding the contract method 0x9de8e3ea.
//
// Solidity: function getLandTransactionHistory(uint256 _landId) view returns(uint256[])
func (_Contract *ContractCaller) GetLandTransactionHistory(opts *bind.CallOpts, _landId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLandTransactionHistory", _landId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLandTransactionHistory is a free data retrieval call binding the contract method 0x9de8e3ea.
//
// Solidity: function getLandTransactionHistory(uint256 _landId) view returns(uint256[])
func (_Contract *ContractSession) GetLandTransactionHistory(_landId *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetLandTransactionHistory(&_Contract.CallOpts, _landId)
}

// GetLandTransactionHistory is a free data retrieval call binding the contract method 0x9de8e3ea.
//
// Solidity: function getLandTransactionHistory(uint256 _landId) view returns(uint256[])
func (_Contract *ContractCallerSession) GetLandTransactionHistory(_landId *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetLandTransactionHistory(&_Contract.CallOpts, _landId)
}

// GetTransactionDetails is a free data retrieval call binding the contract method 0x0fa683d3.
//
// Solidity: function getTransactionDetails(uint256 _transactionId) view returns((uint256,address,address,uint256))
func (_Contract *ContractCaller) GetTransactionDetails(opts *bind.CallOpts, _transactionId *big.Int) (LandTransactionTransaction, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTransactionDetails", _transactionId)

	if err != nil {
		return *new(LandTransactionTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(LandTransactionTransaction)).(*LandTransactionTransaction)

	return out0, err

}

// GetTransactionDetails is a free data retrieval call binding the contract method 0x0fa683d3.
//
// Solidity: function getTransactionDetails(uint256 _transactionId) view returns((uint256,address,address,uint256))
func (_Contract *ContractSession) GetTransactionDetails(_transactionId *big.Int) (LandTransactionTransaction, error) {
	return _Contract.Contract.GetTransactionDetails(&_Contract.CallOpts, _transactionId)
}

// GetTransactionDetails is a free data retrieval call binding the contract method 0x0fa683d3.
//
// Solidity: function getTransactionDetails(uint256 _transactionId) view returns((uint256,address,address,uint256))
func (_Contract *ContractCallerSession) GetTransactionDetails(_transactionId *big.Int) (LandTransactionTransaction, error) {
	return _Contract.Contract.GetTransactionDetails(&_Contract.CallOpts, _transactionId)
}

// Lands is a free data retrieval call binding the contract method 0xe261f1e5.
//
// Solidity: function lands(uint256 ) view returns(address owner, string location, uint256 area, bool isVerified, string detailsHash, string reportHash, string documentsHash, bool isVaild)
func (_Contract *ContractCaller) Lands(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner         common.Address
	Location      string
	Area          *big.Int
	IsVerified    bool
	DetailsHash   string
	ReportHash    string
	DocumentsHash string
	IsVaild       bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lands", arg0)

	outstruct := new(struct {
		Owner         common.Address
		Location      string
		Area          *big.Int
		IsVerified    bool
		DetailsHash   string
		ReportHash    string
		DocumentsHash string
		IsVaild       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Location = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Area = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsVerified = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.DetailsHash = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.ReportHash = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.DocumentsHash = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.IsVaild = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Lands is a free data retrieval call binding the contract method 0xe261f1e5.
//
// Solidity: function lands(uint256 ) view returns(address owner, string location, uint256 area, bool isVerified, string detailsHash, string reportHash, string documentsHash, bool isVaild)
func (_Contract *ContractSession) Lands(arg0 *big.Int) (struct {
	Owner         common.Address
	Location      string
	Area          *big.Int
	IsVerified    bool
	DetailsHash   string
	ReportHash    string
	DocumentsHash string
	IsVaild       bool
}, error) {
	return _Contract.Contract.Lands(&_Contract.CallOpts, arg0)
}

// Lands is a free data retrieval call binding the contract method 0xe261f1e5.
//
// Solidity: function lands(uint256 ) view returns(address owner, string location, uint256 area, bool isVerified, string detailsHash, string reportHash, string documentsHash, bool isVaild)
func (_Contract *ContractCallerSession) Lands(arg0 *big.Int) (struct {
	Owner         common.Address
	Location      string
	Area          *big.Int
	IsVerified    bool
	DetailsHash   string
	ReportHash    string
	DocumentsHash string
	IsVaild       bool
}, error) {
	return _Contract.Contract.Lands(&_Contract.CallOpts, arg0)
}

// LandsTransactionIdList is a free data retrieval call binding the contract method 0xf5b73cbc.
//
// Solidity: function landsTransactionIdList(uint256 , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) LandsTransactionIdList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "landsTransactionIdList", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LandsTransactionIdList is a free data retrieval call binding the contract method 0xf5b73cbc.
//
// Solidity: function landsTransactionIdList(uint256 , uint256 ) view returns(uint256)
func (_Contract *ContractSession) LandsTransactionIdList(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.LandsTransactionIdList(&_Contract.CallOpts, arg0, arg1)
}

// LandsTransactionIdList is a free data retrieval call binding the contract method 0xf5b73cbc.
//
// Solidity: function landsTransactionIdList(uint256 , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) LandsTransactionIdList(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.LandsTransactionIdList(&_Contract.CallOpts, arg0, arg1)
}

// Notaries is a free data retrieval call binding the contract method 0x00e168f0.
//
// Solidity: function notaries(address ) view returns(bool)
func (_Contract *ContractCaller) Notaries(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "notaries", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Notaries is a free data retrieval call binding the contract method 0x00e168f0.
//
// Solidity: function notaries(address ) view returns(bool)
func (_Contract *ContractSession) Notaries(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Notaries(&_Contract.CallOpts, arg0)
}

// Notaries is a free data retrieval call binding the contract method 0x00e168f0.
//
// Solidity: function notaries(address ) view returns(bool)
func (_Contract *ContractCallerSession) Notaries(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Notaries(&_Contract.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contract *ContractCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contract *ContractSession) Oracle() (common.Address, error) {
	return _Contract.Contract.Oracle(&_Contract.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contract *ContractCallerSession) Oracle() (common.Address, error) {
	return _Contract.Contract.Oracle(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// QueryLand is a free data retrieval call binding the contract method 0xe7567696.
//
// Solidity: function queryLand(uint256 landId) view returns((address,string,uint256,bool,string,string,string,bool))
func (_Contract *ContractCaller) QueryLand(opts *bind.CallOpts, landId *big.Int) (LandRegistryLand, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "queryLand", landId)

	if err != nil {
		return *new(LandRegistryLand), err
	}

	out0 := *abi.ConvertType(out[0], new(LandRegistryLand)).(*LandRegistryLand)

	return out0, err

}

// QueryLand is a free data retrieval call binding the contract method 0xe7567696.
//
// Solidity: function queryLand(uint256 landId) view returns((address,string,uint256,bool,string,string,string,bool))
func (_Contract *ContractSession) QueryLand(landId *big.Int) (LandRegistryLand, error) {
	return _Contract.Contract.QueryLand(&_Contract.CallOpts, landId)
}

// QueryLand is a free data retrieval call binding the contract method 0xe7567696.
//
// Solidity: function queryLand(uint256 landId) view returns((address,string,uint256,bool,string,string,string,bool))
func (_Contract *ContractCallerSession) QueryLand(landId *big.Int) (LandRegistryLand, error) {
	return _Contract.Contract.QueryLand(&_Contract.CallOpts, landId)
}

// QueryUserInfo is a free data retrieval call binding the contract method 0x5307d396.
//
// Solidity: function queryUserInfo() view returns((string,uint256[],uint256[],bool))
func (_Contract *ContractCaller) QueryUserInfo(opts *bind.CallOpts) (LandRegistryUser, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "queryUserInfo")

	if err != nil {
		return *new(LandRegistryUser), err
	}

	out0 := *abi.ConvertType(out[0], new(LandRegistryUser)).(*LandRegistryUser)

	return out0, err

}

// QueryUserInfo is a free data retrieval call binding the contract method 0x5307d396.
//
// Solidity: function queryUserInfo() view returns((string,uint256[],uint256[],bool))
func (_Contract *ContractSession) QueryUserInfo() (LandRegistryUser, error) {
	return _Contract.Contract.QueryUserInfo(&_Contract.CallOpts)
}

// QueryUserInfo is a free data retrieval call binding the contract method 0x5307d396.
//
// Solidity: function queryUserInfo() view returns((string,uint256[],uint256[],bool))
func (_Contract *ContractCallerSession) QueryUserInfo() (LandRegistryUser, error) {
	return _Contract.Contract.QueryUserInfo(&_Contract.CallOpts)
}

// Surveyors is a free data retrieval call binding the contract method 0x87b6963f.
//
// Solidity: function surveyors(address ) view returns(bool)
func (_Contract *ContractCaller) Surveyors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "surveyors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Surveyors is a free data retrieval call binding the contract method 0x87b6963f.
//
// Solidity: function surveyors(address ) view returns(bool)
func (_Contract *ContractSession) Surveyors(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Surveyors(&_Contract.CallOpts, arg0)
}

// Surveyors is a free data retrieval call binding the contract method 0x87b6963f.
//
// Solidity: function surveyors(address ) view returns(bool)
func (_Contract *ContractCallerSession) Surveyors(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Surveyors(&_Contract.CallOpts, arg0)
}

// TransactionCounter is a free data retrieval call binding the contract method 0x14e887e8.
//
// Solidity: function transactionCounter() view returns(uint256)
func (_Contract *ContractCaller) TransactionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "transactionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCounter is a free data retrieval call binding the contract method 0x14e887e8.
//
// Solidity: function transactionCounter() view returns(uint256)
func (_Contract *ContractSession) TransactionCounter() (*big.Int, error) {
	return _Contract.Contract.TransactionCounter(&_Contract.CallOpts)
}

// TransactionCounter is a free data retrieval call binding the contract method 0x14e887e8.
//
// Solidity: function transactionCounter() view returns(uint256)
func (_Contract *ContractCallerSession) TransactionCounter() (*big.Int, error) {
	return _Contract.Contract.TransactionCounter(&_Contract.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(uint256 landId, address from, address to, uint256 timestamp)
func (_Contract *ContractCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LandId    *big.Int
	From      common.Address
	To        common.Address
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		LandId    *big.Int
		From      common.Address
		To        common.Address
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LandId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.From = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(uint256 landId, address from, address to, uint256 timestamp)
func (_Contract *ContractSession) Transactions(arg0 *big.Int) (struct {
	LandId    *big.Int
	From      common.Address
	To        common.Address
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.Transactions(&_Contract.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(uint256 landId, address from, address to, uint256 timestamp)
func (_Contract *ContractCallerSession) Transactions(arg0 *big.Int) (struct {
	LandId    *big.Int
	From      common.Address
	To        common.Address
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.Transactions(&_Contract.CallOpts, arg0)
}

// UserMapping is a free data retrieval call binding the contract method 0x118b9535.
//
// Solidity: function userMapping(address ) view returns(string username, bool isVaild)
func (_Contract *ContractCaller) UserMapping(opts *bind.CallOpts, arg0 common.Address) (struct {
	Username string
	IsVaild  bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userMapping", arg0)

	outstruct := new(struct {
		Username string
		IsVaild  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Username = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.IsVaild = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// UserMapping is a free data retrieval call binding the contract method 0x118b9535.
//
// Solidity: function userMapping(address ) view returns(string username, bool isVaild)
func (_Contract *ContractSession) UserMapping(arg0 common.Address) (struct {
	Username string
	IsVaild  bool
}, error) {
	return _Contract.Contract.UserMapping(&_Contract.CallOpts, arg0)
}

// UserMapping is a free data retrieval call binding the contract method 0x118b9535.
//
// Solidity: function userMapping(address ) view returns(string username, bool isVaild)
func (_Contract *ContractCallerSession) UserMapping(arg0 common.Address) (struct {
	Username string
	IsVaild  bool
}, error) {
	return _Contract.Contract.UserMapping(&_Contract.CallOpts, arg0)
}

// LandSurveyingArea is a paid mutator transaction binding the contract method 0xa45c2c4c.
//
// Solidity: function LandSurveyingArea(uint256 _landId, uint256 _area, string _reportHash) returns()
func (_Contract *ContractTransactor) LandSurveyingArea(opts *bind.TransactOpts, _landId *big.Int, _area *big.Int, _reportHash string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "LandSurveyingArea", _landId, _area, _reportHash)
}

// LandSurveyingArea is a paid mutator transaction binding the contract method 0xa45c2c4c.
//
// Solidity: function LandSurveyingArea(uint256 _landId, uint256 _area, string _reportHash) returns()
func (_Contract *ContractSession) LandSurveyingArea(_landId *big.Int, _area *big.Int, _reportHash string) (*types.Transaction, error) {
	return _Contract.Contract.LandSurveyingArea(&_Contract.TransactOpts, _landId, _area, _reportHash)
}

// LandSurveyingArea is a paid mutator transaction binding the contract method 0xa45c2c4c.
//
// Solidity: function LandSurveyingArea(uint256 _landId, uint256 _area, string _reportHash) returns()
func (_Contract *ContractTransactorSession) LandSurveyingArea(_landId *big.Int, _area *big.Int, _reportHash string) (*types.Transaction, error) {
	return _Contract.Contract.LandSurveyingArea(&_Contract.TransactOpts, _landId, _area, _reportHash)
}

// AddNotary is a paid mutator transaction binding the contract method 0x0588565f.
//
// Solidity: function addNotary(address _notary) returns()
func (_Contract *ContractTransactor) AddNotary(opts *bind.TransactOpts, _notary common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addNotary", _notary)
}

// AddNotary is a paid mutator transaction binding the contract method 0x0588565f.
//
// Solidity: function addNotary(address _notary) returns()
func (_Contract *ContractSession) AddNotary(_notary common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddNotary(&_Contract.TransactOpts, _notary)
}

// AddNotary is a paid mutator transaction binding the contract method 0x0588565f.
//
// Solidity: function addNotary(address _notary) returns()
func (_Contract *ContractTransactorSession) AddNotary(_notary common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddNotary(&_Contract.TransactOpts, _notary)
}

// AddSurveyor is a paid mutator transaction binding the contract method 0xad812a52.
//
// Solidity: function addSurveyor(address _surveyor) returns()
func (_Contract *ContractTransactor) AddSurveyor(opts *bind.TransactOpts, _surveyor common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addSurveyor", _surveyor)
}

// AddSurveyor is a paid mutator transaction binding the contract method 0xad812a52.
//
// Solidity: function addSurveyor(address _surveyor) returns()
func (_Contract *ContractSession) AddSurveyor(_surveyor common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddSurveyor(&_Contract.TransactOpts, _surveyor)
}

// AddSurveyor is a paid mutator transaction binding the contract method 0xad812a52.
//
// Solidity: function addSurveyor(address _surveyor) returns()
func (_Contract *ContractTransactorSession) AddSurveyor(_surveyor common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddSurveyor(&_Contract.TransactOpts, _surveyor)
}

// RegisterLand is a paid mutator transaction binding the contract method 0x9437922a.
//
// Solidity: function registerLand(uint256 _landId, string _location, string _detailsHash) returns()
func (_Contract *ContractTransactor) RegisterLand(opts *bind.TransactOpts, _landId *big.Int, _location string, _detailsHash string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerLand", _landId, _location, _detailsHash)
}

// RegisterLand is a paid mutator transaction binding the contract method 0x9437922a.
//
// Solidity: function registerLand(uint256 _landId, string _location, string _detailsHash) returns()
func (_Contract *ContractSession) RegisterLand(_landId *big.Int, _location string, _detailsHash string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterLand(&_Contract.TransactOpts, _landId, _location, _detailsHash)
}

// RegisterLand is a paid mutator transaction binding the contract method 0x9437922a.
//
// Solidity: function registerLand(uint256 _landId, string _location, string _detailsHash) returns()
func (_Contract *ContractTransactorSession) RegisterLand(_landId *big.Int, _location string, _detailsHash string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterLand(&_Contract.TransactOpts, _landId, _location, _detailsHash)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x0ea126f9.
//
// Solidity: function registerUser(address _userAddress, string _userName) returns()
func (_Contract *ContractTransactor) RegisterUser(opts *bind.TransactOpts, _userAddress common.Address, _userName string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerUser", _userAddress, _userName)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x0ea126f9.
//
// Solidity: function registerUser(address _userAddress, string _userName) returns()
func (_Contract *ContractSession) RegisterUser(_userAddress common.Address, _userName string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterUser(&_Contract.TransactOpts, _userAddress, _userName)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x0ea126f9.
//
// Solidity: function registerUser(address _userAddress, string _userName) returns()
func (_Contract *ContractTransactorSession) RegisterUser(_userAddress common.Address, _userName string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterUser(&_Contract.TransactOpts, _userAddress, _userName)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_Contract *ContractSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwner(&_Contract.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwner(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x70435c99.
//
// Solidity: function transferOwnership(uint256 _landId, address _to, uint256 _transactionId) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, _landId *big.Int, _to common.Address, _transactionId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", _landId, _to, _transactionId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x70435c99.
//
// Solidity: function transferOwnership(uint256 _landId, address _to, uint256 _transactionId) returns()
func (_Contract *ContractSession) TransferOwnership(_landId *big.Int, _to common.Address, _transactionId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, _landId, _to, _transactionId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x70435c99.
//
// Solidity: function transferOwnership(uint256 _landId, address _to, uint256 _transactionId) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(_landId *big.Int, _to common.Address, _transactionId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, _landId, _to, _transactionId)
}

// VerifyLand is a paid mutator transaction binding the contract method 0x49695ecc.
//
// Solidity: function verifyLand(uint256 _landId, string _documentsHash, bool _isVerified) returns()
func (_Contract *ContractTransactor) VerifyLand(opts *bind.TransactOpts, _landId *big.Int, _documentsHash string, _isVerified bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifyLand", _landId, _documentsHash, _isVerified)
}

// VerifyLand is a paid mutator transaction binding the contract method 0x49695ecc.
//
// Solidity: function verifyLand(uint256 _landId, string _documentsHash, bool _isVerified) returns()
func (_Contract *ContractSession) VerifyLand(_landId *big.Int, _documentsHash string, _isVerified bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyLand(&_Contract.TransactOpts, _landId, _documentsHash, _isVerified)
}

// VerifyLand is a paid mutator transaction binding the contract method 0x49695ecc.
//
// Solidity: function verifyLand(uint256 _landId, string _documentsHash, bool _isVerified) returns()
func (_Contract *ContractTransactorSession) VerifyLand(_landId *big.Int, _documentsHash string, _isVerified bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyLand(&_Contract.TransactOpts, _landId, _documentsHash, _isVerified)
}

// ContractLandRegisteredIterator is returned from FilterLandRegistered and is used to iterate over the raw logs and unpacked data for LandRegistered events raised by the Contract contract.
type ContractLandRegisteredIterator struct {
	Event *ContractLandRegistered // Event containing the contract specifics and raw log

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
func (it *ContractLandRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLandRegistered)
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
		it.Event = new(ContractLandRegistered)
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
func (it *ContractLandRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLandRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLandRegistered represents a LandRegistered event raised by the Contract contract.
type ContractLandRegistered struct {
	LandId   *big.Int
	Owner    common.Address
	Location string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLandRegistered is a free log retrieval operation binding the contract event 0x9c4ad74928fec1ee86181680df0e76bf575b79ea2a0347f7ed270884bdd26ad7.
//
// Solidity: event LandRegistered(uint256 indexed landId, address indexed owner, string location)
func (_Contract *ContractFilterer) FilterLandRegistered(opts *bind.FilterOpts, landId []*big.Int, owner []common.Address) (*ContractLandRegisteredIterator, error) {

	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LandRegistered", landIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractLandRegisteredIterator{contract: _Contract.contract, event: "LandRegistered", logs: logs, sub: sub}, nil
}

// WatchLandRegistered is a free log subscription operation binding the contract event 0x9c4ad74928fec1ee86181680df0e76bf575b79ea2a0347f7ed270884bdd26ad7.
//
// Solidity: event LandRegistered(uint256 indexed landId, address indexed owner, string location)
func (_Contract *ContractFilterer) WatchLandRegistered(opts *bind.WatchOpts, sink chan<- *ContractLandRegistered, landId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LandRegistered", landIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLandRegistered)
				if err := _Contract.contract.UnpackLog(event, "LandRegistered", log); err != nil {
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

// ParseLandRegistered is a log parse operation binding the contract event 0x9c4ad74928fec1ee86181680df0e76bf575b79ea2a0347f7ed270884bdd26ad7.
//
// Solidity: event LandRegistered(uint256 indexed landId, address indexed owner, string location)
func (_Contract *ContractFilterer) ParseLandRegistered(log types.Log) (*ContractLandRegistered, error) {
	event := new(ContractLandRegistered)
	if err := _Contract.contract.UnpackLog(event, "LandRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLandSurveyingIterator is returned from FilterLandSurveying and is used to iterate over the raw logs and unpacked data for LandSurveying events raised by the Contract contract.
type ContractLandSurveyingIterator struct {
	Event *ContractLandSurveying // Event containing the contract specifics and raw log

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
func (it *ContractLandSurveyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLandSurveying)
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
		it.Event = new(ContractLandSurveying)
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
func (it *ContractLandSurveyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLandSurveyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLandSurveying represents a LandSurveying event raised by the Contract contract.
type ContractLandSurveying struct {
	LandId           *big.Int
	Area             *big.Int
	SurveyorsAddress common.Address
	Timestam         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLandSurveying is a free log retrieval operation binding the contract event 0x9f207f53ad48e112fa882ba14907e3301685086486bac5d73bab867a7c9a3c11.
//
// Solidity: event LandSurveying(uint256 indexed landId, uint256 area, address indexed surveyorsAddress, uint256 timestam)
func (_Contract *ContractFilterer) FilterLandSurveying(opts *bind.FilterOpts, landId []*big.Int, surveyorsAddress []common.Address) (*ContractLandSurveyingIterator, error) {

	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}

	var surveyorsAddressRule []interface{}
	for _, surveyorsAddressItem := range surveyorsAddress {
		surveyorsAddressRule = append(surveyorsAddressRule, surveyorsAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LandSurveying", landIdRule, surveyorsAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractLandSurveyingIterator{contract: _Contract.contract, event: "LandSurveying", logs: logs, sub: sub}, nil
}

// WatchLandSurveying is a free log subscription operation binding the contract event 0x9f207f53ad48e112fa882ba14907e3301685086486bac5d73bab867a7c9a3c11.
//
// Solidity: event LandSurveying(uint256 indexed landId, uint256 area, address indexed surveyorsAddress, uint256 timestam)
func (_Contract *ContractFilterer) WatchLandSurveying(opts *bind.WatchOpts, sink chan<- *ContractLandSurveying, landId []*big.Int, surveyorsAddress []common.Address) (event.Subscription, error) {

	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}

	var surveyorsAddressRule []interface{}
	for _, surveyorsAddressItem := range surveyorsAddress {
		surveyorsAddressRule = append(surveyorsAddressRule, surveyorsAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LandSurveying", landIdRule, surveyorsAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLandSurveying)
				if err := _Contract.contract.UnpackLog(event, "LandSurveying", log); err != nil {
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

// ParseLandSurveying is a log parse operation binding the contract event 0x9f207f53ad48e112fa882ba14907e3301685086486bac5d73bab867a7c9a3c11.
//
// Solidity: event LandSurveying(uint256 indexed landId, uint256 area, address indexed surveyorsAddress, uint256 timestam)
func (_Contract *ContractFilterer) ParseLandSurveying(log types.Log) (*ContractLandSurveying, error) {
	event := new(ContractLandSurveying)
	if err := _Contract.contract.UnpackLog(event, "LandSurveying", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLandVerifiedIterator is returned from FilterLandVerified and is used to iterate over the raw logs and unpacked data for LandVerified events raised by the Contract contract.
type ContractLandVerifiedIterator struct {
	Event *ContractLandVerified // Event containing the contract specifics and raw log

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
func (it *ContractLandVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLandVerified)
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
		it.Event = new(ContractLandVerified)
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
func (it *ContractLandVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLandVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLandVerified represents a LandVerified event raised by the Contract contract.
type ContractLandVerified struct {
	LandId          *big.Int
	DetailsHash     string
	ReportHash      string
	DocumentsHash   string
	IsVerified      bool
	NotariesAddress common.Address
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLandVerified is a free log retrieval operation binding the contract event 0x378e3aa8b6938b8ea84d87fd3355efdb97062e22ca268d202cb08acb23e1f0d9.
//
// Solidity: event LandVerified(uint256 indexed landId, string detailsHash, string reportHash, string documentsHash, bool isVerified, address indexed notariesAddress, uint256 timestamp)
func (_Contract *ContractFilterer) FilterLandVerified(opts *bind.FilterOpts, landId []*big.Int, notariesAddress []common.Address) (*ContractLandVerifiedIterator, error) {

	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}

	var notariesAddressRule []interface{}
	for _, notariesAddressItem := range notariesAddress {
		notariesAddressRule = append(notariesAddressRule, notariesAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LandVerified", landIdRule, notariesAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractLandVerifiedIterator{contract: _Contract.contract, event: "LandVerified", logs: logs, sub: sub}, nil
}

// WatchLandVerified is a free log subscription operation binding the contract event 0x378e3aa8b6938b8ea84d87fd3355efdb97062e22ca268d202cb08acb23e1f0d9.
//
// Solidity: event LandVerified(uint256 indexed landId, string detailsHash, string reportHash, string documentsHash, bool isVerified, address indexed notariesAddress, uint256 timestamp)
func (_Contract *ContractFilterer) WatchLandVerified(opts *bind.WatchOpts, sink chan<- *ContractLandVerified, landId []*big.Int, notariesAddress []common.Address) (event.Subscription, error) {

	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}

	var notariesAddressRule []interface{}
	for _, notariesAddressItem := range notariesAddress {
		notariesAddressRule = append(notariesAddressRule, notariesAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LandVerified", landIdRule, notariesAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLandVerified)
				if err := _Contract.contract.UnpackLog(event, "LandVerified", log); err != nil {
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

// ParseLandVerified is a log parse operation binding the contract event 0x378e3aa8b6938b8ea84d87fd3355efdb97062e22ca268d202cb08acb23e1f0d9.
//
// Solidity: event LandVerified(uint256 indexed landId, string detailsHash, string reportHash, string documentsHash, bool isVerified, address indexed notariesAddress, uint256 timestamp)
func (_Contract *ContractFilterer) ParseLandVerified(log types.Log) (*ContractLandVerified, error) {
	event := new(ContractLandVerified)
	if err := _Contract.contract.UnpackLog(event, "LandVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransactionRecordedIterator is returned from FilterTransactionRecorded and is used to iterate over the raw logs and unpacked data for TransactionRecorded events raised by the Contract contract.
type ContractTransactionRecordedIterator struct {
	Event *ContractTransactionRecorded // Event containing the contract specifics and raw log

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
func (it *ContractTransactionRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransactionRecorded)
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
		it.Event = new(ContractTransactionRecorded)
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
func (it *ContractTransactionRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransactionRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransactionRecorded represents a TransactionRecorded event raised by the Contract contract.
type ContractTransactionRecorded struct {
	TransactionId *big.Int
	LandId        *big.Int
	From          common.Address
	To            common.Address
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionRecorded is a free log retrieval operation binding the contract event 0x4bb86a8504e257f90fe72dcaa3a3d0dfb04b72836b57aa6882b7bc2a1b82a4d1.
//
// Solidity: event TransactionRecorded(uint256 indexed transactionId, uint256 indexed landId, address indexed from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) FilterTransactionRecorded(opts *bind.FilterOpts, transactionId []*big.Int, landId []*big.Int, from []common.Address) (*ContractTransactionRecordedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TransactionRecorded", transactionIdRule, landIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransactionRecordedIterator{contract: _Contract.contract, event: "TransactionRecorded", logs: logs, sub: sub}, nil
}

// WatchTransactionRecorded is a free log subscription operation binding the contract event 0x4bb86a8504e257f90fe72dcaa3a3d0dfb04b72836b57aa6882b7bc2a1b82a4d1.
//
// Solidity: event TransactionRecorded(uint256 indexed transactionId, uint256 indexed landId, address indexed from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) WatchTransactionRecorded(opts *bind.WatchOpts, sink chan<- *ContractTransactionRecorded, transactionId []*big.Int, landId []*big.Int, from []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var landIdRule []interface{}
	for _, landIdItem := range landId {
		landIdRule = append(landIdRule, landIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TransactionRecorded", transactionIdRule, landIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransactionRecorded)
				if err := _Contract.contract.UnpackLog(event, "TransactionRecorded", log); err != nil {
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

// ParseTransactionRecorded is a log parse operation binding the contract event 0x4bb86a8504e257f90fe72dcaa3a3d0dfb04b72836b57aa6882b7bc2a1b82a4d1.
//
// Solidity: event TransactionRecorded(uint256 indexed transactionId, uint256 indexed landId, address indexed from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) ParseTransactionRecorded(log types.Log) (*ContractTransactionRecorded, error) {
	event := new(ContractTransactionRecorded)
	if err := _Contract.contract.UnpackLog(event, "TransactionRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
