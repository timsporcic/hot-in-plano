package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"hotinplano.com/hot"
	"strings"
	"sync"
	"time"
)

type CachedValues struct {
	Timestamp   int64
	Temperature float64
	Humidity    int
	FeelsLike   float64
	Note        string
	State       string
	Prc         string
	mu          sync.Mutex
}

const fiveMin = 60 * 5

func main() {

	values := &CachedValues{}
	updateWeather(values)

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {

		style := "green"
		state := strings.ToUpper(values.State)
		if state != "NORMAL" {
			style = "orange"
		}

		if time.Now().Unix() > values.Timestamp+fiveMin {
			updateWeather(values)
		}

		return c.Render("index", fiber.Map{
			"Temperature": fmt.Sprintf("%.1f", values.Temperature),
			"FeelsLike":   fmt.Sprintf("%.1f", values.FeelsLike),
			"Humidity":    values.Humidity,
			"State":       state,
			"PRC":         values.Prc,
			"Note":        values.Note,
			"Style":       style,
		})
	})

	app.Listen(":8080")
}

func updateWeather(values *CachedValues) {
	values.mu.Lock()
	defer values.mu.Unlock()
	if time.Now().Unix() > values.Timestamp+fiveMin {
		w := hot.GetWeather()
		e := hot.GetErcotData()
		values.Humidity = w.Humidity
		values.FeelsLike = w.FeelsLike
		values.Temperature = w.Temperature
		values.State = e.State
		values.Prc = e.Prc
		values.Note = e.Note
		values.Timestamp = time.Now().Unix()
	}
}
