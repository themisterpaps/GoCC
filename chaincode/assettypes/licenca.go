package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Licenca = assets.AssetType{
	Tag:         "licenca",
	Label:       "Licenca de Construção",
	Description: "Licenca de Construção",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "nr_licenca",
			Label:    "Nr da licença",
			DataType: "string",
		},
		{
			Tag:      "url_planta",
			Label:    "Planta Topografica",
			DataType: "string",
		},
		{
			Tag:      "memoria",
			Label:    "Memoria descritiva",
			DataType: "string",
		},
		{
			Tag:      "certidao",
			Label:    "Certidao Benfeitorias",
			DataType: "string",
		},
		{
			Tag:      "url_licenca",
			Label:    "url da Licensa",
			DataType: "string",
		},
		{
			Tag:      "data_emissao",
			Label:    "data de emissao",
			DataType: "datetime",
		},
		{
			Tag:      "data_validade",
			Label:    "data de validade",
			DataType: "datetime",
		},
	},
}
