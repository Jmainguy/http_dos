package main

import (
    "net/http"
)

func main() {

    http.HandleFunc("/", loadPage)

    http.ListenAndServe(":8000", nil)

}
