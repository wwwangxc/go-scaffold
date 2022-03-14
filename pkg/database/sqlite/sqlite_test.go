package sqlite

import "testing"

func TestBuild(t *testing.T) {
	conf := &Config{
		DSN: "../../db/sqlite/test.db",
	}
	db := conf.Build()
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Log(err)
		return
	}
	t.Log("Success.")
}

func TestStore(t *testing.T) {
	conf := &Config{
		DSN: "../../db/sqlite/test.db",
	}
	conf.Append("test")
	if err := Store("test").Ping(); err != nil {
		t.Log(err)
		return
	}
	t.Log("Success.")
}
