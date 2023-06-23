package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, 9)
	fmt.Fprint(w, "hello2\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func userValidator(w http.ResponseWriter, req *http.Request) {
	userId := ""
	for name, headers := range req.Header {
		for _, h := range headers {
			if name == "userId" {
				userId = h
			}
		}
	}
	if userId == "" {
		fmt.Fprint(w, "top Not validated\n")
		fmt.Fprint(w, req.Body)
	} else {
		sql := "select COUNT(*) From users where userid = " + userId
		rows, err := db.Query(sql)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var (
				count int64
			)
			if err := rows.Scan(&count); err != nil {
				panic(err)
			}
			if count == 1 {
				fmt.Fprint(w, "Validated\n")
			} else {
				fmt.Fprint(w, "bot Not validated\n")
			}
		}
	}
}
func getUsersCount(w http.ResponseWriter, req *http.Request) {
	sql := "select COUNT(*) FROM users "
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
		// saveError := fmt.Sprintf("Error Query, and %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			count int64
		)
		if err := rows.Scan(&count); err != nil {
			panic(err)
		}
		fmt.Fprint(w, count)
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "tournament_app"
)

var db *sql.DB

func main() {
	var err error
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/getUsersCount", getUsersCount)
	http.HandleFunc("/Validator", userValidator)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	http.ListenAndServe("192.168.29.151:1000", nil)
}