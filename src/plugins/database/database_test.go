package database

import (
	"fmt"
	"testing"
	"time"
)

func TestTimezone(t *testing.T) {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	nowloc := time.Now().In(loc)
	fmt.Println(loc.String())
	fmt.Println(nowloc)
	fmt.Println(now)
}

func TestNew(t *testing.T) {
	db := New(WithDSN("root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"))
	sqldb, _ := db.DB()
	fmt.Println(sqldb.Ping())
}
