package model

type MeasurementPostDto struct {
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
	Mac         string  `json:"mac"`
	Password    string  `json:"password"`
}
