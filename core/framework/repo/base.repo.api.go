package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/syariatifaris/arkeus/core/apicalls"
	"github.com/syariatifaris/arkeus/core/config"
	"github.com/syariatifaris/arkeus/core/log/tokolog"
)

//BaseAPIRepository
type BaseAPIRepository interface {
	//Get the base URL
	GetBaseURL() string
}

//APIRepository base structure
type APIRepository struct {
	APIConfig config.APICallConfig
}

//BuildPostHttpAPI will Create HttpAPI
func (a *APIRepository) BuildPostHttpAPI(url string, request interface{}, headers map[string]string) (apicalls.HttpAPI, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return apicalls.HttpAPI{}, errors.New(fmt.Sprintf("error on parsing request body %s : %s", url, err.Error()))
	}

	return apicalls.HttpAPI{
		URL:     url,
		Body:    requestBody,
		Method:  http.MethodPost,
		Headers: headers,
	}, nil
}

//DoRequestSync is Do Request Synchronously
func (a *APIRepository) DoRequest(httpAPI apicalls.HttpAPI) (*http.Response, error) {
	response, err := apicalls.DoRequest(httpAPI)
	if err != nil {
		tokolog.ERROR.Printf("[BaseAPIRepository][DoRequest] Error while calling request. %s/%s, err: %s\n", httpAPI.Method,
			httpAPI.URL, err.Error())
		return response, err
	}

	if response.StatusCode != http.StatusOK {
		tokolog.ERROR.Printf("[BaseAPIRepository][DoRequest] Http status NOT OK. %s/%s. Status: %d\n", httpAPI.Method,
			httpAPI.URL, response.StatusCode)
		return response, errors.New(fmt.Sprintf("%s Http status not OK %d", httpAPI.URL, response.StatusCode))
	}

	return response, err
}

//GetResponseBody gets the response body from http response
func (a *APIRepository) GetResponseBody(response *http.Response) ([]byte, error) {
	return ioutil.ReadAll(response.Body)
}
