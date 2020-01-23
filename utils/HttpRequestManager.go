package utils

import (
    //"fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func DoHttpRequest ( apiEndpoint string ) string {
    apiToken := "RGAPI-5f0536dc-e07f-4970-a64e-2162b8affd4f"
    req, err := http.NewRequest("GET", apiEndpoint, nil)
    if err != nil {
        log.Fatal("Error generating HTTP request: \n%s", apiEndpoint)
    }
    req.Header.Set("X-Riot-Token", apiToken)

    response, err := http.DefaultClient.Do(req)
    defer response.Body.Close()
    if err != nil {
        log.Fatal("Error doing the HTTP request")
    }

    if response.StatusCode != http.StatusOK {
        log.Fatal("Error in your request: ", apiEndpoint, " Error code: ", response.Body)
    }

    bodyBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    return string(bodyBytes)
}