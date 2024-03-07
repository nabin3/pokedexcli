package apinTeractoion


import (
    "fmt"
    "encoding/json"
    "net/http"
)


type ConfigUrls struct {
    Next string
    Previous string
}


type locationAreaData struct {
    Next string `json:"next"`
    Previous string `json:"previous"`
    location_AreaNames []struct {
        Name string `json:"name"`
    } `json:"results"`
}


func (urlHolder *ConfigUrls)FetchLocationAreas(flag bool) ([]string, error) {

    var resp *http.Response
    var Err error
    if flag {
        resp, Err = http.Get(urlHolder.Next)
    } else {
        resp, Err = http.Get(urlHolder.Previous)
    }

    if Err != nil {
        return nil, fmt.Errorf("Failed to fetch: %v", Err)
    }

    body, err0 := io.ReadAll(resp.Body)
    resp.Body.Close()

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("Response falied with status code: %d and body: %s ", resp.statusCode, body)
    }

    if err0 != nil {
        return nil, fmt.Errorf("Error occured in reading from fetched data, error in deatils: %v", err0)
    }

    data := locationAreaData{}
    err1 := json.Unmarshal(body, &data)

    if err1 != nil {
        return nil, fmt.Errorf("Can't convert json to locationData struct, Erro details: %v", err1)
    }

    urlHolder.Next = data.Next
    urlHolder.Previous = data.Previous

    nameList := make([]string, 0, 20)  // Here slice capacity choosed to be 20 because there will be 20 results in a single json
    for _, location_AreaName := range data.location_AreaNames {
        nameList = append(nameList, location_AreaName.Name)
    }

    return nameList, nil

}
