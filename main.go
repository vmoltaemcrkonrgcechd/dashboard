package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
	"math/rand"
	"time"
)

const (
	amount = 5
	min    = 128
	max    = 2048
	delay  = 2
)

type response struct {
	LineChartData     [amount][amount]int `json:"lineChartData"`
	BarChartData      [amount][amount]int `json:"barChartData"`
	RadarChartData    [amount][amount]int `json:"radarChartData"`
	DoughnutChartData [amount]int         `json:"doughnutChartData"`
}

func newResponse() response {
	return response{
		LineChartData:     generateMatrix(),
		BarChartData:      generateMatrix(),
		RadarChartData:    generateMatrix(),
		DoughnutChartData: generateArray(),
	}
}

func generateArray() (arr [amount]int) {
	for i := 0; i < amount; i++ {
		arr[i] = rand.Intn(max-min) + min
	}

	return arr
}

func generateMatrix() (matrix [amount][amount]int) {
	for i := 0; i < amount; i++ {
		matrix[i] = generateArray()
	}

	return matrix
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/", websocket.New(func(conn *websocket.Conn) {
		var err error

		for {
			if err = conn.WriteJSON(newResponse()); err != nil {
				log.Println(err)
				return
			}
			time.Sleep(time.Second * delay)
		}
	}))

	log.Fatal(app.Listen(":80"))
}
