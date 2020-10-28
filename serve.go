package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var targetURL = os.Getenv("URL")

type Example struct {
	Greet string
}

func main() {

	//FileServer := http.FileServer(http.Dir("."))
	//http.Handle("/", FileServer)
	mux := mux.NewRouter()
	mux.HandleFunc("/", hellouniverse).Methods("GET")

	http.ListenAndServe(":8080", mux)
}

func hellouniverse(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(targetURL)
	if err != nil {
		//fmt.Printf("Error!", err)
		errmsg := fmt.Sprintf("Error while http.Get: %s", err.Error())
		log.Println(errmsg)
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Printf("Error!", err)
		errmsg := fmt.Sprintf("Error while read body: %s", err.Error())
		log.Println(errmsg)
		http.Error(w, errmsg, http.StatusServiceUnavailable)
		return

	}

	responseString := string(body)

	Hello := Example{Greet: fmt.Sprintf("Content: %s", responseString)}
	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		errmsg := fmt.Sprintf("Error while parse a file: %s", err.Error())
		log.Println(errmsg)
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(Hello)
	tmp.Execute(w, Hello)

}
