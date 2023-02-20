package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Parecer = assets.AssetType{
	Tag:         "parecer",
	Label:       "Parecer",
	Description: "Parecer",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nr_processo",
			Label:    "Nr processo",
			DataType: "string",
		},
		{
			Tag:      "parecer_sg",
			Label:    "Parecer da SG",
			DataType: "->verificacao",
		},
		{
			// String list
			Tag:      "parecer_rduat",
			Label:    "Parecer do RDUAT",
			DataType: "->verificacao",
		},
		{
			// String list
			Tag:      "parecer_ral",
			Label:    "Parecer do RAL",
			DataType: "->verificacao",
		},
		{
			Tag:      "dir_documento",
			Label:    "url do documento",
			DataType: "string",
		},
		{
			Tag:      "despacho",
			Label:    "url do despacho",
			DataType: "string",
		},
	},
}
