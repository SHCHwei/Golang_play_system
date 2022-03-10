package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)



/*
type Team struct {
	Title string `json:"title"`
	Team string `json:"team"`
	How_many int `json:"how_many"`
}

type Teams struct {
    Teams []Team
}

var TeamList map[string] interface{}
*/


type Profile struct {
	Name		string
	Hobbies []string
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", "marray")
}

func main(){
    http.HandleFunc("/", index)
    http.HandleFunc("/json", handler)

    log.Fatal(http.ListenAndServe(":8099", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
/*
    jsonFile, err := os.Open("team2.json")

    if err != nil {
        log.Fatal(err) // if err exists log fetal and exit
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    //將匯入的文字直接解析到 interface，interface{} 可以用來儲存任意資料型別的物件
	json.Unmarshal([]byte(byteValue), &TeamList)
*/

	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)


}
