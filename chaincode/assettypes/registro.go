package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Registro = assets.AssetType{
	Tag:         "registro",
	Label:       "Registro",
	Description: "Registro",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "nr_certidao",
			Label:    "Numero da Certidao de Benfeitorias",
			DataType: "string",
		},
		{
			Tag:      "descricao",
			Label:    "Descricao",
			DataType: "string",
		},
		{
			Tag:      "data_emissao",
			Label:    "Data de Emissao da Certidao",
			DataType: "datetime",
		},
		{
			Tag:      "data_registro",
			Label:    "Data de registro",
			DataType: "datetime",
		},
	},
}
