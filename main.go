package main

import (
        "net/http"
        "fmt"
        "github.com/gorilla/mux"
)

func main() {
        router := mux.NewRouter()
        router.HandleFunc("/hello/{name}", handleSayHello).Methods("GET")
        http.ListenAndServe(":8080", router)
}

func handleSayHello(rw http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        msg := fmt.Sprint("Hello ", vars["name"])
        fmt.Fprint(rw, msg)
}
