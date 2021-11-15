package main

import (
	"database/sql"
	"fmt"
	web "github.com/igor-koniukhov/webLogger/v2"
	"github.com/subosito/gotenv"
	"goose_project/driver"
	"log"
	"net/http"
)



func init() {

	err := gotenv.Load()
	if err !=nil {
		web.Log.Fatal(err)
	}
}

func main() {
	db, err := driver.ConnectDB()
	if err == nil {
		fmt.Println("Data Base connect")
	} else {
		log.Fatal("DB is close", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	// make migrations

	http.HandleFunc("/", WriteToPage)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func WriteToPage(w http.ResponseWriter, r *http.Request) {
	msg := "hello page"
	_, err := fmt.Fprintf(w, msg)
	if err != nil {
		return
	}

}
