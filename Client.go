package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
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

	fmt.Println("Client.Init Called")
	fmt.Println("-----------------------------")

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

	if IsValidUUID(token) {
		client.CachedToken = token
		return nil
	}

	client.CachedToken = ""
	return errors.New("Bad or missing Sematext API Token")
}

// GetJSON TODO Doc Comment
func (client *Client) GetJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	fmt.Println("GetJSON Called")

	if client.CachedToken == "" {
		return nil, errors.New("coding error - method called without setting API token")
	}

	route := client.BaseURL
	route.Path = path

	req, err := http.NewRequest("GET", route.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("apiKey %s", client.CachedToken))

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return handleAPIResponse(res)

}

// PutJSON TODO Doc Comment
func (client *Client) PutJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	fmt.Println("PutJSON Called")
	fmt.Println("-----------------------------")
	fmt.Println("client.CachedToken")
	fmt.Println(client.CachedToken)
	fmt.Println("-----------------------------")
	fmt.Println("request")
	fmt.Printf("%+v\n", object)
	fmt.Println("-----------------------------")
	if client.CachedToken == "" {
		return nil, errors.New("coding error - method called without setting API token")
	}

	route := client.BaseURL
	route.Path = path

	fmt.Println("-----------------------------")
	fmt.Println("route")
	fmt.Println(route.String())
	fmt.Println("-----------------------------")

	jsn, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", route.String(), bytes.NewBuffer(jsn))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("apiKey %s", client.CachedToken))

	fmt.Println("-----------------------------")
	fmt.Println("req")
	fmt.Printf("--> %s\n\n", formatRequest(req))
	fmt.Println("-----------------------------")

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	fmt.Printf("--> %s\n\n", formatRequest(req))
	fmt.Println("-----------------------------")
	return handleAPIResponse(res)

}

// PostJSON TODO Doc Comment
func (client *Client) PostJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	fmt.Println("PostJSON Called")
	fmt.Println("-----------------------------")
	fmt.Println("client.CachedToken")
	fmt.Println(client.CachedToken)
	fmt.Println("-----------------------------")
	fmt.Println("request")
	fmt.Printf("%+v\n", object)
	fmt.Println("-----------------------------")

	if client.CachedToken == "" {
		return nil, errors.New("coding error - method called without setting API token")
	}

	route := client.BaseURL
	route.Path = path

	fmt.Println("-----------------------------")
	fmt.Println("route")
	fmt.Println(route.String())
	fmt.Println("-----------------------------")

	jsn, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", route.String(), bytes.NewBuffer(jsn))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("apiKey %s", client.CachedToken))

	fmt.Println("-----------------------------")
	fmt.Println("req")
	fmt.Printf("--> %s\n\n", formatRequest(req))
	fmt.Println("-----------------------------")

	res, err := client.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("--> %s\n\n", formatRequest(req))
	fmt.Println("-----------------------------")
	return handleAPIResponse(res)
}

// Delete TODO Doc Comment
func (client *Client) Delete(path string) (*GenericAPIResponse, error) {

	if client.CachedToken == "" {
		return nil, errors.New("coding error - method called without setting API token")
	}

	route := client.BaseURL
	route.Path = path

	req, err := http.NewRequest("DELETE", route.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("apiKey %s", client.CachedToken))

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return handleAPIResponse(res)

}

// IsValidSematextRegion checks sematext api region is valid.
func IsValidSematextRegion(region string) bool {
	switch region {
	case "EU", "US":
		return true
	}
	return false
}

// IsValidUUID checks a string is UUIDv4
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// formatRequest is a debug routine for inspection of an http.Request.
func formatRequest(r *http.Request) string {
	var result []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	result = append(result, fmt.Sprintf("Url: %v", url))
	result = append(result, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			result = append(result, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		result = append(result, "\n")
		result = append(result, r.Form.Encode())
	}
	return strings.Join(result, "\n")
}

// formatResponse is a debug routine for inspection of an http.Response.
func formatResponse(r *http.Response) string {
	var result []string
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			result = append(result, fmt.Sprintf("%v: %v", name, h))
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	result = append(result, string(body))

	return strings.Join(result, "\n")
}

// handleAPIResponse parses responses comsing back from the API.
func handleAPIResponse(response *http.Response) (*GenericAPIResponse, error) {

	var err error
	genericAppResponse := &GenericAPIResponse{}

	err = json.NewDecoder(response.Body).Decode(genericAppResponse)

	switch {
	case err == io.EOF:
		return nil, errors.New("Response from API is unexpectedy empty. ")
	case response.ContentLength == 0:
		return nil, errors.New("ERROR: Response from API is unexpectedy empty. ")
	case err != nil:
		return nil, err
	}

	fmt.Printf("%+v\n", response.StatusCode)

	switch response.StatusCode {
	case 200:
		return genericAppResponse, nil
	case 400:
		fmt.Printf("%+v\n", genericAppResponse)
		return nil, errors.New(genericAppResponse.Message)
	case 401:
		fmt.Printf("%+v\n", genericAppResponse)
		return nil, errors.New(genericAppResponse.Message)
	case 402:
		fmt.Printf("%+v\n", genericAppResponse)
		return nil, errors.New(genericAppResponse.Message)
	case 403:
		fmt.Printf("%+v\n", genericAppResponse)
		return nil, errors.New(genericAppResponse.Message)
	case 404:
		fmt.Printf("%+v\n", genericAppResponse)
		return nil, errors.New(genericAppResponse.Message)
	case 500:
		fmt.Printf("%+v\n", genericAppResponse)
		return nil, errors.New(genericAppResponse.Message)
	default:
		return nil, fmt.Errorf("Unexected status (%d) return from Sematext API", response.StatusCode)
	}
}
