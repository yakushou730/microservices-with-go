package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	"RET": "MMMMNUN",
}

func main() {
	http.HandleFunc("/convert", loggingMiddleware(handler))
	http.HandleFunc("/", loggingMiddleware(notFoundHandler))
	http.ListenAndServe("localhost:3000", nil)
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func loggingMiddleware(handler handlerFunc) handlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s - %s", time.Now().Format("2018-11-25 14:32:58"), r.Method, r.URL.String())
		handler(w, r)
	}
	return fn
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Error 400: The requested URL does not exist.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	timeZone := r.URL.Query().Get("tz")
	if timeZone == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error 400: tz query parameter is required.")
		return
	}

	timeDifference, ok := conversionMap[timeZone]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `Error 404: the tz value: "%s" does not correspond to an existing tz value.`, timeZone)
		return
	}

	currentTimeConverted, err := getCurrentTimeByTimeDifference(timeDifference)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Server Error.")
		return
	}

	w.WriteHeader(http.StatusOK)

	tzc := &timeZoneConversion{
		TimeZone:    timeZone,
		CurrentTime: currentTimeConverted,
	}
	jsonResponse, err := json.Marshal(tzc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Server Error.")
		return
	}

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
