package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Weather struct {
  
  Location struct {
    Region string `json:"region"`
    Country string `json:"country"`
    Timezone string `json:"tz_id"`
    Localtime string `json:"localtime"`
  } `json:"location"`
  Current struct {
    Temperature float64 `json:"temp_c"`
    Condition struct{
      Text string `json:"text"`
    }`json:"condition"`
  }`json:"current"`

}

func main() {
  url :=  "insert -- url"
  weather := Weather{}
  response , err := http.Get(url)

  if err!= nil {
    log.Fatal(err)
  }

  defer response.Body.Close()

  json.NewDecoder(response.Body).Decode(&weather)


  fmt.Printf("|--LOCATION : %s %s                \n",weather.Location.Country,weather.Location.Region)
  fmt.Printf("|--Timezone : %s                \n",weather.Location.Timezone)
  fmt.Printf("|--Current Weather: %s                  \n",weather.Current.Condition.Text)
  fmt.Printf("|--Temperature : %v°C                    \n",weather.Current.Temperature)
  
  i := ""
  fmt.Println("press q to exit")
  fmt.Scan(&i)

  

}

