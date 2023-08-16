package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "io"
	// "log"
	"net/http"
	"strconv"

	// "strings"

	_ "github.com/lib/pq"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func getUsersCount(w http.ResponseWriter, req *http.Request) {
	sql := "select COUNT(*) FROM users WHERE name='mamta' "
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
		fmt.Fprint(w, count)
	}
}

func validate_user(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	fmt.Println(q)

	name, ok := q["name"]

	if !(ok) {
		fmt.Fprintf(w, "no name in query!!")
	} else {
		sql := ("select count(*) from users where username = '" + name[0] + "';")
		rows, err := db.Query(sql)

		if err != nil {
			fmt.Fprintf(w, "Error")
		} else {
			defer rows.Close()
			for rows.Next() {
				var count int64
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
}

type Data struct {
	Name   string
	Number int
}

func create(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	var info Data
	err := decoder.Decode(&info)
	if err != nil {
		fmt.Fprintln(w, "Error Occured")
		panic(err)
	}

	sql := "insert into users (username,contactno) values ('" + info.Name + "', " + strconv.Itoa(info.Number) + ") RETURNING userid;"

	rows, err := db.Query(sql)

	if err != nil {
		fmt.Fprintln(w, "Error")
		panic(err)
	} else {
		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				fmt.Fprintln(w, "Unable to get id")
				panic(err)
			}
			fmt.Fprint(w, id)
		}
	}

}

type For_get struct {
	Id []int
	// name   []string
	// number []string
}

func getrows(w http.ResponseWriter, req *http.Request) {

	dec := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var info For_get
	err := dec.Decode(&info)
	if err != nil {
		fmt.Fprintln(w, "Error in Get")
		panic(err)
	}

	for i := 0; i < len(info.Id); i++ {
		sql := "select * from users where userid=" + strconv.Itoa(info.Id[i]) + ";"

		rows, err := db.Query(sql)
		if err != nil {
			fmt.Fprintln(w, "get traversal error")
			panic(err)
		} else {
			for rows.Next() {
				var name, number string
				var id int
				if err := rows.Scan(&name, &number, &id); err != nil {
					panic(err)
				}
				fmt.Fprintln(w, name+" "+number)
			}
		}
	}
	// sql:="select * from users where userid=1;"
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
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/getUsersCount", getUsersCount)
	http.HandleFunc("/Validator", validate_user)
	http.HandleFunc("/dataentry", create)
	http.HandleFunc("/getrows", getrows)
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

	http.ListenAndServe("192.168.56.1:1000", nil)
}
