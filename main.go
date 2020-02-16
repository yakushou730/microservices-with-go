package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type timeZoneConversion struct {
	TimeZone    string
	CurrentTime string
}

var conversionMap = map[string]string{
	"ASR": "-3h",
	"EST": "-5h",
	"BST": "+1h",
	"IST": "+5h30m",
	"HKT": "+8h",
	"ART": "-3h",
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	timeZone := r.URL.Query().Get("tz")

	timeDifference, _ := conversionMap[timeZone]

	currentTimeConverted, _ := getCurrentTimeByTimeDifference(timeDifference)

	tzc := &timeZoneConversion{
		TimeZone:    timeZone,
		CurrentTime: currentTimeConverted,
	}
	jsonResponse, _ := json.Marshal(tzc)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonResponse))
}

func getCurrentTimeByTimeDifference(timeDifference string) (string, error) {
	now := time.Now().UTC()
	difference, err := time.ParseDuration(timeDifference)
	if err != nil {
		return "", nil
	}
	now = now.Add(difference)
	fmt.Println(now)
	return now.Format("15:04:05"), nil
}
