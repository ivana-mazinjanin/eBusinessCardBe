package controller

import (
	"api/src/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	types "api/src/types"

	"github.com/gorilla/mux"
)

type PlacesController struct{}

func (p PlacesController) GetPlaceDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// read request
		vars := mux.Vars(r)
		placeId := vars["placeId"]

		// call upstream api
		res, err := CallUpstreamAPI(placeId)
		if err != nil {

			// TODO: Return API Error

			// error.ApiError(w, http.StatusInternalServerError, "Failed to get place details! \n"+err.Error())
			return
		}

		// prepare response
		helpers.RespondWithJSON(w, res)
	}
}

func CallUpstreamAPI(placeId string) (*types.PlaceDetails, error) { // TODO: return error

	apiURL := "https://storage.googleapis.com/coding-session-rest-api/" + placeId

	// Make a GET request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}

	// Check if the request was successful (status code 200)
	if response.StatusCode != http.StatusOK {
		//return "", fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res types.PlaceDetails

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	// Convert the response body to a string and return it
	return &res, nil
}
