package main

import (
    "fmt"
    "net/http"
    "io"
    "crypto/md5"
    "time"
    "strconv"
    "html/template"
)

func loadPage(w http.ResponseWriter, r *http.Request) {
       fmt.Println("method:", r.Method)
       crutime := time.Now().Unix()
       h := md5.New()
       io.WriteString(h, strconv.FormatInt(crutime, 10))
       token := fmt.Sprintf("%x", h.Sum(nil))
       t, _ := template.ParseFiles("index.gtpl")
       t.Execute(w, token)
}
