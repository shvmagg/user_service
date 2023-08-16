package routes

import (
	"fmt"
	"net/http"

	"1program.com/tournament_app/db"
	_ "github.com/lib/pq"
)

func getUsersCount(w http.ResponseWriter, req *http.Request) {
	sql := "select COUNT(*) FROM users WHERE name='mamta' "
	rows, err := db.DB().Query(sql)

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
