package main

import (
    "net/http"
    "io"
    "os"
   // "strconv"
    "fmt"
)

func loadPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    Openfile, _ := os.Open("./test.img")
    defer Openfile.Close() //Close after function return
    /*
    FileStat, _ := Openfile.Stat()                     //Get info from file
    FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string
    */
    w.Header().Set("Content-Disposition", "attachment; filename=neat.img")
    w.Header().Set("Content-Length", "53687091")
    io.Copy(w, Openfile)
}
