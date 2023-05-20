package devloper

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthAddKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthAddKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthAddKeyLogic {
	return &DanmuAuthAddKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthAddKeyLogic) DanmuAuthAddKey(req *types.AddKeyRequest) (resp *types.AddKeyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
