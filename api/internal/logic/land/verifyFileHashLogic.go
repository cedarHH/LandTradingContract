package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyFileHashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyFileHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyFileHashLogic {
	return &VerifyFileHashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyFileHashLogic) VerifyFileHash(req *types.VerifyFileHashReq) (resp *types.VerifyFileHashResp, err error) {
	// todo: add your logic here and delete this line
	// 获取交易授权信息
    auth := l.svcCtx.AccountAuth.GetAccountAuth(req.NotaryAddress)
    
    // 调用智能合约的 verifyLand 方法并传入授权信息和文件哈希信息
    tx, err := l.svcCtx.Conn.VerifyLand(auth, big.NewInt(int64(req.LandID)), req.DetailsHash, req.ReportHash, req.documentHash, req.IsVerified)
    if err != nil {
        log.Fatalf("Failed to call VerifyLand: %v", err)
        return &types.VerifyFileHashResp{
            Code: 1,
            Msg:  "Failed to verify file hash",
        }, err
    }

    log.Printf("Transaction hash: %s", tx.Hash().Hex())
    return &types.VerifyFileHashResp{
        Code: 0,
        Msg:  "File hash verified successfully",
    }, nil
}
