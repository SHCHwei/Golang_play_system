package main

import (
    "log"
    "net/http"
    "myapp/handlers"
)



func main(){
/* 暫時關閉
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    w.Header().Set("Content-Type", "application/json")
*/
    http.HandleFunc("/showList", handlers.ShowList)
    http.HandleFunc("/addTeam", handlers.AddTeam)
    /*
    http.HandleFunc("/deleteTeam", handlers.DeleteTeam)
    http.HandleFunc("/updateTeam", handlers.UpdateTeam)
*/

    log.Fatal(http.ListenAndServe(":8099", nil))
}

