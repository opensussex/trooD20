package main

import  "github.com/kataras/iris"
import "math/rand"


type response struct {
    Message string `json:"message"`
    Code string `json:"code"`
    Value int `json:"value"`
}


func main() {

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
