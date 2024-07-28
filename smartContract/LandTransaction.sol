// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import './LandRegistry.sol';
// 土地交易智能合约
contract LandTransaction is LandRegistry{
    
    
    constructor(address _oracle) LandRegistry(_oracle) {
    }
    
    struct Transaction {
         string landId;
         address from;
         address to;
         uint256 timestamp;
         bool isVaild;
     }
     
     mapping(uint256 => Transaction) public transactions;
     // 记录一个土地的交易记录id
     mapping(string => uint256[]) public landsTransactionIdList;
     uint256 public transactionCounter; 
     modifier isYourLands(string memory _landId, address _userAddress) {
         require(_checkLandsOwner(_landId, _userAddress), "not your land");
         _;
     }
     
     event InitiateTransaction(string indexed landId, address indexed from, address to, uint256 timestamp);
     
     event TransactionRecorded(uint256 indexed transactionId, string indexed landId, address indexed from, address to, uint256 timestamp);
     
    
     // 发起土地交易
     function transferOwnership(string memory _landId, address _to) public isYourLands(_landId, msg.sender) {
         require(msg.sender != _to, "Cannot be transferred to oneself");
         // 土地需被公证
         require(lands[_landId].isVerified, "this land is not Verified");
         // 接收者必须存在
         require(userMapping[_to].isVaild, "toUser is not exists");
         // 交易发起事件
         emit InitiateTransaction(_landId, msg.sender, _to, block.timestamp);
     }
     
     // 土地交易验证
     function transferVerify(string memory _landId, address _from, address _to, uint256 _transactionId) onlyOracle public {
         require(!transactions[_transactionId].isVaild, "transactionId is exists");
         require(_checkLandsOwner(_landId, _from), "not fromAddress land");
         require(_from != _to, "Cannot be transferred to oneself");
         // 土地需被公证
         require(lands[_landId].isVerified, "this land is not Verified");
         // 接收者必须存在
         require(userMapping[_to].isVaild, "toUser is not exists");
         
         userMapping[_to].landIdList.push(_landId);
         _removeLandsId(_landId, _from);
         
         userMapping[_to].transactionIdList.push(_transactionId);
         userMapping[_from].transactionIdList.push(_transactionId);
         
         transactions[_transactionId] = Transaction(_landId, _from, _to, block.timestamp, true);
         emit TransactionRecorded(_transactionId, _landId, _from, _to, block.timestamp);
         transactionCounter++;
         landsTransactionIdList[_landId].push(_transactionId);
     }
     
     
     function getTransactionDetails(uint256 _transactionId) public view returns(Transaction memory){
         return transactions[_transactionId];
     } 
    function getLandTransactionHistory(string memory _landId) public view returns (uint256[] memory) {
        return landsTransactionIdList[_landId];
    }
    
    
    function _checkLandsOwner(string memory _landId, address _userAddress) internal view returns(bool) {
        User storage user = userMapping[_userAddress];
        for (uint256 i = 0; i < user.landIdList.length; i++) {
            if (keccak256(bytes(user.landIdList[i])) == keccak256(bytes(_landId))) {
                return true;
            }
        }
        return false;
    }
    
      function _removeLandsId(string memory _landId, address _userAddress) internal {
        User storage user = userMapping[_userAddress];
        for (uint256 i = 0; i < user.landIdList.length; i++) {
            if (keccak256(bytes(user.landIdList[i])) == keccak256(bytes(_landId))) {
                user.landIdList[i] = user.landIdList[user.landIdList.length - 1];
                user.landIdList.pop();
            }
        }
    }
  
    
}