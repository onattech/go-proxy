package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(getTitle())
	})

	// Make HTTP request
	alisSatis := getTitle()
	log.Println(alisSatis)
	// indexTest()

	log.Fatal(app.Listen(":4000"))
}

// <div class="box-borderless">

func getTitle() string {
	// Make HTTP GET request
	response, err := http.Get("https://www.kuveytturk.com.tr/finans-portali/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	// log.Println(pageContent)

	// Find Dollar index
	dollarIndex := strings.Index(pageContent, "USD (Amerikan Doları)")
	log.Println(dollarIndex, "!!!!!")

	// Find Alis index
	alisIndex := strings.Index(pageContent[dollarIndex:], "<p>") + dollarIndex
	log.Println(alisIndex, "!!!!!")
	chunk := pageContent[alisIndex : alisIndex+800]

	// Set Alis index
	// Create a regular expression to find comments
	re := regexp.MustCompile(`\d\d,\d\d\d\d`)
	comments := re.FindAllString(string(chunk), -1)

	return comments[0] + " " + comments[1]
}

func indexTest() {
	x := "Hello"
	y := x[2:]
	fmt.Println(y)
	// s := "12121211122"
	// first3 := s[0:3]
	// last3 := s[len(s)-3:]
}