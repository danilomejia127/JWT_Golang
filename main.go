package main

import "jwt-golang/database"

func main() {
	// Initialize Database
	database.Connect("root:secure_pass_here@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()
	println("running...")
}
