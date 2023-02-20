package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Verificacao = assets.AssetType{
	Tag:         "verificacao",
	Label:       "Verificacao",
	Description: "Verificacao do Tecnico",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nome",
			Label:    "Nome do Verificador",
			DataType: "string",
		},
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "data",
			Label:    "Data da Verificacao",
			DataType: "datetime",
		},
		{
			/// Reference to another asset
			Tag:      "aprovacao",
			Label:    "resposta de aprovacao",
			DataType: "string",
		},
		{
			// String list
			Tag:      "cargo",
			Label:    "Cargo do Profissional",
			DataType: "string",
		},
	},
}
