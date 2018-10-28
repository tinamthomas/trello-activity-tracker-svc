package main

import (
    "log"
    "net/http"
    "./controllers"
)

func main() {
    http.HandleFunc("/", controllers.WelcomeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}