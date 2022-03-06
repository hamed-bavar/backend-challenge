package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func StartApp() {
	// Start the app
	router := mux.NewRouter()
	router.
		HandleFunc("/customers", getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")

}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hey")
}
