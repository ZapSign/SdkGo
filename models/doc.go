package models

type Doc struct {
	Sandbox             bool     `json:"sandbox"`
	Name                string   `json:"name"`
	Url_pdf             string   `json:"url_pdf"`
	Brand_primary_color string   `json:"brand_primary_color"`
	Lang                string   `json:"lang"`
	Signers             []Signer `json:"signers"`
}

func (doc Doc) CreateDoc(signers []Signer) Doc {
	return Doc{
		Sandbox:             true,
		Name:                "Golang Example",
		Url_pdf:             "https://zapsign.s3.amazonaws.com/2022/1/docs/d7660fd2-fe74-4691-bec8-5c42c0ae2b3f/39a35070-8987-476d-86e3-75d91f588a5a.docx",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Signers:             signers,
	}
}
