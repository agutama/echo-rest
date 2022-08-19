package models

import (
	"fmt"
	"net/http"

	"github.com/agutama/echo-rest/db"
)

type Company struct {
	id      string `json: "id"`
	Name    string `json: "name"`
	Address string `json: "address,omitempty`
}

func FectAllCompany() (Response, error) {
	var obj Company
	var arrobj []Company
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT id,name,COALESCE(address,'') FROM alamisharia.company LIMIT 10"

	rows, err := con.Query(sqlStatement)

	// defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&obj.id, &obj.Name, &obj.Address)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil

}

func GetCompanyByID(id string) (Response, error) {
	var obj Company
	var arrobj []Company
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id,name,COALESCE(address,'') FROM alamisharia.company WHERE id= $1"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.id, &obj.Name, &obj.Address)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func GetCompanyInID(id string) (Response, error) {
	var obj Company
	var arrobj []Company
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id,name,COALESCE(address,'') FROM alamisharia.company WHERE id IN ($1)"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.id, &obj.Name, &obj.Address)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	fmt.Println(rows)
	return res, nil

}
