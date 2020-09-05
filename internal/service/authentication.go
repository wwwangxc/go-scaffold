package service

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/model"
	"go-scaffold/pkg/util"
	"go-scaffold/pkg/xredis"
	"time"

	"github.com/pkg/errors"
)

// AuthenticationService ..
type AuthenticationService struct{}

// Login ..
func (t *AuthenticationService) Login(username string, password string) (string, error) {
	if len(username) == 0 || len(password) == 0 {
		return "", errors.New("用户名或密码错误")
	}
	user := &model.Admin{}
	if err := user.GetByUsername(username); err != nil {
		return "", err
	}
	if user.Password != util.WithSecret(password, user.Salt) {
		return "", errors.New("用户名或密码错误")
	}

	sessionID := util.GenUUID()
	sessionKey := constant.RedisKeySession + sessionID
	fields := map[string]interface{}{
		"userID":   user.ID,
		"username": user.Username,
	}
	if !xredis.Cli.HMSet(sessionKey, fields, 15*time.Minute) {
		return "", errors.New("用户授权失败")
	}

	return sessionID, nil
}

// Logout ..
func (t *AuthenticationService) Logout(sessionID string) {
	xredis.Cli.Del(constant.RedisKeySession + sessionID)
}
