package main

import (
	"go-vue-sns/db"
	"go-vue-sns/route"
)

func main() {
	db.Init()
	e := route.Init()
	e.Logger.Fatal(e.Start(":8083"))
}
