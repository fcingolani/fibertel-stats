package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	log.SetOutput(&lumberjack.Logger{
		Filename: "./fibertel-stats.log",
		MaxSize:  50,
	})

	logCurrentValues()

	for _ = range time.Tick(time.Minute) {
		logCurrentValues()
	}

}

func logCurrentValues() {

	values, err := getCurrentValues()

	if err != nil {
		log.Println("✕", err.Error())
		fmt.Print("✕")
	} else {
		log.Println("✓", values)
		fmt.Print("✓")
	}

}

func getCurrentValues() ([]string, error) {
	url := "http://provisioning.fibertel.com.ar/asp/nivelesPrima.asp"

	timeout := time.Duration(3 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	bodyString := string(bodyBytes)

	lines := strings.Split(bodyString, "\n")

	values := make([]string, 0)

	for _, line := range lines {

		if strings.Contains(line, "valor") {
			line = strings.Split(line, ">")[1]
			line = strings.Split(line, " ")[0]
			line = strings.Replace(line, ",", ".", 1)
			_, err := strconv.ParseFloat(line, 64)

			if err == nil {
				values = append(values, line)
			}

		}

	}

	return values, nil

}
