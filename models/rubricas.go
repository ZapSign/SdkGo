package models

type Rubricas struct {
	Page                     int64   `json:"page"`
	Relative_position_bottom float32 `json:"relative_position_bottom"`
	Relative_position_left   float32 `json:"relative_position_left"`
	Relative_size_x          float32 `json:"relative_size_x"`
	Relative_size_y          float32 `json:"relative_size_y"`
	Type                     string  `json:"type"`
	Signer_token             string  `json:"signer_token"`
}