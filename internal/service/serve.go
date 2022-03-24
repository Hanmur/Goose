package service

import (
	"Goose/global"
	"Goose/internal/dao"
	"context"
)

type Service struct {
	ctx  context.Context
	dao  *dao.Dao
	pool *dao.Pool
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.DBEngine)
	svc.pool = dao.NewPool(global.RedisPool)
	return svc
}
