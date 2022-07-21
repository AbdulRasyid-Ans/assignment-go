package model

type Weather struct {
	Status WeatherStatus `json:"status"`
}

type WeatherStatus struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
