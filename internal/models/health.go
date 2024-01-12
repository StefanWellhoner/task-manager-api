package model

type Health struct {
	Health   string `json:"health"`
	Database string `json:"database"`
}

const (
	StatusGreen  = "green"
	StatusOrange = "orange"
	StatusRed    = "red"
)
