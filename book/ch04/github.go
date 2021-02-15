package ch04

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Issues struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	Title     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Login string
}

const issueURL = "https://api.github.com/search/issues"

func SearchIssues(terms []string) (*Issues, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Issues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
