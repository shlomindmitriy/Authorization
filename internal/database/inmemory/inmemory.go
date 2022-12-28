package inmemory

import (
	"context"
	"github.com/Authorization/internal/models"
)

type InMemory struct {
	db map[int]models.UserInfo
}

func NewInMemory(db map[int]models.UserInfo) *InMemory {
	return &InMemory{}
}

func (i *InMemory) GetUserInfo(ctx context.Context, id int) {

}

func (i *InMemory) PutUserInfo(ctx context.Context, userInfo models.UserInfo) {

}
