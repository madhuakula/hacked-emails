package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	hackedEmailsAPIURI = "https://hacked-emails.com/api"
)

type Response struct {
	Data    []Data `json:"data"`
	Query   string `json:"query"`
	Status  string `json:"status"`
	Results int64  `json:"results"`
}

type Data struct {
	Source_url      string `json:"source_url"`
	Source_lines    int64  `json:"source_lines"`
	Source_size     int64  `json:"source_size"`
	Source_network  string `json:"source_network"`
	Source_provider string `json:"source_provider"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Date_created    string `json:"date_created"`
	Date_leaked     string `json:"date_leaked"`
	Emails_count    int64  `json:"emails_count"`
	Verified        bool   `json:"verified"`
	Details         string `json:"details"`
}

// Check returns the details from hacked-emails for a given email.
func Check(email string) (response *Response, err error) {
	endpoint := fmt.Sprintf("%s?q=%s", hackedEmailsAPIURI, url.QueryEscape(email))
	resp, err := http.Get(endpoint)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
