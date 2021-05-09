# Facade Pattern

A facade, in architectural terms, is the front wall that hides the rooms and corridors of a building. It protects its inhabitants from cold and rain, and provides them privacy. It orders and divides the dwellings.

The Facade design pattern does the same, but in our code. It shields the code from
unwanted access, orders some calls, and hides the complexity scope from the user. (Contreras, 2017)

## Implementation

Use of HTTP REST API with abstracting complex handling of HTTP requests.

For example, we might have the following method:

```go
c.doRequest(...)
```

Which might do a _magic_ under the hood:

```go
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
```

## Summary

The Facade Pattern might be used on a package (library) level to abstract complex operations (Contreras, 2017)
