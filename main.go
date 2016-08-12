package main

import  (
    "github.com/kataras/iris"
    "math/rand"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type response struct {
    Message string `json:"message"`
    Code string `json:"code"`
    Value int `json:"value"`
}

type Character struct {
  gorm.Model
  Name string
  Age uint
}

func main() {
    db, err := gorm.Open("sqlite3", "trood20.db")
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Character{})
    iris.Get("/", welcome)
    iris.Get("/dice", dice)

    iris.Listen(":8080")
}

func welcome(ctx *iris.Context) {
    message := response{Message: "Welcome", Code: "200"}
    ctx.JSON(iris.StatusOK, message)
}

func dice(ctx *iris.Context) {
    roll := roll()
    message := response{Message: "Dice", Code: "200", Value: roll}
    ctx.JSON(iris.StatusOK, message)
}

func roll() int {
    return rand.Intn(20) + 1
}
