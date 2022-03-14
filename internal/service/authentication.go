package service

import (
	"context"
	"fmt"
	"go-scaffold/internal/constant"
	"go-scaffold/internal/model"
	"go-scaffold/pkg/cache/xredis"
	"go-scaffold/pkg/util"
	"time"

	"github.com/pkg/errors"
)

// global error
var (
	ErrInvalidUsernameOrPassword = errors.New("用户名或密码错误")
	ErrAuthFailed                = errors.New("用户授权失败")
)

// AuthenticationService ..
type AuthenticationService struct{}

// Login ..
func (t *AuthenticationService) Login(ctx context.Context, userName, password string) (string, error) {
	if len(userName) == 0 || len(password) == 0 {
		return "", ErrInvalidUsernameOrPassword
	}

	user, err := model.FindUserByUserName(ctx, userName)
	if err != nil {
		return "", fmt.Errorf("find user by user name fail. err:%w", err)
	}

	if user.Password != util.WithSecret(password, user.Salt) {
		return "", ErrInvalidUsernameOrPassword
	}

	sessionID := util.GenUUID()
	sessionKey := constant.RedisKeySession + sessionID
	fields := map[string]interface{}{
		"userID":   user.ID,
		"username": user.UserName,
	}
	if !xredis.Store(constant.RedisStoreNameDB0).HMSet(sessionKey, fields, 15*time.Minute) {
		return "", ErrAuthFailed
	}

	return sessionID, nil
}

// Logout ..
func (t *AuthenticationService) Logout(sessionID string) {
	xredis.Store(constant.RedisStoreNameDB0).Del(constant.RedisKeySession + sessionID)
}
