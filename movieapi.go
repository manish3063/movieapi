package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setupRoutes(r *gin.Engine) {
	r.GET("/movies/year/:year", Dummyyear)
	r.GET("/movies/rating/:rating", Dummyrating)
	r.GET("/movies/genre/:genre", Dummygenre)

}

//Dummy3 function
func Dummyyear(c *gin.Context) {
	year, ok := c.Params.Get("year")
	records := readCsvFile("./movies.csv")
	movies := getyear(records, year)
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	/*
		city := ""
	*/
	res := gin.H{
		"year":       year,
		"movie name": movies,
	}
	c.JSON(http.StatusOK, res)
}

func getyear(records [][]string, input string) []string {
	//var app [][] string

	//var year string
	var movies = []string{}
	for i := 0; i < 8; i++ {

		//fmt.Println(records[0][0], i)
		if records[i][7] == input {
			//year = records[i][7]
			//movies:=records[i][0]
			movies = append(movies, records[i][0])

		}

	}
	return movies
}

//
//Dummy rating...
func Dummyrating(c *gin.Context) {
	rating, ok := c.Params.Get("rating")
	records := readCsvFile("./movies.csv")
	movies := getrating(records, rating)
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	/*
		city := ""

	*/
	res := gin.H{
		"rating":     rating,
		"movie name": movies,
	}
	c.JSON(http.StatusOK, res)
}

func getrating(records [][]string, input string) []string {
	//var app [][] string

	//var year string
	var movies = []string{}
	for i := 0; i < 8; i++ {

		//fmt.Println(records[0][0], i)
		if records[i][5] <= input {
			//year = records[i][7]
			//movies:=records[i][0]
			movies = append(movies, records[i][0])

		}

	}
	return movies
}

//
//Dummy3 genre
func Dummygenre(c *gin.Context) {
	genre, ok := c.Params.Get("genre")
	records := readCsvFile("./movies.csv")
	movies := getgenre(records, genre)
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	/*
		city := ""
	*/
	res := gin.H{
		"genre":      genre,
		"movie name": movies,
	}
	c.JSON(http.StatusOK, res)
}

func getgenre(records [][]string, input string) []string {
	//var app [][] string

	//var year string
	var genre = []string{}
	for i := 0; i < 8; i++ {

		//fmt.Println(records[0][0], i)
		if records[i][1] == input {
			//year = records[i][7]
			//genre:=records[i][0]
			genre = append(genre, records[i][0])

		}

	}
	return genre
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
