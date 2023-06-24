package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"hotinplano.com/hot"
	"strings"
)

func main() {

	//w := hot.GetWeather()
	//
	//fmt.Printf("%+v\n", w)
	//
	//e := hot.GetErcotData()
	//
	//fmt.Printf("%+v\n", e)

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {

		w := hot.GetWeather()
		e := hot.GetErcotData()
		style := "green"
		state := strings.ToUpper(e.State)
		if state != "NORMAL" {
			style = "orange"
		}

		return c.Render("index", fiber.Map{
			"Temperature": fmt.Sprintf("%.1f", w.Temperature),
			"FeelsLike":   fmt.Sprintf("%.1f", w.FeelsLike),
			"Humidity":    w.Humidity,
			"State":       state,
			"PRC":         e.Prc,
			"Note":        e.Note,
			"Style":       style,
		})
	})

	app.Listen(":8080")
}
