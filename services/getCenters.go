package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/udaya2899/covid-vaccine-notify/model"
)

const RequiredMinimumAge = 18

func GetCenters(districtID, date string) (model.CowinResponse, error) {
	fmt.Printf("\n\n******\n\nCOVID Vaccine Notify\n\n******\n\n")

	cowinResponse, err := getCentersFromAPI(districtID, date)
	if err != nil {
		return model.CowinResponse{}, err
	}

	log.Printf("COWIN Response obtained")

	result := queryJSON(cowinResponse)

	return result, nil
}

func queryJSON(data model.CowinResponse) model.CowinResponse {
	result := model.CowinResponse{}
	for _, center := range data.Centers {
		for _, session := range center.Sessions {
			if session.MinAgeLimit == RequiredMinimumAge {
				result.Centers = append(result.Centers, center)
			}
		}
	}

	return result
}

func getCentersFromAPI(districtID, date string) (model.CowinResponse, error) {
	url := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/calendarByDistrict?district_id=%s&date=%s", districtID, date)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return model.CowinResponse{}, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.CowinResponse{}, err
	}

	cowinResponse := model.CowinResponse{}
	jsonErr := json.Unmarshal(body, &cowinResponse)
	if jsonErr != nil {
		return model.CowinResponse{}, err
	}

	return cowinResponse, nil
}
