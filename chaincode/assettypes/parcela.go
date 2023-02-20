package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Parcela = assets.AssetType{
	Tag:         "parcela",
	Label:       "Parcela",
	Description: "Parcela",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nr_parcela",
			Label:    "Nr Parcela",
			DataType: "string",
		},

		{
			/// Reference to another asset
			Tag:      "ano_ocupacao",
			Label:    "Ano de Ocupacao",
			DataType: "datetime",
		},
		{
			// String list
			Tag:      "endereco",
			Label:    "endereco",
			DataType: "->endereco",
		},
	},
}
