package models

type DocFromTemplate struct {
	Sandbox                   bool             `json:"sandbox"`
	Name                      string           `json:"name"`
	Lang                      string           `json:"lang"`
	Disable_signer_emails     bool             `json:"disable_signer_emails"`
	Signed_file_only_finished bool             `json:"signed_file_only_finished"`
	Brand_logo                string           `json:"brand_logo"`
	Brand_primary_color       string           `json:"brand_primary_color"`
	Brand_name                string           `json:"brand_name"`
	External_id               string           `json:"external_id"`
	Folder_path               string           `json:"folder_path"`
	Signature_order_active    bool             `json:"signature_order_active"`
	Observers                 []string         `json:"observers"`
	Reminder_every_n_days     int              `json:"reminder_every_n_days"`
	Signer_name               string           `json:"signer_name"`
	Template_id               string           `json:"template_id"`
	Deparafromtemplate        []DeParaTemplate `json:"data"`
}

func (docFromTemplate DocFromTemplate) CreateDocFromTemplate(deParaTemplateMock []DeParaTemplate) DocFromTemplate {
	return DocFromTemplate{
		Sandbox:             true,
		Template_id:         "aea8a7f9-5591-4a3e-b243-38926c2b2e00",
		Signer_name:         "João dos Santos",
		External_id:         "123",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Deparafromtemplate:  deParaTemplateMock,
	}

}
