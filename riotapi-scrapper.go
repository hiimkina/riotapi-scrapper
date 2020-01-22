package main

import (
    "fmt"
    "net/http"
)

func main () {
    req, err := http.NewRequest("GET", "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/FrFkina", nil)
    if err != nil {
    	// handle err
    }
    req.Header.Set("X-Riot-Token", "RGAPI-80136d22-3d28-43ff-a80c-fee11ce34569")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
    	// handle err
    }

    fmt.Printf("%s", resp)

    defer resp.Body.Close()
}
