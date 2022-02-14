package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/onattech/go-proxy/utils"
)

func main() {
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		// AllowHeaders: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString(getTitle())
		return c.JSON(getTitle())
	})

	// Make HTTP request
	alisSatis := getTitle()
	log.Println(alisSatis)
	// indexTest()

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}

// <div class="box-borderless">

type ExRate struct {
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
}

func getTitle() ExRate {
	// Make HTTP GET request
	response, err := http.Get("https://www.kuveytturk.com.tr/finans-portali/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Get the response body as a string
	dataInBytes, _ := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	// log.Println(pageContent)

	// Find Dollar index
	dollarIndex := strings.Index(pageContent, "USD (Amerikan Doları)")
	// log.Println(dollarIndex, "!!!!!")

	// Find Buy index
	buyIndex := strings.Index(pageContent[dollarIndex:], "<p>") + dollarIndex
	chunk := pageContent[buyIndex : buyIndex+800]
	// log.Println(buyIndex, "!!!!!")

	// Set Buy index
	// Create a regular expression to find buySell
	re := regexp.MustCompile(`\d\d,\d\d\d\d`)
	buysell := re.FindAllString(string(chunk), -1)
	buy := utils.EasyFloat(buysell[0])
	sell := utils.EasyFloat(buysell[1])

	resp := ExRate{buy, sell}

	return resp

}
