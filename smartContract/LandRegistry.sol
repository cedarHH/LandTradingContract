// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import './Ownable.sol';

contract LandRegistry is Ownable{

    constructor(address _oracle) Ownable(_oracle) {
    }

    mapping(address => User) public userMapping;
    mapping(string => Land) public lands;
    
    mapping(address => bool) public surveyors;
    mapping(address => bool) public notaries;
    

    uint256 landCount = 0;
    struct User {
        string username;
        string[] landIdList;
        uint256[] transactionIdList;
        bool isVaild;
    }
    

    struct Land {
         address owner;
         string location;
         uint256 area;
         bool isVerified;
         string detailsHash;
         string reportHash;
         string documentsHash;
         bool isVaild;
    } 



    event LandRegistered(string indexed landId, address indexed owner, string location);

    event LandVerified(string indexed landId, string  detailsHash, string reportHash, string  documentsHash,
                       bool isVerified, address indexed notariesAddress, uint256 timestamp);

    event LandSurveying(string indexed landId, uint256 area, address indexed surveyorsAddress, uint256 timestam);


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
    

    function registerUser(address _userAddress, string memory _userName) public onlyOwner {
        require(!userMapping[_userAddress].isVaild, "user is exists");
        userMapping[_userAddress] = User(_userName, new string[](0),new uint256[](0), true);
    }


    function registerLand(string memory _landId, string memory _location, string memory _detailsHash) public onlyRegistered {
        require(!lands[_landId].isVaild, "land is exists");
        lands[_landId] = Land(msg.sender, _location, 0, false, "", "", "", true);
        lands[_landId].detailsHash = _detailsHash;
        userMapping[msg.sender].landIdList.push(_landId);
        emit LandRegistered(_landId, msg.sender, _location);
        landCount++;
    }

    function LandSurveyingArea(string memory _landId, uint256 _area, string memory _reportHash) public onlySurveyors {
        require(lands[_landId].isVaild, "land is not exists");
        require(msg.sender == oracle, 'Only oracle can call');
        lands[_landId].area = _area;
        lands[_landId].reportHash = _reportHash;
        emit LandSurveying(_landId, _area, msg.sender, block.timestamp);
    }


    function verifyLand(string memory _landId, string memory _documentsHash, bool _isVerified)
    public onlyNotaries {
        require(lands[_landId].isVaild, "land is not exists");
        require(msg.sender == oracle, 'Only oracle can call');
        lands[_landId].documentsHash = _documentsHash;
        lands[_landId].isVerified = _isVerified;
        emit LandVerified(_landId, lands[_landId].detailsHash,lands[_landId].reportHash,_documentsHash,_isVerified, msg.sender, block.timestamp);
    }

    function queryLand(string memory landId) public view returns (Land memory) {
        require(lands[landId].isVerified, "land is not verified");
        return lands[landId];
    }
    
    
    function queryUserInfo() public view returns (User memory) {
        return userMapping[msg.sender];
    }

}