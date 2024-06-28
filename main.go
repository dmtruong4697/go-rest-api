package main

import (
	"fmt"
	_ "go-rest-api/docs"
	"go-rest-api/src/database"
	"go-rest-api/src/routes"
	"log"
	"net/http"
)

func main() {

	// Khởi tạo kết nối cơ sở dữ liệu
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Khởi tạo router và đăng ký các route
	router := routes.Setup(db)

	// Khởi động server và lắng nghe cổng 8080
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
