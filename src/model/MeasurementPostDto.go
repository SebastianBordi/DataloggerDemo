package model

type MeasurementPostDto struct {
	Temperature float64 `json:"temperature"`
	Humedity    int     `json:"humedity"`
	Mac         string  `json:"mac"`
	Password    string  `json:"password"`
}
