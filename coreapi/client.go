package coreapi

type CoreAPIClient struct {
	baseURL      string
	scans        IScanAPI
	repositories IRepositoryAPI
}

func NewCoreAPIClient[R, A any](baseURL string) CoreAPIClient {
	scanClient := NewScanAPI[R](baseURL)
	repositoryClient := NewRepositoryAPI[A](baseURL)
	return CoreAPIClient[R, A]{
		baseURL:      baseURL,
		scans:        scanClient,
		repositories: repositoryClient,
	}
}

func (c CoreAPIClient[R, _]) Scans() IScanAPI[R] {
	return c.scans
}

func (c CoreAPIClient[_, A]) Repositories() IRepositoryAPI[A] {
	return c.repositories
}
