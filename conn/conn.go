package conn

import (
	"fmt"
	"net/http"
	"strconv"
	// _ "github.com/lib/pq"
)

func Start() {
	fmt.Println("Successfully connected!")

	http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
}
