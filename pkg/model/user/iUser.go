package user

import (
	"context"
	"github.com/google/uuid"
	"restfullApi/pkg/errors"
	"time"
)

type User struct {
	id          uuid.UUID
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	status      Status
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	createdAt   time.Time
	updatedAt   time.Time
	lastLogin   time.Time
}

type IUser interface {
	Create(ctx context.Context, user *User) *errors.Error
	GetByUserName(ctx context.Context, userName string) (User, *errors.Error)
	Update(ctx context.Context, user *User) *errors.Error
	Remove(ctx context.Context, userName string) *errors.Error
}

func UserInit(u *User) *User {
	u.id = uuid.New()
	u.createdAt = time.Now()
	return u
}
