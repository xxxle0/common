package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ScanAPI[R any] struct {
	baseURL string
	v2      string
}

type IScanAPI[R any] interface {
	GetScansOfRepository(repositoryId int) ([]R, error)
	GetScans(repositoryId int, createdAt string) ([]R, error)
	GetScanByIds(scanIds []string) ([]R, error)
	InsertScans(scans []R) ([]R, error)
	GetScanById(scanId string) (*R, error)
}

func NewScanAPI[R any](baseURL string) IScanAPI[R] {
	v2 := fmt.Sprintf("%s/v2/scans", baseURL)
	return ScanAPI[R]{
		baseURL: baseURL,
		v2:      v2,
	}
}

func (s ScanAPI[R]) GetScansOfRepository(repositoryId int) ([]R, error) {
	var scans []R
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

func (s ScanAPI[R]) GetScans(repositoryId int, createdAt string) ([]R, error) {
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
	var scans []R
	err = json.NewDecoder(res.Body).Decode(&scans)
	if err != nil {
		return nil, err
	}
	return scans, nil
}

func (s ScanAPI[R]) InsertScans(scans []R) ([]R, error) {
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
	var insertedScans []R
	err = json.NewDecoder(res.Body).Decode(&insertedScans)
	if err != nil {
		return nil, err
	}
	return insertedScans, nil
}

func (s ScanAPI[R]) GetScanById(scanId string) (*R, error) {
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
	var scan R
	err = json.NewDecoder(res.Body).Decode(&scan)
	if err != nil {
		return nil, err
	}
	return &scan, nil
}

func (s ScanAPI[R]) GetScanByIds(scanIds []string) ([]R, error) {
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
	var scans []R
	err = json.NewDecoder(res.Body).Decode(&scans)
	if err != nil {
		return nil, err
	}
	return scans, nil
}
