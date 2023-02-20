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
			Tag:      "nuit",
			Label:    "NUIT",
			DataType: "number",
		},

		{

			Tag:      "nome",
			Label:    "Nome Completo",
			DataType: "string",
		},
		{

			Tag:      "dataDeNascimento",
			Label:    "Data de Nascimento",
			DataType: "datetime",
		},
		{
			Tag:      "estadoCivil",
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
			Tag:      "nuic",
			Label:    "NUIC",
			DataType: "string",
		},
		{
			// String list
			Tag:      "declaracao",
			Label:    "Declaracao do bairro",
			DataType: "string",
		},
		{
			// Objecto ID
			Tag:      "id",
			Label:    "Identificacao",
			DataType: "->id",
		},
		{
			// Objecto endereco
			Tag:      "endereco",
			Label:    "Endereco",
			DataType: "[]->endereco",
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
			DataType: "number",
		},
	},
}
