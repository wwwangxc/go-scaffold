package jwt

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	conf := &Config{
		Issuer: "xc",
		Secret: "盐",
		TTL:    5 * time.Minute,
	}
	conf.Init()

	user := &User{
		Name:     "李狗蛋",
		Birthday: time.Now(),
	}
	token, err := Gen(user)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("token: " + token)

	user2 := &User{}
	err = Parse(token, user2)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(user2)
}

type User struct {
	Name     string
	Birthday time.Time
}
