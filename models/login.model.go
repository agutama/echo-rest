package models

import (
	"database/sql"
	"fmt"

	"github.com/agutama/echo-rest/db"
	"github.com/agutama/echo-rest/helpers"
)

type User struct {
	id       int    `json: "id"`
	Username string `json: "username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username= $1"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match.")
		return false, err
	}

	return true, nil

}
