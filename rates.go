package gorbkrates

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Now get current rate. Show ISO 4217
func Now(number string) (rate float64, err error) {
	return OnDate(number, time.Now())
}

// OnDate on date rate. Show ISO 4217
func OnDate(number string, date time.Time) (rate float64, err error) {

	http, err := http.Get("http://cbrates.rbc.ru/tsv/" + number + "/" + date.Format("2006/01/02") + ".tsv")
	if http.StatusCode != 200 || err != nil {
		return 0, err
	}

	defer http.Body.Close()

	b, err := ioutil.ReadAll(http.Body)
	if err != nil {
		return 0, err
	}

	arr := strings.Split(string(b), "\n")
	if len(arr) == 0 {
		return 0, errors.New("Empty .tsv")
	}

	res := strings.Split(arr[0], "\t")
	if len(res) == 0 {
		return 0, errors.New("Error Split string")
	}

	return strconv.ParseFloat(res[1], 64)
}

// ForPeriod rate for period. Show ISO 4217
func ForPeriod(number string, start, end time.Time) (rates map[time.Time]float64, err error) {

	http, err := http.Get("http://cbrates.rbc.ru/tsv/cb/" + number + ".tsv")
	if http.StatusCode != 200 || err != nil {
		return nil, err
	}

	defer http.Body.Close()

	b, err := ioutil.ReadAll(http.Body)
	if err != nil {
		return nil, err
	}

	arr := strings.Split(string(b), "\n")
	if len(arr) == 0 {
		return nil, errors.New("Empty .tsv")
	}

	result := map[time.Time]float64{}

	for i := len(arr)-1; i > 0; i-- {
		
		res := strings.Split(arr[i], "\t")
		t, _ := time.Parse("20060102", res[0])


		if !(start.Before(t) && end.After(t))	{
			continue
		}

		c, _ := strconv.ParseFloat(res[2], 64)

		result[t] = c

	}
	
	return result, nil
}

