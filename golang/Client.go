package api

// TODO - The API Key needs to be passed as Header parameter with name Authorization and value should be in the format apiKey <Value>. e.g. apiKey e5f18450-205a-48eb-8589-7d49edaea813

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"time"
)

// Client TODO Doc Comment
type Client struct {
	BaseURL          url.URL
	UserAgent        string
	httpClient       *http.Client
	CachedToken      string
	TerraformVersion string
}

// Init TODO Doc Comment
func (client *Client) Init(region string, terraformVersion string) error {

	baseURL, err := url.Parse("https://apps.sematext.com")
	if err != nil {
		return err
	}

	switch region {
	case "US":
		baseURL, err = url.Parse("https://apps.sematext.com")
	case "EU":
		baseURL, err = url.Parse("https://apps.eu.sematext.com")
	default:
		err = errors.New("sematext_region must be one of EU, US")
	}

	if err != nil {
		return err
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		DisableKeepAlives:     true,
		TLSClientConfig:       &tls.Config{},
	}

	client.BaseURL = *baseURL
	client.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraformVersion)
	client.httpClient = &http.Client{Transport: transport}
	client.CachedToken = ""
	client.TerraformVersion = terraformVersion

	return nil
}

// SetAuthorization TODO Doc Comment
func (client *Client) SetAuthorization(token string) error {
	re := regexp.MustCompile(`[0-9]{8}-[0-9]{4}-[0-9]{4}-[0-9]{4}-[0-9]{8}`)
	if re.Match([]byte(token)) {
		client.CachedToken = fmt.Sprintf("apiKey %s", token)
		return nil
	}
	client.CachedToken = ""
	return errors.New("Bad or missing Token")
}

// GetJSON TODO Doc Comment
func (client *Client) GetJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	if client.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := client.BaseURL
	route.Path = path

	req, err := http.NewRequest("GET", route.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", client.CachedToken)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		genericAppResponse := &GenericAPIResponse{}
		err = json.NewDecoder(res.Body).Decode(genericAppResponse)
		if err != nil {
			return nil, err

		}
		return genericAppResponse, nil
	}

	return nil, fmt.Errorf("Unexected status (%d) return from Sematext API", res.StatusCode)
}

// PutJSON TODO Doc Comment
func (client *Client) PutJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	if client.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := client.BaseURL
	route.Path = path

	jsn, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", route.String(), bytes.NewBuffer(jsn))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", client.CachedToken)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		genericAppResponse := new(GenericAPIResponse)
		err = json.NewDecoder(res.Body).Decode(genericAppResponse)
		if err != nil {
			return nil, err

		}
		return genericAppResponse, nil
	}

	return nil, fmt.Errorf("Unexected status (%d) return from Sematext API", res.StatusCode)
}

// PostJSON TODO Doc Comment
func (client *Client) PostJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	if client.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := client.BaseURL
	route.Path = path

	jsn, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", route.String(), bytes.NewBuffer(jsn))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", client.CachedToken)

	res, err := client.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		genericAppResponse := &GenericAPIResponse{}
		err = json.NewDecoder(res.Body).Decode(genericAppResponse)
		if err != nil {
			return nil, err

		}
		return genericAppResponse, nil
	}

	return nil, fmt.Errorf("Unexected status (%d) return from Sematext API", res.StatusCode)
}

// Delete TODO Doc Comment
func (client *Client) Delete(path string) error {

	if client.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := client.BaseURL
	route.Path = path

	req, err := http.NewRequest("DELETE", route.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", client.CachedToken)

	res, err := client.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("Unexected status (%d) return from Sematext API on Delete", res.StatusCode)
	}
	return nil
}
