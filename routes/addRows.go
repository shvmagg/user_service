package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"1program.com/tournament_app/db"
)

type Data struct {
	Name   string
	Number string
}

func addRows(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	var info Data
	err := decoder.Decode(&info)
	if err != nil {
		fmt.Fprintln(w, "Error Occured")
		panic(err)
	}

	sql := "insert into users (username,contactno) values ('" + info.Name + "', " + info.Number + ") RETURNING userid;"

	rows, err := db.DB().Query(sql)

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
