package main

import (
	"github.com/agutama/echo-rest/db"
	"github.com/agutama/echo-rest/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
