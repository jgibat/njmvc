package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Location struct {
	LocationId    int
	FirstOpenSlot string
}

func main() {
	resp, err := http.Get("https://telegov.njportal.com/njmvc/AppointmentWizard/7")
	if err != nil {
		fmt.Errorf("error while trying to GET appointments %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(body))
	r := regexp.MustCompile(`var timeData = (.*)`)
	timeData := r.FindStringSubmatch(string(body))
	fmt.Print(timeData[1])

	var Locations []Location
	err = json.Unmarshal([]byte(timeData[1]), &Locations)
	if err != nil {
		log.Fatalf("error unmarshaling response %s\n", err)
	}
	for _, l := range Locations {
		fmt.Printf("location %d, appointments %s\n", l.LocationId, l.FirstOpenSlot)
	}
}
