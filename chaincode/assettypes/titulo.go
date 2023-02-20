package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Titulo = assets.AssetType{
	Tag:         "titulo",
	Label:       "Titulo",
	Description: "Titulo",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nr_processo",
			Label:    "Numero do processo",
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
	},
}
