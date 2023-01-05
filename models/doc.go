package models

type Doc struct {
	Sandbox             bool     `json:"sandbox"`
	Name                string   `json:"name"`
	Url_pdf             string   `json:"url_pdf"`
	Brand_primary_color string   `json:"brand_primary_color"`
	Lang                string   `json:"lang"`
	Signers             []Signer `json:"signers"`
}
