package utils

import (
    //"fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func DoHttpRequest ( apiEndpoint string ) string {
    apiToken := "RGAPI-80136d22-3d28-43ff-a80c-fee11ce34569"
    req, err := http.NewRequest("GET", apiEndpoint, nil)
    if err != nil {
        log.Fatal("Error generating HTTP request")
    }
    req.Header.Set("X-Riot-Token", apiToken)

    response, err := http.DefaultClient.Do(req)
    defer response.Body.Close()
    if err != nil {
        log.Fatal("Error doing the HTTP request")
    }

    if response.StatusCode != http.StatusOK {
        log.Fatal("Error in your request: %s", response.StatusCode)
    }

    bodyBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    return string(bodyBytes)
}