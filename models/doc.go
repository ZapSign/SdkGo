package models

type Doc struct {
	Sandbox             bool     `json:"sandbox"`
	Name                string   `json:"name"`
	Url_pdf             string   `json:"url_pdf"`
	Brand_primary_color string   `json:"brand_primary_color"`
	Lang                string   `json:"lang"`
	Signers             []Signer `json:"signers"`
}

func (doc Doc) CreateDocWithUrlPdf(signers []Signer) Doc {
	return Doc{
		Sandbox:             true,
		Name:                "Golang Example",
		Url_pdf:             "https://zapsign.s3.amazonaws.com/2022/1/pdf/63d19807-cbfa-4b51-8571-215ad0f4eb98/ca42e7be-c932-482c-b70b-92ad7aea04be.pdf",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Signers:             signers,
	}
}

func (doc Doc) CreateDocWithDocPdf(signers []Signer) Doc {
	return Doc{
		Sandbox:             true,
		Name:                "Golang Example",
		Url_pdf:             "https://zapsign.s3.amazonaws.com/2022/1/pdf/63d19807-cbfa-4b51-8571-215ad0f4eb98/ca42e7be-c932-482c-b70b-92ad7aea04be.pdf",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Signers:             signers,
	}
}
