package main

import (
	"devweek/repositories"
	"devweek/types"
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", HandleGetAll)
	app.Post("/", HandleCreate)

	app.Listen(":3000")
}

func HandleGetAll(ctx *fiber.Ctx) error {
	allTodos := repositories.GetAll()

	return ctx.JSON(allTodos)
}

func HandleCreate(ctx *fiber.Ctx) error {
	todo := &types.TodoType{}

	ctx.BodyParser(todo)

	err := repositories.Create(todo)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	fmt.Printf("%+v", todo)

	return ctx.JSON(todo)
}
