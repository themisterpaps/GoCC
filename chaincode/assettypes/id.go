package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var ID = assets.AssetType{
	Tag:         "id",
	Label:       "Identificacao",
	Description: "Identificacao",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nr",
			Label:    "Numero do Documento",
			DataType: "string",
		},

		{
			Tag:      "tipo",
			Label:    "Tipo de Documento",
			DataType: "string",
		},
		{

			Tag:      "emissao",
			Label:    "Data de emissao",
			DataType: "datetime",
		},
		{

			Tag:      "validade",
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
			Label:    "url do documento",
			DataType: "string",
		},
	},
}
