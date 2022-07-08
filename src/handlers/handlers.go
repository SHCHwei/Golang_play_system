package handlers

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
	Time string `json:time`
}

type Teams struct {
    Teams []Team
}

var TeamList map[string] interface{}


func ShowList(w http.ResponseWriter, r *http.Request) {

    jsonFile, err := os.Open("team.json")
    play, err := ioutil.ReadAll(jsonFile)

    if err != nil {
        log.Fatal(err) // if err exists log fetal and exit
    }

    defer jsonFile.Close()

    //將匯入的文字直接解析到 interface，interface{} 可以用來儲存任意資料型別的物件
	json.Unmarshal([]byte(play), &TeamList)

    w.Write(play)

}

func AddTeam(w http.ResponseWriter, r *http.Request) {


    if err := r.ParseForm(); err != nil {
    	fmt.Fprintf(w, "ParseForm() err: %v", err)
    }

    for _,value := range r.ParseForm{
        fmt.Fprintf(value)
    }


}



func DeleteTeam(w http.ResponseWriter, r *http.Request) {


    if err := r.ParseForm(); err != nil {
    	fmt.Fprintf(w, "ParseForm() err: %v", err)
    }

}


func UpdateTeam(w http.ResponseWriter, r *http.Request) {


    if err := r.ParseForm(); err != nil {
    	fmt.Fprintf(w, "ParseForm() err: %v", err)
    }


}