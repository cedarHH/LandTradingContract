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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"LandRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"area\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"surveyorsAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestam\",\"type\":\"uint256\"}],\"name\":\"LandSurveying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"detailsHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reportHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"documentsHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notariesAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LandVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_area\",\"type\":\"uint256\"}],\"name\":\"LandSurveyingArea\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"addNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_surveyor\",\"type\":\"address\"}],\"name\":\"addSurveyor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"}],\"name\":\"getLandTransactionHistory\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"getTransactionDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLandTransaction.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lands\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"area\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"detailsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reportHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"documentsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"landsTransactionIdList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"notaries\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"}],\"name\":\"queryLand\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"area\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"detailsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reportHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"documentsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"internalType\":\"structLandRegistry.Land\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"landIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"transactionIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"internalType\":\"structLandRegistry.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"registerLand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_userName\",\"type\":\"string\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"surveyors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"landId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMapping\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_landId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_detailsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_reportHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_documentsHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isVerified\",\"type\":\"bool\"}],\"name\":\"verifyLand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f6005553480156012575f80fd5b505f80546001600160a01b03191633179055611eb5806100315f395ff3fe608060405234801561000f575f80fd5b5060043610610126575f3560e01c806387b6963f116100a9578063c35ce0541161006e578063c35ce054146103be578063e261f1e5146103d1578063e7567696146103f8578063eb4c277414610418578063f5b73cbc1461042b575f80fd5b806387b6963f146102ce5780638da5cb5b146102f05780639ace38c21461031a5780639de8e3ea1461038b578063ad812a52146103ab575f80fd5b806314e887e8116100ef57806314e887e81461026957806347b24c08146102805780634fb2e45d146102935780635307d396146102a657806370435c99146102bb575f80fd5b8062e168f01461012a5780630588565f146101615780630ea126f9146101765780630fa683d314610189578063118b953514610248575b5f80fd5b61014c610138366004611797565b60046020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61017461016f366004611797565b61043e565b005b610174610184366004611856565b610476565b6102076101973660046118a1565b60408051608080820183525f8083526020808401829052838501829052606093840182905294815260068552839020835191820184528054825260018101546001600160a01b03908116958301959095526002810154909416928101929092526003909201549181019190915290565b60408051825181526020808401516001600160a01b039081169183019190915283830151169181019190915260609182015191810191909152608001610158565b61025b610256366004611797565b61059a565b6040516101589291906118e6565b61027260085481565b604051908152602001610158565b61017461028e366004611909565b61063e565b6101746102a1366004611797565b610863565b6102ae6108e3565b6040516101589190611971565b6101746102c93660046119e0565b610a79565b61014c6102dc366004611797565b60036020525f908152604090205460ff1681565b5f54610302906001600160a01b031681565b6040516001600160a01b039091168152602001610158565b61035b6103283660046118a1565b60066020525f9081526040902080546001820154600283015460039093015491926001600160a01b039182169291169084565b60405161015894939291909384526001600160a01b03928316602085015291166040830152606082015260800190565b61039e6103993660046118a1565b610d30565b6040516101589190611a13565b6101746103b9366004611797565b610d8f565b6101746103cc366004611a55565b610dc7565b6103e46103df3660046118a1565b610ec9565b604051610158989796959493929190611a75565b61040b6104063660046118a1565b611136565b6040516101589190611af9565b610174610426366004611bc5565b61147e565b610272610439366004611a55565b6115da565b5f546001600160a01b03163314610453575f80fd5b6001600160a01b03165f908152600460205260409020805460ff19166001179055565b5f546001600160a01b0316331461048b575f80fd5b6001600160a01b0382165f9081526001602052604090206003015460ff16156104ec5760405162461bcd60e51b815260206004820152600e60248201526d757365722069732065786973747360901b60448201526064015b60405180910390fd5b6040805160808101825282815281515f808252602080830185528084019290925283518181528083018552838501526001606084018190526001600160a01b038716825290915291909120815181906105459082611cf8565b50602082810151805161055e926001850192019061171f565b506040820151805161057a91600284019160209091019061171f565b50606091909101516003909101805460ff19169115159190911790555050565b60016020525f90815260409020805481906105b490611c75565b80601f01602080910402602001604051908101604052809291908181526020018280546105e090611c75565b801561062b5780601f106106025761010080835404028352916020019161062b565b820191905f5260205f20905b81548152906001019060200180831161060e57829003601f168201915b5050506003909301549192505060ff1682565b335f9081526001602052604090206003015460ff1661068d5760405162461bcd60e51b815260206004820152600b60248201526a34b9903737ba103ab9b2b960a91b60448201526064016104e3565b5f8281526002602052604090206007015460ff16156106df5760405162461bcd60e51b815260206004820152600e60248201526d6c616e642069732065786973747360901b60448201526064016104e3565b604080516101008101825233815260208082018481525f838501819052606084018190528451808401865281815260808501528451808401865281815260a08501528451808401865281815260c0850152600160e08501819052878252600290935293909320825181546001600160a01b0319166001600160a01b039091161781559251919291908201906107749082611cf8565b5060408201516002820155606082015160038201805460ff1916911515919091179055608082015160048201906107ab9082611cf8565b5060a082015160058201906107c09082611cf8565b5060c082015160068201906107d59082611cf8565b5060e091909101516007909101805460ff1916911515919091179055335f81815260016020818152604080842083018054938401815584529220018490555183907f9c4ad74928fec1ee86181680df0e76bf575b79ea2a0347f7ed270884bdd26ad790610843908590611db3565b60405180910390a360058054905f61085a83611dd9565b91905055505050565b5f546001600160a01b03163314610878575f80fd5b6001600160a01b03811661088a575f80fd5b5f80546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35f80546001600160a01b0319166001600160a01b0392909216919091179055565b61090f60405180608001604052806060815260200160608152602001606081526020015f151581525090565b335f908152600160205260409081902081516080810190925280548290829061093790611c75565b80601f016020809104026020016040519081016040528092919081815260200182805461096390611c75565b80156109ae5780601f10610985576101008083540402835291602001916109ae565b820191905f5260205f20905b81548152906001019060200180831161099157829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610a0457602002820191905f5260205f20905b8154815260200190600101908083116109f0575b5050505050815260200160028201805480602002602001604051908101604052809291908181526020018280548015610a5a57602002820191905f5260205f20905b815481526020019060010190808311610a46575b50505091835250506003919091015460ff161515602090910152919050565b82610a8381611605565b610abf5760405162461bcd60e51b815260206004820152600d60248201526c1b9bdd081e5bdd5c881b185b99609a1b60448201526064016104e3565b6001600160a01b0383163303610b175760405162461bcd60e51b815260206004820181905260248201527f43616e6e6f74206265207472616e7366657272656420746f206f6e6573656c6660448201526064016104e3565b5f8481526002602052604090206003015460ff16610b775760405162461bcd60e51b815260206004820152601960248201527f74686973206c616e64206973206e6f742056657269666965640000000000000060448201526064016104e3565b6001600160a01b0383165f9081526001602052604090206003015460ff16610bd85760405162461bcd60e51b8152602060048201526014602482015273746f55736572206973206e6f742065786973747360601b60448201526064016104e3565b6001600160a01b0383165f9081526001602081815260408320820180549283018155835290912001849055610c0c84611660565b6001600160a01b038381165f8181526001602081815260408084206002908101805480860182559086528386200189905533808652828620820180548087018255908752848720018a905582516080810184528c815280850182815281850189815242606084018181528e8b5260068952998790209351845591519783018054988c166001600160a01b0319998a1617905551938201805494909a16939096169290921790975593516003909401939093558251938452830152869185917f4bb86a8504e257f90fe72dcaa3a3d0dfb04b72836b57aa6882b7bc2a1b82a4d1910160405180910390a460088054905f610d0483611dd9565b9091555050505f9283526007602090815260408420805460018101825590855293209092019190915550565b5f81815260076020908152604091829020805483518184028101840190945280845260609392830182828015610d8357602002820191905f5260205f20905b815481526020019060010190808311610d6f575b50505050509050919050565b5f546001600160a01b03163314610da4575f80fd5b6001600160a01b03165f908152600360205260409020805460ff19166001179055565b335f9081526003602052604090205460ff16610e165760405162461bcd60e51b815260206004820152600e60248201526d2737ba10309039bab93b32bcb7b960911b60448201526064016104e3565b5f8281526002602052604090206007015460ff16610e6b5760405162461bcd60e51b81526020600482015260126024820152716c616e64206973206e6f742065786973747360701b60448201526064016104e3565b5f828152600260208190526040918290200182905551339083907f9f207f53ad48e112fa882ba14907e3301685086486bac5d73bab867a7c9a3c1190610ebd9085904290918252602082015260400190565b60405180910390a35050565b60026020525f9081526040902080546001820180546001600160a01b039092169291610ef490611c75565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2090611c75565b8015610f6b5780601f10610f4257610100808354040283529160200191610f6b565b820191905f5260205f20905b815481529060010190602001808311610f4e57829003601f168201915b50505050600283015460038401546004850180549495929460ff909216935090610f9490611c75565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc090611c75565b801561100b5780601f10610fe25761010080835404028352916020019161100b565b820191905f5260205f20905b815481529060010190602001808311610fee57829003601f168201915b50505050509080600501805461102090611c75565b80601f016020809104026020016040519081016040528092919081815260200182805461104c90611c75565b80156110975780601f1061106e57610100808354040283529160200191611097565b820191905f5260205f20905b81548152906001019060200180831161107a57829003601f168201915b5050505050908060060180546110ac90611c75565b80601f01602080910402602001604051908101604052809291908181526020018280546110d890611c75565b80156111235780601f106110fa57610100808354040283529160200191611123565b820191905f5260205f20905b81548152906001019060200180831161110657829003601f168201915b5050506007909301549192505060ff1688565b6111876040518061010001604052805f6001600160a01b03168152602001606081526020015f81526020015f151581526020016060815260200160608152602001606081526020015f151581525090565b5f8281526002602052604090206003015460ff166111de5760405162461bcd60e51b81526020600482015260146024820152731b185b99081a5cc81b9bdd081d995c9a599a595960621b60448201526064016104e3565b5f828152600260209081526040918290208251610100810190935280546001600160a01b03168352600181018054919284019161121a90611c75565b80601f016020809104026020016040519081016040528092919081815260200182805461124690611c75565b80156112915780601f1061126857610100808354040283529160200191611291565b820191905f5260205f20905b81548152906001019060200180831161127457829003601f168201915b505050918352505060028201546020820152600382015460ff16151560408201526004820180546060909201916112c790611c75565b80601f01602080910402602001604051908101604052809291908181526020018280546112f390611c75565b801561133e5780601f106113155761010080835404028352916020019161133e565b820191905f5260205f20905b81548152906001019060200180831161132157829003601f168201915b5050505050815260200160058201805461135790611c75565b80601f016020809104026020016040519081016040528092919081815260200182805461138390611c75565b80156113ce5780601f106113a5576101008083540402835291602001916113ce565b820191905f5260205f20905b8154815290600101906020018083116113b157829003601f168201915b505050505081526020016006820180546113e790611c75565b80601f016020809104026020016040519081016040528092919081815260200182805461141390611c75565b801561145e5780601f106114355761010080835404028352916020019161145e565b820191905f5260205f20905b81548152906001019060200180831161144157829003601f168201915b50505091835250506007919091015460ff16151560209091015292915050565b335f9081526004602052604090205460ff166114cd5760405162461bcd60e51b815260206004820152600e60248201526d4e6f742061206e6f74617269657360901b60448201526064016104e3565b5f8581526002602052604090206007015460ff166115225760405162461bcd60e51b81526020600482015260126024820152716c616e64206973206e6f742065786973747360701b60448201526064016104e3565b5f85815260026020526040902060040161153c8582611cf8565b505f8581526002602052604090206005016115578482611cf8565b505f8581526002602052604090206006016115728382611cf8565b505f8581526002602052604090819020600301805460ff191683151517905551339086907f378e3aa8b6938b8ea84d87fd3355efdb97062e22ca268d202cb08acb23e1f0d9906115cb9088908890889088904290611df1565b60405180910390a35050505050565b6007602052815f5260405f2081815481106115f3575f80fd5b905f5260205f20015f91509150505481565b335f908152600160205260408120815b6001820154811015611657578382600101828154811061163757611637611e3e565b905f5260205f2001540361164f575060019392505050565b600101611615565b505f9392505050565b335f908152600160205260408120905b600182015481101561171a578282600101828154811061169257611692611e3e565b905f5260205f20015403611712576001808301805490916116b291611e52565b815481106116c2576116c2611e3e565b905f5260205f2001548260010182815481106116e0576116e0611e3e565b5f91825260209091200155600182018054806116fe576116fe611e6b565b600190038181905f5260205f20015f905590555b600101611670565b505050565b828054828255905f5260205f20908101928215611758579160200282015b8281111561175857825182559160200191906001019061173d565b50611764929150611768565b5090565b5b80821115611764575f8155600101611769565b80356001600160a01b0381168114611792575f80fd5b919050565b5f602082840312156117a7575f80fd5b6117b08261177c565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126117da575f80fd5b813567ffffffffffffffff8111156117f4576117f46117b7565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715611823576118236117b7565b60405281815283820160200185101561183a575f80fd5b816020850160208301375f918101602001919091529392505050565b5f8060408385031215611867575f80fd5b6118708361177c565b9150602083013567ffffffffffffffff81111561188b575f80fd5b611897858286016117cb565b9150509250929050565b5f602082840312156118b1575f80fd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b604081525f6118f860408301856118b8565b905082151560208301529392505050565b5f806040838503121561191a575f80fd5b82359150602083013567ffffffffffffffff81111561188b575f80fd5b5f8151808452602084019350602083015f5b82811015611967578151865260209586019590910190600101611949565b5093949350505050565b602081525f82516080602084015261198c60a08401826118b8565b90506020840151601f198483030160408501526119a98282611937565b9150506040840151601f198483030160608501526119c78282611937565b9150506060840151151560808401528091505092915050565b5f805f606084860312156119f2575f80fd5b83359250611a026020850161177c565b929592945050506040919091013590565b602080825282518282018190525f918401906040840190835b81811015611a4a578351835260209384019390920191600101611a2c565b509095945050505050565b5f8060408385031215611a66575f80fd5b50508035926020909101359150565b6001600160a01b0389168152610100602082018190525f90611a999083018a6118b8565b88604084015287151560608401528281036080840152611ab981886118b8565b905082810360a0840152611acd81876118b8565b905082810360c0840152611ae181866118b8565b91505082151560e08301529998505050505050505050565b60208152611b136020820183516001600160a01b03169052565b5f60208301516101006040840152611b2f6101208401826118b8565b9050604084015160608401526060840151611b4e608085018215159052565b506080840151838203601f190160a0850152611b6a82826118b8565b91505060a0840151601f198483030160c0850152611b8882826118b8565b91505060c0840151601f198483030160e0850152611ba682826118b8565b91505060e0840151611bbd61010085018215159052565b509392505050565b5f805f805f60a08688031215611bd9575f80fd5b85359450602086013567ffffffffffffffff811115611bf6575f80fd5b611c02888289016117cb565b945050604086013567ffffffffffffffff811115611c1e575f80fd5b611c2a888289016117cb565b935050606086013567ffffffffffffffff811115611c46575f80fd5b611c52888289016117cb565b92505060808601358015158114611c67575f80fd5b809150509295509295909350565b600181811c90821680611c8957607f821691505b602082108103611ca757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561171a57805f5260205f20601f840160051c81016020851015611cd25750805b601f840160051c820191505b81811015611cf1575f8155600101611cde565b5050505050565b815167ffffffffffffffff811115611d1257611d126117b7565b611d2681611d208454611c75565b84611cad565b6020601f821160018114611d58575f8315611d415750848201515b5f19600385901b1c1916600184901b178455611cf1565b5f84815260208120601f198516915b82811015611d875787850151825560209485019460019092019101611d67565b5084821015611da457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f6117b060208301846118b8565b634e487b7160e01b5f52601160045260245ffd5b5f60018201611dea57611dea611dc5565b5060010190565b60a081525f611e0360a08301886118b8565b8281036020840152611e1581886118b8565b90508281036040840152611e2981876118b8565b94151560608401525050608001529392505050565b634e487b7160e01b5f52603260045260245ffd5b81810381811115611e6557611e65611dc5565b92915050565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220331448c7ed8fa31183890ede279a1ab26f4d19463abc797ace54448c8f1e785b64736f6c634300081a0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
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
	TransactOpts bind.TransactOpts // Transaction accountAuth options to use throughout this session
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
	TransactOpts bind.TransactOpts   // Transaction accountAuth options to use throughout this session
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

// LandSurveyingArea is a paid mutator transaction binding the contract method 0xc35ce054.
//
// Solidity: function LandSurveyingArea(uint256 _landId, uint256 _area) returns()
func (_Contract *ContractTransactor) LandSurveyingArea(opts *bind.TransactOpts, _landId *big.Int, _area *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "LandSurveyingArea", _landId, _area)
}

// LandSurveyingArea is a paid mutator transaction binding the contract method 0xc35ce054.
//
// Solidity: function LandSurveyingArea(uint256 _landId, uint256 _area) returns()
func (_Contract *ContractSession) LandSurveyingArea(_landId *big.Int, _area *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LandSurveyingArea(&_Contract.TransactOpts, _landId, _area)
}

// LandSurveyingArea is a paid mutator transaction binding the contract method 0xc35ce054.
//
// Solidity: function LandSurveyingArea(uint256 _landId, uint256 _area) returns()
func (_Contract *ContractTransactorSession) LandSurveyingArea(_landId *big.Int, _area *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LandSurveyingArea(&_Contract.TransactOpts, _landId, _area)
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

// RegisterLand is a paid mutator transaction binding the contract method 0x47b24c08.
//
// Solidity: function registerLand(uint256 _landId, string _location) returns()
func (_Contract *ContractTransactor) RegisterLand(opts *bind.TransactOpts, _landId *big.Int, _location string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerLand", _landId, _location)
}

// RegisterLand is a paid mutator transaction binding the contract method 0x47b24c08.
//
// Solidity: function registerLand(uint256 _landId, string _location) returns()
func (_Contract *ContractSession) RegisterLand(_landId *big.Int, _location string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterLand(&_Contract.TransactOpts, _landId, _location)
}

// RegisterLand is a paid mutator transaction binding the contract method 0x47b24c08.
//
// Solidity: function registerLand(uint256 _landId, string _location) returns()
func (_Contract *ContractTransactorSession) RegisterLand(_landId *big.Int, _location string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterLand(&_Contract.TransactOpts, _landId, _location)
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

// VerifyLand is a paid mutator transaction binding the contract method 0xeb4c2774.
//
// Solidity: function verifyLand(uint256 _landId, string _detailsHash, string _reportHash, string _documentsHash, bool _isVerified) returns()
func (_Contract *ContractTransactor) VerifyLand(opts *bind.TransactOpts, _landId *big.Int, _detailsHash string, _reportHash string, _documentsHash string, _isVerified bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifyLand", _landId, _detailsHash, _reportHash, _documentsHash, _isVerified)
}

// VerifyLand is a paid mutator transaction binding the contract method 0xeb4c2774.
//
// Solidity: function verifyLand(uint256 _landId, string _detailsHash, string _reportHash, string _documentsHash, bool _isVerified) returns()
func (_Contract *ContractSession) VerifyLand(_landId *big.Int, _detailsHash string, _reportHash string, _documentsHash string, _isVerified bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyLand(&_Contract.TransactOpts, _landId, _detailsHash, _reportHash, _documentsHash, _isVerified)
}

// VerifyLand is a paid mutator transaction binding the contract method 0xeb4c2774.
//
// Solidity: function verifyLand(uint256 _landId, string _detailsHash, string _reportHash, string _documentsHash, bool _isVerified) returns()
func (_Contract *ContractTransactorSession) VerifyLand(_landId *big.Int, _detailsHash string, _reportHash string, _documentsHash string, _isVerified bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyLand(&_Contract.TransactOpts, _landId, _detailsHash, _reportHash, _documentsHash, _isVerified)
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
