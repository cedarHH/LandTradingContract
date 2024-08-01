// Code generated by goctl. DO NOT EDIT.
package types

type AddNotaryReq struct {
	SenderKey     string `json:"senderKey"`
	NotaryAddress string `json:"notaryAddress"`
}

type AddNotaryResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type AddSurveyorReq struct {
	SenderKey       string `json:"senderKey"`
	SurveyorAddress string `json:"surveyorAddress"`
}

type AddSurveyorResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetPresignedUrlReq struct {
	FileName string `json:"fileName"`
}

type GetPresignedUrlResp struct {
	Code int    `json:"code"`
	Url  string `json:"url"`
	Msg  string `json:"msg"`
}

type GetTransactionIdResp struct {
	Code int    `json:"code"`
	Data int64  `json:"Data"`
	Msg  string `json:"msg"`
}

type GetUuidResp struct {
	Code int    `json:"code"`
	Uuid string `json:"uuid"`
	Msg  string `json:"msg"`
}

type LandDetail struct {
	Owner      string `json:"owner"`
	Location   string `json:"location"`
	Area       int64  `json:"area"`
	IsVerified bool   `json:"isVerified"`
	Details    string `json:"details"`
	Report     string `json:"report"`
	Document   string `json:"document"`
}

type RegisterLandReq struct {
	Senderkey string `json:"senderkey"`
	LandId    string `json:"landId"`
	Location  string `json:"location"`
	Details   string `json:"details"`
}

type RegisterLandResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RegisterUserReq struct {
	Senderkey   string `json:"senderkey"`
	UserAddress string `json:"userAddress"`
	UserName    string `json:"userName"`
}

type RegisterUserResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SurveyLandReq struct {
	Senderkey string `json:"senderkey"`
	LandId    string `json:"landId"`
	Area      int64  `json:"area"`
	Report    string `json:"report"`
}

type SurveyLandResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type TransactionDetail struct {
	LandId    string `json:"landId"`
	From      string `json:"from"`
	To        string `json:"to"`
	Timestamp uint64 `json:"timestamp"`
}

type TransferContractOwnershipReq struct {
	Senderkey string `json:"senderkey"`
	NewOwner  string `json:"newOwner"`
}

type TransferContractOwnershipResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type TransferVerifyReq struct {
	TransactionId uint64 `json:"transactionId"`
	LandId        string `json:"landId"`
	Address_from  string `json:"addressFrom"`
	Address_to    string `json:"addressTo"`
}

type TransferVerifyResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserDetail struct {
	Username          string   `json:"username"`
	LandIdList        []string `json:"landIdList"`
	TransactionIdList []int64  `json:"transactionIdList"`
}

type VerifyData struct {
	IsTampered bool `json:"isTampered"`
}

type VerifyLandReq struct {
	Senderkey string `json:"senderkey"`
	LandId    string `json:"landId"`
	IsVerify  bool   `json:"isVerify"`
	Document  string `json:"document"`
}

type VerifyLandResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type VerifyOwnershipData struct {
	IsOwner bool `json:"isOwner"`
}

type GenerateProofOfOwnershipReq struct {
	Owner  string `json:"owner"`
	LandId string `json:"landId"`
}

type GenerateProofOfOwnershipResp struct {
	Code  int    `json:"code"`
	Proof string `json:"proof"`
	Msg   string `json:"msg"`
}

type GetLandTransactionHistoryReq struct {
	LandId string `json:"landId"`
}

type GetLandTransactionHistoryResp struct {
	Code int     `json:"code"`
	Data []int64 `json:"data"`
	Msg  string  `json:"msg"`
}

type GetTransactionDetailsReq struct {
	TransactionId uint64 `json:"TransactionId"`
}

type GetTransactionDetailsResp struct {
	Code int               `json:"code"`
	Data TransactionDetail `json:"data"`
	Msg  string            `json:"msg"`
}

type QueryLandReq struct {
	LandId string `json:"landId"`
}

type QueryLandResp struct {
	Code int        `json:"code"`
	Data LandDetail `json:"data"`
	Msg  string     `json:"msg"`
}

type QueryUserInfoReq struct {
	Address string `json:"Address"`
}

type QueryUserInfoResp struct {
	Code int        `json:"code"`
	Data UserDetail `json:"data"`
	Msg  string     `json:"msg"`
}

type TransferOwnershipReq struct {
	Senderkey  string `json:"senderkey"`
	LandId     string `json:"landId"`
	Address_to string `json:"address_to"`
}

type TransferOwnershipResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type VerifyFileHashReq struct {
	LandId string `json:"landId"`
}

type VerifyFileHashResp struct {
	Code int        `json:"code"`
	Data VerifyData `json:"Data"`
	Msg  string     `json:"msg"`
}

type VerifyProofOfOwnershipReq struct {
	LandId string `json:"landId"`
	Proof  string `json:"proof"`
}

type VerifyProofOfOwnershipResp struct {
	Code int                 `json:"code"`
	Data VerifyOwnershipData `json:"Data"`
	Msg  string              `json:"msg"`
}
