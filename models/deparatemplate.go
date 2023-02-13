package models

type DeParaTemplate struct {
	De   string `json:"de"`
	Para string `json:"para"`
}

func (deParaTemplate DeParaTemplate) CreateDeParaTemplate() []DeParaTemplate {
	var fakeDeParaVariables = []DeParaTemplate{
		{
			De:   "{{EMAIL CLIENTE}}",
			Para: "cliente@gmail.com",
		},
		{
			De:   "{{NOME COMPLETO}}",
			Para: "João dos Santos",
		},
		{
			De:   "{{NÚMERO DO CPF}}",
			Para: "123.456.789-10",
		},
	}

	return fakeDeParaVariables
}
