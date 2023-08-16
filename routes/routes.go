package routes

import (
	"net/http"
)

func AttachRoutes() {
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/getUsersCount", getUsersCount)
	http.HandleFunc("/Validator", validate_user)
	http.HandleFunc("/dataentry", addRows)
	http.HandleFunc("/getrows", getRows)

}
