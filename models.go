package main

import "time"

type Data struct {
	CreatedAt   time.Time `json:"created_at"`
	Temperature float64   `json:"temperature"`
	Luminosity  float64   `json:"luminosity"`
	Pressure    float64   `json:"pressure"`
	Humidity    float64   `json:"humidity"`
}
