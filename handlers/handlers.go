package handlers

import (
	"GLNG_Assgn_3_Batch5_JanuwarByKhaqi/data"
	"encoding/json"
	"net/http"
	"time"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	dataInstance := data.NewData()
	go dataInstance.Update()

	for {
		err := json.NewEncoder(w).Encode(dataInstance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		time.Sleep(2 * time.Second)
	}
}
