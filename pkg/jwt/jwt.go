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

// global error
var (
	ErrNotInit = errors.New("gen jwt: need initialize")
	ErrInvalid = errors.New("parse jwt: the token is invalid")
)

type claims struct {
	jwtgo.StandardClaims

	JSON []byte
}

// Gen ..
func Gen(obj interface{}) (string, error) {
	if config == nil {
		return "", ErrNotInit
	}
	json, err := jsoniter.Marshal(obj)
	if err != nil {
		return "", err
	}
	c := &claims{
		JSON: json,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.TTL) * time.Minute).Unix(),
			Issuer:    viper.GetString(config.Issuer),
		},
	}
	token := jwtgo.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Secret))
}

// Parse ..
func Parse(token string, obj interface{}) error {
	if config == nil {
		return ErrNotInit
	}
	c, err := parse(token)
	if err != nil {
		return err
	}
	return json.Unmarshal(c.JSON, obj)
}

// Expired ..
func Expired(token string) (bool, error) {
	if config == nil {
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
		return []byte(config.Secret), nil
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
