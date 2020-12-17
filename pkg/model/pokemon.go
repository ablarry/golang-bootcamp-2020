package model

// Pokemon entity
type Pokemon struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Weakness string `json:"wakness"`
}
