package service

import (
	"context"
	"fmt"
	"github.com/mhthrh/common_pkg/pkg/logger"
	"github.com/mhthrh/common_pkg/pkg/model/user"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"go.uber.org/zap"
	"restfullApi/pkg/dao"
	"restfullApi/pkg/xValidation"
)

type Service struct {
	logger logger.ILogger
	dao    user.IUser
	valid  xValidation.IValidation
}

func New() *Service {
	return &Service{
		logger: logger.NewLogger(),
		dao:    dao.New(),
		valid:  xValidation.New(),
	}
}
func (s Service) Create(ctx context.Context, user *user.User) *xErrors.Error {
	s.logger.Info(ctx, "start method Create", zap.Any("user object", user))

	s.logger.Info(ctx, "start parameter validation")
	if err := s.valid.Create(user); err != nil {
		s.logger.Error(ctx, fmt.Sprintf("create user validation error, %v", err), zap.Any("error", err))
		return err
	}
	s.logger.Info(ctx, "create user validation was successfully")

	if err := s.dao.Create(ctx, user); err != nil {
		s.logger.Error(ctx, "error in create user", zap.Any("create error", err))
		return err
	}
	s.logger.Info(ctx, "user created successfully")
	return xErrors.Success()

}

func (s Service) GetByUserName(ctx context.Context, userName string) (user.User, *xErrors.Error) {
	s.logger.Info(ctx, "start method GetByUserName", zap.String("user name", userName))

	if err := s.valid.GetByUserName(userName); err != nil {
		s.logger.Error(ctx, fmt.Sprintf("GetByUserName validation error, %v", err), zap.Any("error", err))
		return user.User{}, err
	}
	s.logger.Info(ctx, "validation was successful")

	s.logger.Info(ctx, "start dao call")
	u, err := s.dao.GetByUserName(ctx, userName)
	if err != nil {
		s.logger.Error(ctx, fmt.Sprintf("GetByUserName dao call error, %v", err), zap.Any("error", err))
		return user.User{}, err
	}
	s.logger.Info(ctx, "GetByUserName call was successful")
	return u, xErrors.Success()
}

func (s Service) Update(ctx context.Context, user *user.User) *xErrors.Error {
	s.logger.Info(ctx, "start method Update", zap.Any("user name", user))

	if err := s.valid.Update(user); err != nil {
		s.logger.Error(ctx, fmt.Sprintf("Update validation error, %v", err), zap.Any("error", err))
		return err
	}
	s.logger.Info(ctx, "validation was successful")

	s.logger.Info(ctx, "start dao call")

	err := s.dao.Update(ctx, user)
	if err != nil {
		s.logger.Error(ctx, fmt.Sprintf("Update dao call error, %v", err), zap.Any("error", err))
		return err
	}
	s.logger.Info(ctx, "update call was successful")
	return xErrors.Success()
}

func (s Service) Remove(ctx context.Context, userName string) *xErrors.Error {
	s.logger.Info(ctx, "start method Remove", zap.String("user name", userName))

	if err := s.valid.Remove(userName); err != nil {
		s.logger.Error(ctx, fmt.Sprintf("Remove validation error, %v", err), zap.Any("error", err))
		return err
	}
	s.logger.Info(ctx, "validation was successful")

	s.logger.Info(ctx, "start dao call")

	err := s.dao.Remove(ctx, userName)
	if err != nil {
		s.logger.Error(ctx, fmt.Sprintf("Remove dao call error, %v", err), zap.Any("error", err))
		return err
	}
	s.logger.Info(ctx, "Remove call was successful")
	return xErrors.Success()
}
