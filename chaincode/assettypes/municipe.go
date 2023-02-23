package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Municipe = assets.AssetType{
	Tag:         "municipe",
	Label:       "Municipe",
	Description: "Municipe",

	Props: []assets.AssetProp{
		{
			// Single Key
			Required: true,
			IsKey:    true,
			Tag:      "nr_nuit",
			Label:    "NUIT",
			DataType: "number",
		},
		{
			Tag:      "nome",
			Label:    "Nome Completo",
			DataType: "string",
		},
		{

			Tag:      "data_nascimento",
			Label:    "Data de Nascimento",
			DataType: "datetime",
		},
		{
			Tag:      "estado_civil",
			Label:    "Estado Civil",
			DataType: "string",
		},
		{
			Tag:      "genero",
			Label:    "Genero",
			DataType: "string",
		},
		{
			Tag:      "naturalidade",
			Label:    "Naturalidade",
			DataType: "string",
		},
		{
			Tag:      "url_declaracao",
			Label:    "URL Declaracao do bairro",
			DataType: "string",
		},
		//IDENTIFICACAO
		{
			Tag:      "nr_BI",
			Label:    "Numero do BI",
			DataType: "string",
		},
		{
			Tag:      "tipo_documentos",
			Label:    "Tipo de Documento",
			DataType: "string",
		},
		{

			Tag:      "data_emissao",
			Label:    "Data de emissao",
			DataType: "datetime",
		},
		{

			Tag:      "data_validade",
			Label:    "Data de validade",
			DataType: "datetime",
		},
		{
			Tag:      "local",
			Label:    "Local de Emissao",
			DataType: "string",
		},
		{
			Tag:      "url",
			Label:    "url do BI",
			DataType: "string",
		},

		//Endereço
		{

			Tag:      "bairro",
			Label:    "Bairro",
			DataType: "string",
		},
		{
			Tag:      "provincia",
			Label:    "provincia",
			DataType: "string",
		},
		{

			Tag:      "municipio",
			Label:    "Municipio",
			DataType: "string",
		},
		{
			Tag:      "distrito",
			Label:    "Distrito Municipal",
			DataType: "string",
		},
		{
			Tag:      "posto",
			Label:    "Posto administrativo",
			DataType: "string",
		},
		{
			// String list
			Tag:      "contacto",
			Label:    "Contacto",
			DataType: "[]number",
		},
		{
			// String list
			Tag:      "email",
			Label:    "email",
			DataType: "string",
		},
	},
}
