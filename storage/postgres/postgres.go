package postgres

import (
	"database/sql"
	"fmt"
	"payment-service/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	config := config.Load()
	connector := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode = disable",
				config.DB_HOST,config.DB_PORT,config.DB_USER,config.DB_NAME,config.DB_PASSWORD)

	db,err := sql.Open("postgres",connector)
	if err != nil{
		return nil,err
	}
	return db,nil
}
