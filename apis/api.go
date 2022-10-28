package apis

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ApiTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the test go")
}

func ApiDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running ")
}

func ApiRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	age := vars["age"]
	fmt.Fprintf(w, "Your Name is : %s and your Age is : %s \n", name, age)
}

func ApiQueryParam(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	fmt.Fprintf(w, "Query param is %s \n", param)
}

func ApiQueryParams(w http.ResponseWriter, r *http.Request) {
	one := r.URL.Query().Get("one")
	two := r.URL.Query().Get("two")
	fmt.Fprintf(w, "First Query param is %s \nand second one is %s  ", one, two)
}

func ApiPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed : "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "POST" {
		fmt.Fprintf(w, "This is a post request")

	}
}
