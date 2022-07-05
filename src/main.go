package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "encoding/json"
    "io/ioutil"
)



type Team struct {
	Title string `json:"title"`
	Team string `json:"team"`
	How_many int `json:"how_many"`
}

type Teams struct {
    Teams []Team
}

var TeamList map[string] interface{}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", "marray")
}

func main(){
    http.HandleFunc("/", index)
    http.HandleFunc("/json", handler)

    log.Fatal(http.ListenAndServe(":8099", nil))
}

func index(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

    jsonFile, err := os.Open("team2.json")

    if err != nil {
        log.Fatal(err) // if err exists log fetal and exit
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    //將匯入的文字直接解析到 interface，interface{} 可以用來儲存任意資料型別的物件
	json.Unmarshal([]byte(byteValue), &TeamList)

    w.Write(byteValue)

}
