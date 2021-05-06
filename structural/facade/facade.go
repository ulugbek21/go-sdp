package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// CurrentWeatherDataRetriever ...
type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

// CurrentWeatherData ...
type CurrentWeatherData struct {
	APIkey string
}

// GetByGeoCoordinates ...
func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, c.APIkey))
}

// GetByCityAndCountryCode ...
func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s", city, countryCode, c.APIkey))
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("status code was %d, aborting. Error message was:\n%s", resp.StatusCode, errMsg)

		return
	}

	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()

	return
}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}
