package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Etihad_body struct {
	CabinClass     string        `json:"cabinClass"`
	AwardBooking   string        `json:"awardBooking"`
	PromoCodes     []interface{} `json:"promoCodes"`
	SearchType     string        `json:"searchType"`
	ItineraryParts []struct {
		From struct {
			UseNearbyLocations bool   `json:"useNearbyLocations"`
			Code               string `json:"code"`
		} `json:"from"`
		To struct {
			UseNearbyLocations bool   `json:"useNearbyLocations"`
			Code               string `json:"code"`
		} `json:"to"`
		When struct {
			Date string `json:"date"`
		} `json:"when"`
	} `json:"itineraryParts"`
	Passengers struct {
		Adt int `json:"ADT"`
	} `json:"passengers"`
}

func Generate_body(from string, to string, when string) []byte {
	f, err := os.Open("./json/etihad_body.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	var body Etihad_body

	json.Unmarshal(byteValue, &body)
	body.ItineraryParts[0].From.Code = from
	body.ItineraryParts[0].To.Code = to
	body.ItineraryParts[0].When.Date = when

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// type airport struct {
// 	useNearbyLocations bool   `json:"useNearbyLocations"`
// 	code               string `json: "code"`
// }

// type when struct {
// 	date string `json: "date"`
// }

// type itinerary struct {
// 	from airport `json: "from"`
// 	to   airport `json: "to"`
// 	when when    `json: "when"`
// }

// type passenger_data struct {
// 	ADT int `json:"ADT"`
// }

// type Etihad_body struct {
// 	cabinClass     string         `json:"cabinClass"`
// 	awardBooking   string         `json:"awardBooking"`
// 	promoCodes     []string       `json:"promoCodes"`
// 	searchType     string         `json:"searchType"`
// 	itineraryParts []itinerary    `json:"itineraryParts"`
// 	passengers     passenger_data `json:"passengers"`
// }

// func Generate_body(from string, to string, when string) {
// f, err := os.Open("./json/etihad_body.json")
// if err != nil {
// 	log.Fatal(err)
// }
// defer f.Close()

// 	byteValue, _ := ioutil.ReadAll(f)

// 	var body Etihad_body

// 	json.Unmarshal(byteValue, &body)

// 	fmt.Println(body.cabinClass)
// 	fmt.Println(byteValue)
// }
