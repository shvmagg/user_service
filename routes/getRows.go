package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"1program.com/tournament_app/db"
	"1program.com/tournament_app/proj_utils"
)

type getRowsReqModel struct {
	Id []int
}

func getRows(w http.ResponseWriter, req *http.Request) {

	dec := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var info getRowsReqModel
	err := dec.Decode(&info)
	if err != nil {
		fmt.Fprintln(w, "Error in Get")
		panic(err)
	}

	sql := "select username, contactno from users"

	if info.Id != nil && len(info.Id) > 0 {

		tempIds := proj_utils.ArrayToString(info.Id, ", ")

		sql += " where userid in (" + tempIds + ")"

	}

	sql += ";"

	rows, err := db.DB().Query(sql)
	if err != nil {
		fmt.Fprintln(w, "get traversal error")
		panic(err)
	} else {
		for rows.Next() {
			var name, number string
			if err := rows.Scan(&name, &number); err != nil {
				panic(err)
			}
			fmt.Fprintln(w, name+" "+number)
		}
	}
}
