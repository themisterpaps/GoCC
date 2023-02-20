package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Licensa = assets.AssetType{
	Tag:         "licensa",
	Label:       "Licensa",
	Description: "Licensa",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "nr_processo",
			Label:    "Nr do Processo",
			DataType: "string",
		},
		{
			Tag:      "planta",
			Label:    "Planta Topografica",
			DataType: "string",
		},
		{
			Tag:      "area",
			Label:    "Area ",
			DataType: "number",
		},
		{
			Tag:      "memoria",
			Label:    "Memoria ",
			DataType: "string",
		},
		{
			Tag:      "certidao",
			Label:    "Certidao",
			DataType: "string",
		},
		{
			Tag:      "url",
			Label:    "url da Licensa",
			DataType: "string",
		},
		{
			Tag:      "data",
			Label:    "data de emissao",
			DataType: "datetime",
		},
		{
			Tag:      "prazo",
			Label:    "prazo de validade",
			DataType: "datetime",
		},
	},
}
