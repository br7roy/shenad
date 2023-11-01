package go_test

import (
	_ "github.com/mattn/go-sqlite3"
	"shenad/main/opts"
	"testing"
)

func TestInit(t *testing.T) {
	opts.LoadConfig("../conf.toml")
	opts.InitDB()
	var user opts.User
	token := "bf17f001-6656-49c0-9331-3a71a98afc90"
	entry, err := user.QueryByToken(token)
	if err != nil {
		t.Error("err:", err.Error())
	} else {
		t.Log(entry)
	}

}
