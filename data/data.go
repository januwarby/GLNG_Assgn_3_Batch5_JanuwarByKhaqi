package data

import (
	"GLNG_Assgn_3_Batch5_JanuwarByKhaqi/utils"
	"math/rand"
	"time"
)

type Data struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

func NewData() *Data {
	return &Data{}
}

func (d *Data) Update() {
	for {
		d.Water = rand.Intn(100) + 1
		d.Wind = rand.Intn(100) + 1
		d.StatusWater = utils.GetStatus(d.Water, "water")
		d.StatusWind = utils.GetStatus(d.Wind, "wind")

		time.Sleep(2 * time.Second)
	}
}
