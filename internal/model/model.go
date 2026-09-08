package model

type Dataset struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Organization string `json:"organization"`
	Category     string `json:"category"`
	Format       string `json:"format"`
	Updated      string `json:"updated"`
	Records      int    `json:"records"`
}
