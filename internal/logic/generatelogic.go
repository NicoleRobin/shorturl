package logic

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/nicolerobin/shorturl/internal/constant"
	"github.com/nicolerobin/shorturl/internal/svc"
	"github.com/nicolerobin/shorturl/internal/types"
	"github.com/nicolerobin/shorturl/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"path"
)

type GenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateLogic {
	return &GenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func shortUrl(id int64) string {
	var res bytes.Buffer
	for id > 0 {
		remainder := id % 62
		if remainder >= 0 && remainder < 10 {
			res.WriteByte(byte('0' + remainder))
		} else if remainder >= 10 && remainder < 36 {
			res.WriteByte(byte(int64('A') + remainder))
		} else if remainder >= 36 {
			res.WriteByte(byte(int64('a') + remainder))
		}
		id = id / 62
	}
	return res.String()
}

func (l *GenerateLogic) Generate(req *types.GenerateReq) (resp *types.GenerateRes, err error) {
	logx.Infof("req:%+v", req)

	// 判断是否已存在
	shortUrlModel, err := l.svcCtx.TShorturlModel.FindOneByUrl(l.ctx, req.Url)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorf("TShorturlModel.FindOneByUrl() failed, url:%s, err:%s", req.Url, err)
		return nil, err
	}
	if err == nil {
		return &types.GenerateRes{
			ShortUrl: shortUrlModel.ShortUrl,
		}, nil
	}

	// 获取自增ID
	autoIncrKey := fmt.Sprintf("%s:%s", l.svcCtx.Config.Name, constant.AutoIncKey)
	autoIncrId, err := l.svcCtx.Redis.Incr(autoIncrKey)
	if err != nil {
		logx.Errorf("Redis.Incr() failed, key:%s, err:%s", autoIncrKey, err)
		return nil, err
	}

	// 生成短网址
	shortUrl := shortUrl(autoIncrId)
	shortUrl = path.Join(l.svcCtx.Config.BaseUrl, shortUrl)
	logx.Infof("shortUrl:%s", shortUrl)
	shortUrlModel = &model.TShorturl{
		Url:      req.Url,
		ShortUrl: shortUrl,
	}
	result, err := l.svcCtx.TShorturlModel.Insert(l.ctx, shortUrlModel)
	if err != nil {
		logx.Errorf("TShorturlModel.Insert() failed, err:%s", err)
		return nil, err
	}
	logx.Infof("result:%+v", result)

	return &types.GenerateRes{
		ShortUrl: shortUrl,
	}, nil
}
