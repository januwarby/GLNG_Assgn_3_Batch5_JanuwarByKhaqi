package main

import (
	"GLNG_Assgn_3_Batch5_JanuwarByKhaqi/data"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	dataInstance := data.NewData()
	go dataInstance.Update()

	for {
		output := map[string]interface{}{
			"water": dataInstance.Water,
			"wind":  dataInstance.Wind,
		}
		statusWater := fmt.Sprintf("status water: %s", dataInstance.StatusWater)
		statusWind := fmt.Sprintf("status wind: %s", dataInstance.StatusWind)

		jsonOutput, _ := json.MarshalIndent(output, "", "")

		fmt.Printf("%s\n%s\n%s\n", string(jsonOutput), statusWater, statusWind)

		time.Sleep(2 * time.Second)
	}
}
