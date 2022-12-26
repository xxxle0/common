package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepositoryAPI[R any] struct {
	baseURL string
	v2      string
}

type IRepositoryAPI[R any] interface {
	GetRepositoryById(repositoryId int) (*R, error)
}

func NewRepositoryAPI[R any](baseURL string) IRepositoryAPI[R] {
	v2 := fmt.Sprintf("%s/v2/repositories", baseURL)
	return RepositoryAPI[R]{
		baseURL: baseURL,
		v2:      v2,
	}
}

func (c RepositoryAPI[R]) GetRepositoryById(repositoryId int) (*R, error) {
	var repository R
	url := fmt.Sprintf("%s/getRepositoryById", c.v2)
	reqBody := map[string]interface{}{
		"repositoryId": repositoryId,
	}
	json_data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return nil, err
	}
	return &repository, nil
}
