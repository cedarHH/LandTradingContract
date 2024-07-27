// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import './LandRegistry.sol';
// 土地交易智能合约
contract LandTransaction is LandRegistry{
    
    
    struct Transaction {
         uint256 landId;   
         address from;
         address to;
         uint256 timestamp;
     }
     
     mapping(uint256 => Transaction) public transactions;
     // 记录一个土地的交易记录id
     mapping(uint256 => uint256[]) public landsTransactionIdList;
     uint256 public transactionCounter; 
     modifier isYourLands(uint256 _landId) {
         require(_checkLandsOwner(_landId), "not your land");
         _;
     }
     
     event TransactionRecorded(uint256 indexed transactionId, uint256 indexed landId, address indexed from, address to, uint256 timestamp);
     
    
     function transferOwnership(uint256 _landId, address _to, uint256 _transactionId) public isYourLands(_landId) {
         require(msg.sender != _to, "Cannot be transferred to oneself");
         // 土地需被公证
         require(lands[_landId].isVerified, "this land is not Verified");
         // 接收者必须存在
         require(userMapping[_to].isVaild, "toUser is not exists");
         userMapping[_to].landIdList.push(_landId);
         _removeLandsId(_landId);
         
         userMapping[_to].transactionIdList.push(_transactionId);
         userMapping[msg.sender].transactionIdList.push(_transactionId);
         
         transactions[_transactionId] = Transaction(_landId, msg.sender, _to, block.timestamp);
         emit TransactionRecorded(_transactionId, _landId, msg.sender, _to, block.timestamp);
         transactionCounter++;
         landsTransactionIdList[_landId].push(_transactionId);
         
     }
     function getTransactionDetails(uint256 _transactionId) public view returns(Transaction memory){
         return transactions[_transactionId];
     } 
    function getLandTransactionHistory(uint256 _landId) public view returns (uint256[] memory) {
        return landsTransactionIdList[_landId];
    }
    
    
    function _checkLandsOwner(uint256 _landId) internal view returns(bool) {
        User storage user = userMapping[msg.sender];
        for (uint256 i = 0; i < user.landIdList.length; i++) {
            if (user.landIdList[i] == _landId) {
                return true;
            }
        }
        return false;
    }
    
      function _removeLandsId(uint256 _landId) internal {
        User storage user = userMapping[msg.sender];
        for (uint256 i = 0; i < user.landIdList.length; i++) {
            if (user.landIdList[i] == _landId) {
                user.landIdList[i] = user.landIdList[user.landIdList.length - 1];
                user.landIdList.pop();
            }
        }
    }
  
    
}