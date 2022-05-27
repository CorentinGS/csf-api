package main

import "time"

type Data struct {
	CreatedAt   time.Time `json:"created_at"`
	Temperature float64   `json:"temperature"`
	Luminosity  float64   `json:"luminosity"`
	Wind        float64   `json:"wind"`
	Humidity    float64   `json:"humidity"`
}
