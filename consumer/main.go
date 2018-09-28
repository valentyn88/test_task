package main

import (
	custom_db "test_task/consumer/db"
	"net/http"
	"encoding/json"
	"test_task/consumer/user"
	"fmt"
)

type API struct {
	db custom_db.DB
}

func NewAPI(db custom_db.DB) API {
	return API{db:db}
}

func main()  {
	db := custom_db.NewDB()
	api := NewAPI(db)
	http.HandleFunc("/", api.userHandler)
	http.ListenAndServe(":9090", nil)
}

func (api API) userHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var u user.User
	if err := decoder.Decode(&u); err != nil {
		http.Error(w, fmt.Sprintf("Error while parsing request from reader service %s", err), http.StatusInternalServerError)
	}
	api.db.AddRow(u)
}
