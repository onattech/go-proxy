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
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(getRates())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	log.Fatal(app.Listen(":" + port))
}

type ExRate struct {
	USD  BuySell
	Gold BuySell
}

type BuySell struct {
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
}

func getRates() ExRate {
	// Make HTTP GET request
	response, err := http.Get("https://www.kuveytturk.com.tr/finans-portali/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Get the response body as a string
	dataInBytes, _ := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	buyUSD, sellUSD := getEXRate(pageContent, "USD (Amerikan Doları)")
	buyGold, sellGold := getEXRate(pageContent, "ALT (Gram Altın)")

	resp := ExRate{
		USD:  BuySell{buyUSD, sellUSD},
		Gold: BuySell{buyGold, sellGold},
	}

	resPretty := utils.PrettyStruct(resp)
	log.Println("Response: ", resPretty)

	return resp
}

func getEXRate(pageContent string, currencyPhrase string) (float64, float64) {
	// Find currency index
	currencyIndex := strings.Index(pageContent, currencyPhrase)

	// Find Buy index and get the chunk buy/sell are in
	buyIndex := strings.Index(pageContent[currencyIndex:], "<p>") + currencyIndex
	chunk := pageContent[buyIndex : buyIndex+800]

	// Create a regular expression to find buySell
	re := regexp.MustCompile(`[0-9]*,[0-9]*`)

	// Extract buy/sell from the chunk
	buysell := re.FindAllString(string(chunk), -1)

	buy := utils.EasyFloat(buysell[0])
	sell := utils.EasyFloat(buysell[1])

	return buy, sell
}
