package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"goose_project/models"
	"log"
	"os"
)

var (
	db  *sql.DB
	err error
)

func ConnectDB() (*sql.DB, error) {
	dbSet := models.Setting{
		DriverDB: os.Getenv("DRIVER_DB"),
		User:     os.Getenv("USR"),
		Pass:     os.Getenv("PASS"),
		Port:     os.Getenv("PORT"),
		Name:     os.Getenv("NAME"),
		Reload:   false,
	}
	stmt := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", dbSet.User, dbSet.Pass, dbSet.Port, dbSet.Name)
	db, err = sql.Open(dbSet.DriverDB, stmt)
	err := db.Ping()
	if err != nil {
		fmt.Println("DB is close")
		return nil, err
	}
	if dbSet.Reload{
		log.Println("Start reloading database")
		err := goose.DownTo(db, ".", 0)
		if err !=nil{
			return nil, err
		}
	}
	log.Printf("Start migration database \n")
	err = goose.Up(db, ".")
	if err != nil {
		log.Println(err, "migration error")
	}
	return db, nil
}
