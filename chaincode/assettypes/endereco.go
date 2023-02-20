package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Endereco = assets.AssetType{
	Tag:         "endereco",
	Label:       "Endereco",
	Description: "Enderecos",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
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
	},
}
