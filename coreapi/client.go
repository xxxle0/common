package coreapi

type CoreAPIClient struct {
	baseURL string
	Scans   IScanAPI
}

func NewCoreAPIClient(baseURL string) CoreAPIClient {
	scanClient := NewScanAPI(baseURL)
	return CoreAPIClient{
		baseURL: baseURL,
		Scans:   scanClient,
	}
}
