// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import './Ownable.sol';
// 土地注册智能合约
contract LandRegistry is Ownable{

    // 存储注册用户，使用映射以用户地址为键，注册状态为值
    mapping(address => User) public userMapping;
    // 存储土地信息，使用映射以土地ID为键，土地结构为值
    mapping(uint256 => Land) public lands;
    
    mapping(address => bool) public surveyors;
    mapping(address => bool) public notaries;
    
    // 记录注册土地数量
    uint256 landCount = 0;
    struct User {
        string username; // 用户名
        uint256[] landIdList; // 拥有土地id列表
        uint256[] transactionIdList; // 交易列表
        bool isVaild; // 用户状态
    }
    
        
    // 土地结构体，包含土地ID、所有者、位置、详细信息和验证状态
    struct Land {
         address owner;
         string location; // 位置
         uint256 area; // 面积
         bool isVerified; // 是否审核通过
         string detailsHash; //  数据文件hash
         string reportHash; // 数据文件hash
         string documentsHash; // 数据文件hash
         bool isVaild; // 判断是否存在
    } 

    // 事件声明，用于触发事件
    
    // 注册土地事件
    event LandRegistered(uint256 indexed landId, address indexed owner, string location);
    // 公证员--土地校验事件
    event LandVerified(uint256 indexed landId, string  detailsHash, string reportHash, string  documentsHash, 
                       bool isVerified, address indexed notariesAddress, uint256 timestamp);
    // 测量员 -- 土地面积测量事件
    event LandSurveying(uint256 indexed landId, uint256 area, address indexed surveyorsAddress, uint256 timestam);


    // 只有注册用户可以执行的修饰符
    modifier onlyRegistered() {
        require(userMapping[msg.sender].isVaild, "is not user");
        _;
    }
    
   modifier onlySurveyors() {
        require(surveyors[msg.sender], "Not a surveyor");
        _;
   }
  
  modifier onlyNotaries() {
       require(notaries[msg.sender], "Not a notaries");
        _;
   }

     function addSurveyor(address _surveyor) public onlyOwner {
         surveyors[_surveyor] = true;
     } 
     function addNotary(address _notary) public onlyOwner {
         notaries[_notary] = true;
     } 
    

    // 注册新用户，只有合约所有者可以调用
    function registerUser(address _userAddress, string memory _userName) public onlyOwner {
        require(!userMapping[_userAddress].isVaild, "user is exists");
        userMapping[_userAddress] = User(_userName, new uint256[](0),new uint256[](0), true);
    }


    // 注册土地，只有注册用户可以调用
    function registerLand(uint256 _landId, string memory _location) public onlyRegistered {
        require(!lands[_landId].isVaild, "land is exists");
        lands[_landId] = Land(msg.sender, _location, 0, false, "", "", "", true);
        userMapping[msg.sender].landIdList.push(_landId);
        emit LandRegistered(_landId, msg.sender, _location);
        landCount++;
    }


    // 测量人员测量土地后传入土地面积
    function LandSurveyingArea(uint256 _landId, uint256 _area) public onlySurveyors {
        require(lands[_landId].isVaild, "land is not exists");
        lands[_landId].area = _area;
        emit LandSurveying(_landId, _area, msg.sender, block.timestamp);
    }


    // 土地公证人员校验土地信息
    function verifyLand(uint256 _landId, string memory _detailsHash, string memory _reportHash, string memory _documentsHash, bool _isVerified) 
    public onlyNotaries {
        require(lands[_landId].isVaild, "land is not exists");
        lands[_landId].detailsHash = _detailsHash;
        lands[_landId].reportHash = _reportHash;
        lands[_landId].documentsHash = _documentsHash;
        lands[_landId].isVerified = _isVerified;
        emit LandVerified(_landId, _detailsHash,_reportHash,_documentsHash,_isVerified, msg.sender, block.timestamp);
    }

    // 查询土地信息，任何人都可以调用
    function queryLand(uint256 landId) public view returns (Land memory) {
        require(lands[landId].isVerified, "land is not verified");
        return lands[landId];
    }
    
    
    function queryUserInfo() public view returns (User memory) {
        return userMapping[msg.sender];
    }

}