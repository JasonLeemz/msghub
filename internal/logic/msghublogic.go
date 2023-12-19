package logic

import (
	"context"

	"msghub/internal/svc"
	"msghub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsghubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsghubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsghubLogic {
	return &MsghubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsghubLogic) Msghub(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
