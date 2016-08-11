package main

import "github.com/kataras/iris"

type response struct {
    Message string `json:"message"`
    Code string `json:"code"`
}

func main() {

    iris.Get("/", welcome)

    iris.Listen(":8080")
}

func welcome(ctx *iris.Context) {
    message := response{Message: "Welcome", Code: "200"}
    ctx.JSON(iris.StatusOK, message)
}
