package dao

import (
	"context"
	"restfullApi/pkg/errors"
	"restfullApi/pkg/model/user"
	"sync"
)

type Dao struct {
	users map[string]user.User
	lock  *sync.Mutex
}

func New() user.IUser {
	return Dao{make(map[string]user.User), &sync.Mutex{}}
}

func (d Dao) Create(ctx context.Context, user *user.User) *errors.Error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.users[user.UserName]; ok {
		return errors.NewErrUsrExist(nil, nil)
	}
	d.users[user.UserName] = *user
	return nil
}

func (d Dao) GetByUserName(ctx context.Context, userName string) (user.User, *errors.Error) {
	if u, ok := d.users[userName]; ok {
		return u, nil
	}
	return user.User{}, errors.NewErrUsrNotExist(nil, nil)
}

func (d Dao) Update(ctx context.Context, user *user.User) *errors.Error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.users[user.UserName]; !ok {
		return errors.NewErrUsrNotExist(nil, nil)
	}
	delete(d.users, user.UserName)
	d.users[user.UserName] = *user
	return nil
}

func (d Dao) Remove(ctx context.Context, userName string) *errors.Error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.users[userName]; !ok {
		return errors.NewErrUsrNotExist(nil, nil)
	}
	delete(d.users, userName)
	return nil
}
