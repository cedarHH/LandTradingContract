package land

import (
	"context"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserInfoLogic {
	return &QueryUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserInfoLogic) QueryUserInfo(
	req *types.QueryUserInfoReq) (resp *types.QueryUserInfoResp, err error) {

	//callOpts := &bind.CallOpts{
	//	From: common.HexToAddress("0x0352782844B97688590dfe893DD3062654De7EE6"),
	//}
	//info, err := l.svcCtx.Conn.QueryUserInfo(callOpts)
	//if err != nil {
	//	log.Fatalf("Failed to call queryUserInfo: %v", err)
	//}
	//
	//return &types.QueryUserInfoResp{
	//	Code: 0,
	//	Data: types.UserDetail{
	//		Username:          info.Username,
	//		LandIdList:        info.LandIdList,
	//		TransactionIdList: []int64{},
	//	},
	//	Msg: "",
	//}, nil
	return
}
