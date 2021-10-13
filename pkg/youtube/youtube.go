package youtube

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Items []Items `json:"items"`
}

const URL = "https://www.googleapis.com/youtube/v3/channels"

func TryRequest(config Config) ([]Items, error) {

	var response Response

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", config.YoutubeKey)
	q.Add("id", config.ChannelId)
	q.Add("part", config.Part)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}