package apicall

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
)

type HeaderDetails struct {
	Key, Value string
}

// APICall makes an API call with the provided parameters
func APICallWithCookie(url, methodType, jsonData string, headers []HeaderDetails, req *http.Request) (string, *http.Response, error) {
	var response *http.Response
	var body []byte
	// Create a request with the specified method, URL, and JSON data
	reqBody := bytes.NewBuffer([]byte(jsonData))
	request, err := http.NewRequest(methodType, url, reqBody)
	if err != nil {
		return "", response, err
	}

	// Set additional headers
	for _, header := range headers {
		request.Header.Set(header.Key, header.Value)
	}

	// Copy cookies from the provided request to the new request
	if req != nil {
		request.Header.Set("Cookie", req.Header.Get("Cookie"))
	}

	// Perform the API call
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return "", response, err
	} else {
		defer response.Body.Close()
		//log.Println("apiUtil.Header", response.Header)

		body, err = io.ReadAll(response.Body)
		if err != nil {

			return "", response, err
		}
	}

	return string(body), response, nil
}

// APICall makes an HTTP request with custom headers.
// It takes a URL, JSON payload, HTTP method, an array of custom headers, and a source parameter.
// It returns the response body as a string, the http.Response object, and an error (if any).
func APICall(pURl string, pJson string, pMethod string, pHeaders []HeaderDetails, pSource string) (string, *http.Response, error) {
	log.Println("APICall (+)  ")
	//create local variable to store data
	var lRespJson string
	var lResponse *http.Response

	// Create a payload from the JSON string
	lPayload := strings.NewReader(string(pJson))

	// Create an HTTP client
	lClient := &http.Client{}

	// Create an HTTP request with the specified method, URL, and payload
	lRequest, lErr := http.NewRequest(pMethod, pURl, lPayload)
	if lErr != nil {
		log.Println("APICall:001", lErr.Error())
		return lRespJson, lResponse, lErr
	}

	// Add a default Content-Type header
	lRequest.Header.Add("Content-Type", "application/json")

	// Add custom headers from the slice
	for _, lHeader := range pHeaders {
		lRequest.Header.Add(lHeader.Key, lHeader.Value)
	}

	// Perform the HTTP request
	lResponse, lErr = lClient.Do(lRequest)
	if lErr != nil {
		log.Println("APICall:002", lErr.Error())
		return lRespJson, lResponse, lErr
	}
	defer lResponse.Body.Close()

	// Read the response body
	body, lErr := io.ReadAll(lResponse.Body)
	if lErr != nil {
		log.Println("APICall:003", lErr.Error())
		return lRespJson, lResponse, lErr
	} else {
		// Convert the response body to a string
		lRespJson = string(body)
	}
	log.Println("APICall (-)  ")
	return lRespJson, lResponse, nil
}
