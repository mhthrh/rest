package dao

import (
	"context"
	"github.com/mhthrh/common_pkg/pkg/model/user"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"sync"
)

type Dao struct {
	users map[string]user.User
	lock  *sync.Mutex
}

func New() user.IUser {
	return Dao{make(map[string]user.User), &sync.Mutex{}}
}

func (d Dao) Create(ctx context.Context, user *user.User) *xErrors.Error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.users[user.UserName]; ok {
		return xErrors.NewErrUsrExist(nil, nil)
	}
	d.users[user.UserName] = *user
	return nil
}

func (d Dao) GetByUserName(ctx context.Context, userName string) (user.User, *xErrors.Error) {
	if u, ok := d.users[userName]; ok {
		return u, nil
	}
	return user.User{}, xErrors.NewErrUsrNotExist(nil, nil)
}

func (d Dao) Update(ctx context.Context, user *user.User) *xErrors.Error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.users[user.UserName]; !ok {
		return xErrors.NewErrUsrNotExist(nil, nil)
	}
	delete(d.users, user.UserName)
	d.users[user.UserName] = *user
	return nil
}

func (d Dao) Remove(ctx context.Context, userName string) *xErrors.Error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.users[userName]; !ok {
		return xErrors.NewErrUsrNotExist(nil, nil)
	}
	delete(d.users, userName)
	return nil
}
