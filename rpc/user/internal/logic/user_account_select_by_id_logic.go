package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAccountSelectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAccountSelectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAccountSelectByIdLogic {
	return &UserAccountSelectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAccountSelectByIdLogic) UserAccountSelectById(in *user.GetByIdReq) (*user.AccountResp, error) {
	// todo: add your logic here and delete this line

	return &user.AccountResp{}, nil
}
