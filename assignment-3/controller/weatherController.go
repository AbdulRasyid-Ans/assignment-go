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

	weather.Status.Water = rand.Intn(100)
	weather.Status.Wind = rand.Intn(100)

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
