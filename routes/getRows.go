package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"1program.com/tournament_app/db"
)

type For_get struct {
	Id []int
	// name   []string
	// number []string
}

func getRows(w http.ResponseWriter, req *http.Request) {

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

		rows, err := db.DB().Query(sql)
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
