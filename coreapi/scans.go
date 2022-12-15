package coreapi

import (
	"bytes"
	"common/coreapi/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type ScanAPI struct {
	baseURL string
	v2      string
}

type IScanAPI interface{}

func NewScanAPI(baseURL string) IScanAPI {
	v2 := fmt.Sprintf("%s/v2/scans", baseURL)
	return ScanAPI{
		baseURL: baseURL,
		v2:      v2,
	}
}

func (s ScanAPI) GetScansOfRepository(repositoryId int) ([]model.Scan, error) {
	var scans []model.Scan
	url := fmt.Sprintf("%s/getScansOfRepository", s.v2)
	body := map[string]int{
		"repositoryId": repositoryId,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&scans)
	if err != nil {
		return nil, err
	}
	return scans, nil
}

func (s ScanAPI) GetScans(repositoryId int, createdAt string) ([]model.Scan, error) {
	url := fmt.Sprintf("%s/getScans", s.v2)
	requestBody := map[string]interface{}{
		"repositoryId": repositoryId,
		"createdAt":    createdAt,
	}
	b, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var scans []model.Scan
	err = json.NewDecoder(res.Body).Decode(&scans)
	if err != nil {
		return nil, err
	}
	return scans, nil
}

func (s ScanAPI) InsertScans(scans []model.Scan) ([]model.Scan, error) {
	url := fmt.Sprintf("%s/insertScans", s.v2)
	requestBody := map[string]interface{}{
		"scans": scans,
	}
	b, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var insertedScans []model.Scan
	err = json.NewDecoder(res.Body).Decode(&insertedScans)
	if err != nil {
		return nil, err
	}
	return insertedScans, nil
}

func (s ScanAPI) GetScanById(scanId string) (*model.Scan, error) {
	url := fmt.Sprintf("%s/getScanById", s.v2)
	requestBody := map[string]interface{}{
		"scanId": scanId,
	}
	b, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var scan model.Scan
	err = json.NewDecoder(res.Body).Decode(&scan)
	if err != nil {
		return nil, err
	}
	return &scan, nil
}

func (s ScanAPI) GetScanByIds(scanIds []string) ([]model.Scan, error) {
	url := fmt.Sprintf("%s/getScanById", s.v2)
	requestBody := map[string]interface{}{
		"scanIds": scanIds,
	}
	b, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var scans []model.Scan
	err = json.NewDecoder(res.Body).Decode(&scans)
	if err != nil {
		return nil, err
	}
	return scans, nil
}
