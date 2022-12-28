package service

import (
	"context"

	"github.com/Authorization/internal/models"
)

type Service struct {
	Db database
}

func NewService(db database) *Service {
	return &Service{}
}

type database interface {
	GetUserInfo(ctx context.Context, id int)
	PutUserInfo(ctx context.Context, userInfo models.UserInfo)
}

func (s *Service) Login(ctx context.Context, logName string, password string) (models.UserInfo, error) {
	return models.UserInfo{}, nil
}

func (s *Service) Register(ctx context.Context, user models.UserInfo) (int, error) {
	return 0, nil
}
