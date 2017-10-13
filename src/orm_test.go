package ormgo

import (
	"fmt"
	"testing"
	//"utils/error"

	_ "github.com/go-sql-driver/mysql"
	//"utils/logger"
)

type Gorm struct {
	Name string
}

func GetDB() DB {
	db, err :=Open("mysql", "ormgo:ormgo@tcp(localhost:3306)/ormgo?charset=utf8")
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
		fmt.Println(err)
	}
	fmt.Println("00")
	return db

}

func TestWhere(t *testing.T) {
	GetDB()
	fmt.Println("11111111111")
	t.Log("99")

}
