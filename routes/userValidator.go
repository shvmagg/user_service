package routes

import (
	"fmt"
	"net/http"

	"1program.com/tournament_app/db"
)

func validate_user(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	fmt.Println(q)

	name, ok := q["name"]

	if !(ok) {
		fmt.Fprintf(w, "no name in query!!")
	} else {
		sql := ("select count(*) from users where username = '" + name[0] + "';")
		rows, err := db.DB().Query(sql)

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
