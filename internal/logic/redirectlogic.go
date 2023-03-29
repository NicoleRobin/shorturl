package logic

import (
	"context"

	"github.com/nicolerobin/shorturl/internal/svc"
	"github.com/nicolerobin/shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectLogic {
	return &RedirectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectLogic) Redirect(req *types.RedirectReq) (resp *types.RedirectRes, err error) {
	logx.Infof("req:%+v", req)
	shortUrlModel, err := l.svcCtx.TShorturlModel.FindOneByShortUrl(l.ctx, req.ShortUrl)
	if err != nil {
		logx.Errorf("TShorturlModel().FindOneByShortUrl() failed, shortUrl:%s, err:%s", shortUrl, err)
		return nil, err
	}
	logx.Infof("shortUrlModel:%+v", shortUrlModel)

	// #TODO: 怎么设置返回的httpStatus值
	return &types.RedirectRes{}, nil
}
