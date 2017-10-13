package main

import (
	"database/sql"
	"utils/error"
	"utils/logger"

	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
func init(){
	fmt.Println("88888888888")
}

func init(){
	fmt.Println("999999999999")
}

func main() {

	logger.SetLevel(logger.DEBUG)
	logger.EnableColor()
	logger.Debug("11111111111111")
	db, err := sql.Open("mysql", "ormgo:ormgo@tcp(localhost:3306)/ormgo?charset=utf8")
	error.CheckErr(err)
	// stmt, err := db.Prepare(`INSERT  ormgo.user(name,phone) values (?,?)`)
	stmt, err := db.Prepare(`select * from user where id = ?`)
	error.CheckErr(err)
	stmt.Exec("1")

	/*	var (
			id int
			name string
		)
		rows, err := db.Query("select id, name from ormgo.user where id = ?", 2)
		if err != nil {
			logger.Error(err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				logger.Error(err)
				return
			}
			fmt.Println(id, name)
		}
		err = rows.Err()
		if err != nil {
			logger.Error(err)
			return
		}*/

		var name2 string
		err = db.QueryRow("select name from user where id = ?", 6).Scan(&name2)
		if err != nil {
			logger.Error(err)
			return
		}
		fmt.Println("name2:",name2)
	var name3 string
	rows, err := db.Query("select name from user")
	rows.Scan(name3)
	fmt.Println("name3:",name3)
	fmt.Println("rows:",rows)
stringColumns,err := rows.Columns()
fmt.Println("stringColumns:",stringColumns)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		logger.Debug(id, name)
	}
	err = rows.Err() // get any error encountered during iteration

}
