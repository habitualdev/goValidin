package goValidin

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Validin struct {
	ApiKey string
}

func NewValidin(apiKey string) *Validin {
	return &Validin{
		ApiKey: apiKey,
	}
}

func (v *Validin) DoQuery(baseUrl string, parameters map[string]string) ([]byte, error) {
	queryString := ""
	if len(parameters) != 0 {
		for key, value := range parameters {
			if value == "" {
				continue
			}
			queryString += key + "=" + value + "&"
		}
		queryString = queryString[:len(queryString)-1]
		queryString = "?" + queryString
	}
	req, err := http.NewRequest("GET", baseUrl+queryString, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.URL.String())
	req.Header.Add("Authorization", "BEARER "+v.ApiKey)
	req.Header.Add("content-type", "application/json")
	//req.Header.Add("User-Agent", "GoValidin/1.0.0")
	//client := &http.Client{}
	fmt.Println(req.Header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}
	return respBytes, nil
}
