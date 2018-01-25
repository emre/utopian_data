package client

import (
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	BaseUrl string
}

func (c *Client) generateURL(action string) string {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, action)
	log.Println(url)
	return url
}

func (c *Client) GetModerators() ModeratorResponse {

	utopianHttpClient := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	var req *http.Request
	req, err := http.NewRequest("GET", c.generateURL("moderators"), nil)
	if err != nil {
		panic(err)
	}

	response, err := utopianHttpClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	moderatorResponse := ModeratorResponse{}
	err = json.Unmarshal(body, &moderatorResponse)
	if err != nil {
		panic(err)
	}

	return moderatorResponse
}

func (c *Client) GetSponsors() SponsorResponse {

	utopianHttpClient := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	var req *http.Request
	req, err := http.NewRequest("GET", c.generateURL("sponsors"), nil)
	if err != nil {
		panic(err)
	}

	response, err := utopianHttpClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	sponsorResponse := SponsorResponse{}
	err = json.Unmarshal(body, &sponsorResponse)
	if err != nil {
		panic(err)
	}

	return sponsorResponse
}

func (c *Client) GetPosts(limit *int, skip *int, postList []Post, hidden bool) []Post {

	defaultLimit := 1000
	defaultSkip := 0

	utopianHttpClient := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	if skip == nil {
		skip = &defaultSkip
	}
	if limit == nil {
		limit = &defaultLimit
	}

	params := fmt.Sprintf("limit=%d&skip=%d", *limit, *skip)
	if hidden {
		params += "&status=flagged"
	}

	var req *http.Request
	req, err := http.NewRequest("GET", c.generateURL(fmt.Sprintf("posts?%s", params)), nil)
	if err != nil {
		panic(err)
	}

	response, err := utopianHttpClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	postResponse := PostResponse{}
	err = json.Unmarshal(body, &postResponse)
	if err != nil {
		panic(err)
	}

	postList = append(postList, postResponse.Posts...)
	if *skip + *limit < postResponse.Total {
		newSkip := *skip + *limit
		return c.GetPosts(limit, &newSkip, postList, hidden)
	}

	return postList
}
