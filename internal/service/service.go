package service

// Service holds general information about the service
type Service struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Description string `json:"description"`
	NumVersions int    `json:"numVersions"`
}
