package main

import (
    "github.com/kataras/iris/v12"
    "io/ioutil"
    "os"
    "log"
    "encoding/json"
)

type webConfig struct {
	Driver string `json:"driver"`
	DBName string `json:"dbname"`
	User   string `json:"user"`
}

var WebConfig webConfig



type Team struct {
	Title string `json:"title"`
	Team string `json:"team"`
	How_many int `json:"how_many"`
}

type Teams struct {
    Teams []Team
}

var TeamList map[string] interface{}

func main(){

    // iris.New() 和 iris.Default()
    // iris.New() -> Creates an iris application without any middleware by default
    //app := iris.New()

    app.UseRouter(crs)

    //傳入index方法
    app.Get("/", crs, index)

    app.Get("/json", crs, func(ctx iris.Context) {


    //傳入index方法
    app.Post("/",index)
    app.Get("/",index)

    app.Post("/json", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
    })

    app.Get("/json_secure", crs, func(ctx iris.Context) {
        response := []string{"val1", "val2", "val3"}
        options := iris.JSON{Indent: ""}
        ctx.JSON(response, options)
    })

    app.Listen(":8099")
}

func index(ctx iris.Context){

    ctx.Header("Access-Control-Allow-Origin", "*")
    ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
    ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")

    jsonFile, err := os.Open("team.json")
    jsonFile2, err := os.Open("team2.json")

    if err != nil {
        log.Fatal(err) // if err exists log fetal and exit
    }

    defer jsonFile.Close()
    defer jsonFile2.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    byteValue2, _ := ioutil.ReadAll(jsonFile2)

    //將匯入的文字直接解析到 interface，interface{} 可以用來儲存任意資料型別的物件
	json.Unmarshal([]byte(byteValue), &TeamList)


    var list Teams

    // Golang中字串更改，需要先轉換成[]byte 型別 如範例一
    json.Unmarshal([]byte(byteValue2), &list)

    /*  範例一 : 將hello 切割 string 變成 byte型別的slice
        s := "hello"
        c := []byte(s)  // 將字串 s 轉換為 []byte 型別
        c[0] = 'c'
        s2 := string(c)  // 再轉換回 string 型別
    */


    //從interface 中取出資料
    fmt.Println(list.Teams[0])

    //data := dbcon.DBT()

    //回傳json型態
    ctx.JSON(list)

}
