package main

import "github.com/fanqie/dcmigrate-example/dc_migrations"

func main() {
	if dc_migrations.TryStartUpDcMigrate() {
		return
	}
	// Your code here...
	// gin.Run()
}
