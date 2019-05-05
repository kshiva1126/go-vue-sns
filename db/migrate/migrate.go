package main

import (
	d "go-vue-sns/db"
)

func main() {
	d.Init()
	db := d.GetConnection()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	if !db.HasTable(d.Message{}) {
		db.AutoMigrate(&d.Message{})
	}

	if !db.HasTable(d.User{}) {
		db.AutoMigrate(&d.User{})
	}
}
