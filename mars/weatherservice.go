package mars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"serverless-golang-api-with-aws/appsettings"
	"serverless-golang-api-with-aws/structs"
	"time"
)

var nasaMarsSolsWeatherBaseAPIURL string
var marsWeatherAPIKey string
var marsWeatherAPIFeedbackType string
var marsWeatherAPIVersion string

func init() {
	nasaMarsSolsWeatherBaseAPIURL = "https://api.nasa.gov/insight_weather/"
	marsWeatherAPIKey = appsettings.GetFromEnvironment("MARS_WEATHER_API_KEY", "DEMO_KEY")
	marsWeatherAPIFeedbackType = "json"
	marsWeatherAPIVersion = appsettings.GetFromEnvironment("MARS_WEATHER_API_VER", "1.0")
}

func GetAllSolsWeather() (structs.FriendlyResponse, error) {
	client := getHttpClient()
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf(nasaMarsSolsWeatherBaseAPIURL+"?api_key=%s&feedType=%s&ver=%s", marsWeatherAPIKey, marsWeatherAPIFeedbackType, marsWeatherAPIVersion), nil)
	if err != nil {
		return structs.FriendlyResponse{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return structs.FriendlyResponse{}, err
	}

	responseJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return structs.FriendlyResponse{}, err
	}

	var solsWeatherInfo structs.SolWeatherResponse
	if err := json.Unmarshal(responseJson, &solsWeatherInfo); err != nil {
		return structs.FriendlyResponse{}, err
	}

	return solsWeatherInfo.ToFriendlyResponse(), err

}

func getHttpClient() *http.Client {
	tr := &http.Transport{
		IdleConnTimeout: 120 * time.Second,
	}
	return &http.Client{Transport: tr}
}
