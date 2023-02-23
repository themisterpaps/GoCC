package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Duat_provisorio = assets.AssetType{
	Tag:         "duat_provisorio",
	Label:       "Duat Provisorio",
	Description: "Duat Provisorio",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nr_autorizacao",
			Label:    "Numero da Autorização",
			DataType: "string",
		},
		{
			Tag:      "codigo_cadastral",
			Label:    "Codico Cadastral",
			DataType: "string",
		},
		{
			Tag:      "local_emissao",
			Label:    "Local de Emissão",
			DataType: "string",
		},
		{
			Tag:      "tipo_exploracao",
			Label:    "Tipo de Exploracao",
			DataType: "string",
		},
		{
			Tag:      "url",
			Label:    "url do titulo",
			DataType: "string",
		},
		{
			Tag:      "data_emissao",
			Label:    "Data de emissão",
			DataType: "datetime",
		},
	},
}
