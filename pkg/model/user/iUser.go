package user

import (
	"context"
	"github.com/google/uuid"
	"restfullApi/pkg/errors"
	"time"
)

type User struct {
	id          uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Status      Status
	UserName    string
	Password    string
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
