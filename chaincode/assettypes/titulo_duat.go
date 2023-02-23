package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Titulo_duat = assets.AssetType{
	Tag:         "titulo_duat",
	Label:       "Titulo de DUAT",
	Description: "Titulo de DUAT",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "nr_titulo",
			Label:    "Numero do Titulo",
			DataType: "string",
		},

		{
			Tag:      "finalidade",
			Label:    "Finalidade de Uso",
			DataType: "string",
		},
		{
			Tag:      "prazo",
			Label:    "prazo",
			DataType: "number",
		},
		{
			Tag:      "url",
			Label:    "url do titulo",
			DataType: "datetime",
		},
		{
			Tag:      "data_emissao",
			Label:    "data de emissao",
			DataType: "datetime",
		},

	},
}
