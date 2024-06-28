package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:truong123456@tcp(localhost:3306)/")
	if err != nil {
		return nil, err
	}

	// Chọn cơ sở dữ liệu
	_, err = db.Exec("USE sys")
	if err != nil {
		// Xử lý lỗi
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
