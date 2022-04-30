package coingeckoapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	client *http.Client
}

func New() *Client {
	hc := &http.Client{
		Timeout: 10 * time.Second,
	}
	return &Client{
		client: hc,
	}
}

func (c *Client) do(product, method, path string, data interface{}, sign bool, stream bool) (response []byte, err error) {
	var ENDPOINT string
	switch product {
	case "spot":
		ENDPOINT = "https://api.coingecko.com/api/v3"
	default:
		// pass
	}
	values, err := query.Values(data)
	if err != nil {
		return nil, err
	}
	payload := values.Encode()

	var req *http.Request
	if method == http.MethodGet {
		req, err = http.NewRequest(method, fmt.Sprintf("%s/%s?%s", ENDPOINT, path, payload), nil)
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%s/%s", ENDPOINT, path), strings.NewReader(payload))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	//req.Header.Add("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d: %v", resp.StatusCode, string(response))
	}
	return response, err
}

func TimeFromUnixTimestampInt(raw interface{}) (time.Time, error) {
	ts, ok := raw.(int64)
	if !ok {
		return time.Time{}, errors.New(fmt.Sprintf("unable to parse, value not int64: %T", raw))
	}
	return time.Unix(0, ts*int64(time.Millisecond)), nil
}
