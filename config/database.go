package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"fmt"

	"novanxyz/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	driver   = utils.Getenv("DB_TYPE", "mysql")
	host     = utils.Getenv("DB_HOST", "host.docker.internal")
	port     = utils.Getenv("DB_PORT", "3306")
	user     = utils.Getenv("DB_USER", "")
	password = utils.Getenv("DB_PASS", "")
	dbName   = utils.Getenv("DB_NAME", "")
)

func DatabaseConnection() *gorm.DB {
	fmt.Sprintln("%s://%s:%s@%s:%s/%s?sslmode=disable", driver, user, password, host, port, dbName)

	conn, err := sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName))
	utils.ErrorPanic(err)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}
