package test

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

//hashedPwd, err = bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)

func TestGenratePass(t *testing.T) {
	password := "admin123"
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Logf("generate password err:%s", err.Error())
		return
	}

	strPwd := string(hashedPwd)
	t.Logf("strPwd:%s", strPwd)

}

func TestTime(t *testing.T) {
	timestamp := int64(1602469602000)
	//at, _ := time.Parse(time.RFC3339, timestamp)
	ts1 := time.Unix(timestamp/1e3, 0)
	ts2 := time.Unix(0, timestamp*int64(time.Millisecond))
	t.Logf("ts1: %v", ts1)
	t.Logf("ts2: %v", ts2)

}
