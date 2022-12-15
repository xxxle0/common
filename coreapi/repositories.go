package coreapi

import (
	"bytes"
	"common/coreapi/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepositoryAPI struct {
	baseURL string
	v2      string
}

type IRepositoryAPI interface{}

func NewRepositoryAPI(baseURL string) IRepositoryAPI {
	v2 := fmt.Sprintf("%s/v2/repositories", baseURL)
	return RepositoryAPI{
		baseURL: baseURL,
		v2:      v2,
	}
}

func (c RepositoryAPI) GetRepositoryById(repositoryId int) (*model.Repository, error) {
	var repository model.Repository
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
