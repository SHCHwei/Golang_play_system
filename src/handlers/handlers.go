package handlers

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "time"
)


type Team struct {
	Title string `json:"title"`
	Team string `json:"team"`
	How_many string `json:"how_many"`
}

type Teams struct {
    Teams []Team
}


func ShowList(w http.ResponseWriter, r *http.Request) {

    jsonFile, err := os.Open("team.json")
    play, err := ioutil.ReadAll(jsonFile)

    if err != nil {
        log.Fatal(err) // if err exists log fetal and exit
    }

    defer jsonFile.Close()

    w.Write(play)

}

func AddTeam(w http.ResponseWriter, r *http.Request) {

    //建立新隊伍
    newTeam :=  []byte([{ 'title':r.FormValue("title"), 'team': r.FormValue("team"), 'how_many': r.FormValue("how_many"), 'time': time.Now().Unix()}])
    fmt.Println(newTeam)


    //讀取全部隊伍
    jsonFile, err := os.Open("team2.json")
    play, err := ioutil.ReadAll(jsonFile)

    if err != nil {
        log.Fatal(err) // if err exists log fetal and exit
    }

    defer jsonFile.Close()

    var playTeam []Teams
	json.Unmarshal(play, &playTeam)

    //新舊合併
    //append(play, newTeam)
    w.Write(play)

}

/*

func DeleteTeam(w http.ResponseWriter, r *http.Request) {

    if err := r.ParseForm(); err != nil {
    	fmt.Println("ParseForm() err:")
    }

}


func UpdateTeam(w http.ResponseWriter, r *http.Request) {


    if err := r.ParseForm(); err != nil {
    	fmt.Fprintf(w, "ParseForm() err: %v", err)
    }


}
*/