package controller

import (
	"assignment-3/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ReadAndUpdateWeather(ctx *gin.Context) {
	var weather model.Weather

	file, err := os.OpenFile("assets/weather.json", os.O_RDWR|os.O_APPEND, 0666)
	defer func() {
		errFile := file.Close()
		if errFile != nil {
			if err != nil {
				log.Println(err.Error())
			}
		}
	}()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	weatherByte, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(weatherByte, &weather)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	weather.Status.Water = generateStatus(rand.Intn(100), "water")
	weather.Status.Wind = generateStatus(rand.Intn(100), "wind")

	newWeatherByte, err := json.MarshalIndent(&weather, "", "    ")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ioutil.WriteFile("assets/weather.json", newWeatherByte, 0666)

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"wind":           weather.Status.Wind,
		"water":          weather.Status.Water,
		"title":          "Status Air dan Angin",
		"updateInterval": 15,
	})
}

func generateStatus(percentage int, statusType string) (res int) {
	const maxBahaya = 100
	var aman, minSiaga, maxSiaga, minBahaya int
	if statusType == "water" {
		minSiaga = 6
		maxSiaga = 8
		aman = minSiaga - 1
		minBahaya = maxSiaga + 1
	} else {
		minSiaga = 7
		maxSiaga = 15
		aman = minSiaga - 1
		minBahaya = maxSiaga + 1
	}
	if percentage < 50 {
		res = rand.Intn(aman)
	} else if percentage >= 50 && percentage < 80 {
		res = rand.Intn(maxSiaga-minSiaga) + minSiaga
	} else {
		res = rand.Intn(maxBahaya-minBahaya) + minBahaya
	}
	return
}
