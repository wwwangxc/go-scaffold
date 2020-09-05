package jwt

import (
	"encoding/json"
	"errors"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	ErrNotInit = errors.New("gen jwt: 未初始化配置信息")
	ErrInvalid = errors.New("parse jwt: the token is invalid")
)

type claims struct {
	Json []byte
	jwtgo.StandardClaims
}

// Gen ..
func Gen(obj interface{}) (string, error) {
	if _config == nil {
		return "", ErrNotInit
	}
	json, err := jsoniter.Marshal(obj)
	if err != nil {
		return "", err
	}
	c := &claims{
		Json: json,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(_config.TTL) * time.Minute).Unix(),
			Issuer:    viper.GetString(_config.Issuer),
		},
	}
	token := jwtgo.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(_config.Secret))
}

// Parse ..
func Parse(token string, obj interface{}) error {
	if _config == nil {
		return ErrNotInit
	}
	c, err := parse(token)
	if err != nil {
		return err
	}
	return json.Unmarshal(c.Json, obj)
}

// Expired ..
func Expired(token string) (bool, error) {
	if _config == nil {
		return false, ErrNotInit
	}
	c, err := parse(token)
	if err != nil {
		return false, err
	}
	return time.Now().Unix() > c.ExpiresAt, nil
}

func parse(tokenStr string) (*claims, error) {
	token, err := jwtgo.ParseWithClaims(tokenStr, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(_config.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	c, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return nil, ErrInvalid
	}
	return c, nil
}