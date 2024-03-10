package apinTeraction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ConfigUrls struct {
	Next     string
	Previous string
}

type locationAreaData struct {
	Next               *string `json:"next"`
	Previous           *string `json:"previous"`
	Location_AreaNames []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (urlHolder *ConfigUrls) FetchLocationAreas(flag bool) ([]string, error) {

	// Making a custom client for setting timeout on api call
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Making Get request to location-area endpoint
	var resp *http.Response
	var Err error
	if flag {
		resp, Err = client.Get(urlHolder.Next)
	} else {
		resp, Err = client.Get(urlHolder.Previous)
	}

	// Error checking for api call
	if Err != nil {
		return nil, fmt.Errorf("failed to fetch: %v", Err)
	}

	body, err0 := io.ReadAll(resp.Body)
	resp.Body.Close()

	// Error checking for reading the response
	if err0 != nil {
		return nil, fmt.Errorf("error occured in reading from fetched data, error in deatils: %v", err0)
	}

	// If failed to fetch location_area names
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("response falied with status code: %d and body: %s ", resp.StatusCode, body)
	}

	// Converting json response from api call to go struct
	data := locationAreaData{}
	err1 := json.Unmarshal(body, &data)

	// Error checking for conversion of json to struct process
	if err1 != nil {
		return nil, fmt.Errorf("can't convert json to locationData struct, Erro details: %v", err1)
	}

	// Updating configUrls struct's instance, here first checking for null pointer
	if data.Next != nil {
		urlHolder.Next = *(data.Next)
	}
	if data.Previous != nil {
		urlHolder.Previous = *(data.Previous)
	}

	// Loading the nameList slice
	nameList := make([]string, 0, 20) // Here slice capacity choosed to be 20 because there will be 20 results in a single json
	for _, Location_AreaName := range data.Location_AreaNames {
		nameList = append(nameList, Location_AreaName.Name)
	}

	return nameList, nil

}
