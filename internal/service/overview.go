package service

// Overview holds general information about the service
type Overview struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Description string `json:"description"`
	NumVersions int    `json:"numVersions"`
}
