package main

import (
    "net/http"
)

func main() {

    http.HandleFunc("/", loadPage)

    http.ListenAndServe(":25550", nil)

}
