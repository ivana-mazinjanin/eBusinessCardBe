package controller

import (
	"api/src/helpers"
	"api/src/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type PlacesController struct{}

func (p PlacesController) GetPlaceDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		error := helpers.CustomError{}

		// read request
		vars := mux.Vars(r)
		placeId := vars["placeId"]

		// call upstream 	API
		apiURL := "https://storage.googleapis.com/coding-session-rest-api/" + placeId
		response, err := http.Get(apiURL)
		if err != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed to fetch place data!")
			return
		}

		// Check if the request was successful (status code 200)
		if response.StatusCode != http.StatusOK {
			error.ApiError(w, response.StatusCode, "Failed to fetch place data!")
			return
		}

		daysOfWeek := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

		// Read upstream data and prepare response
		placeDetails, err := readUpstreamResponse(response)
		out := types.PlaceDetailsOut{
			Name:         placeDetails.Name,
			Address:      placeDetails.Address,
			IsOpen:       isOpen(placeDetails.OpeningHours),
			NextChange:   nextChange(placeDetails.OpeningHours, daysOfWeek),
			OpeningHours: []*types.OpeningHoursOut{},
		}

		for _, dayOfTheWeek := range daysOfWeek {

			a := placeDetails.OpeningHours.Days[dayOfTheWeek]

			isSet := false

			for _, x := range out.OpeningHours {
				if areEqualHours(x.WorkingBlocks, a) {
					x.Days = append(x.Days, dayOfTheWeek)
					isSet = true
					break
				}
			}

			if isSet == false {
				newEntry := types.OpeningHoursOut{
					Days:          []string{dayOfTheWeek},
					WorkingBlocks: a,
				}
				out.OpeningHours = append(out.OpeningHours, &newEntry)
			}
		}

		// prepare response
		helpers.RespondWithJSON(w, out)
	}
}

func readUpstreamResponse(response *http.Response) (*types.PlaceDetails, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var placeDetails types.PlaceDetails

	err = json.Unmarshal(body, &placeDetails)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	return &placeDetails, nil
}

func areEqualHours(blocksA []types.WorkingBlock, blocksB []types.WorkingBlock) bool {
	if len(blocksA) != len(blocksB) {
		return false
	}

	for i := 0; i < len(blocksA); i++ {
		// TODO: Compare Type ?
		if blocksA[i].Start != blocksB[i].Start || blocksA[i].End != blocksB[i].End {
			return false
		}
	}

	return true
}

func isOpen(openingHours types.OpeningHours) bool {
	// now := time.Now()
	now := time.Date(2023, time.Month(11), 27, 18, 35, 30, 0, time.UTC)

	currentDay := strings.ToLower(now.Weekday().String())
	currentTimeStr := now.Format("15:04")

	blocks := openingHours.Days[currentDay]

	if blocks == nil {
		return false
	}

	for _, block := range blocks {
		if currentTimeStr >= block.Start && currentTimeStr < block.End {
			return true
		}
	}

	return false
}

func nextChange(openingHours types.OpeningHours, daysOfWeek []string) string {
	// now := time.Now()
	now := time.Date(2023, time.Month(11), 27, 18, 35, 30, 0, time.UTC)
	currentDay := strings.ToLower(now.Weekday().String())
	currentTimeStr := now.Format("15:04")

	blocks := openingHours.Days[currentDay]

	if blocks != nil {
		for _, block := range blocks {

			// if restaurtant is open, find closing time
			if currentTimeStr >= block.Start && currentTimeStr < block.End {
				return block.End
			}

			// if next opening block found during current day, return opening time
			if currentTimeStr < block.Start {
				return block.Start
			}
		}
	}

	// find next day
	nextDay := 0
	for i, day := range daysOfWeek {
		if currentDay == day {
			nextDay = (i + 1) % 7
		}
	}

	// go day by day as long as next working day is found (needed in case place is closed next day...)
	var blocksNextDay []types.WorkingBlock = nil
	for {
		blocksNextDay = openingHours.Days[daysOfWeek[nextDay]]
		if blocksNextDay != nil {
			break
		}
		nextDay = (nextDay + 1) % 7
	}

	return blocksNextDay[0].Start

}
