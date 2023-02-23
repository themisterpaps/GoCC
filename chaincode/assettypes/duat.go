package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Duat = assets.AssetType{
	Tag:         "duat",
	Label:       "Duat",
	Description: "Duat",

	Props: []assets.AssetProp{
		{
			//Key
			Required: true,
			IsKey:    true,
			Tag:      "nr_processo",
			Label:    "Nr processo",
			DataType: "string",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "nr_nuit",
			Label:    "NUIT",
			DataType: "->municipe",
		},
		{
			/// Reference to another asset
			Tag:      "nr_parcela",
			Label:    "Informação da Parcela de Terra",
			DataType: "->parcela",
		},
		{
			Tag:      "parecer",
			Label:    "Parecer Tecnico",
			DataType: "->parecer",
		},
		{
			Tag:      "duat_provisorio",
			Label:    "DUAT provisório",
			DataType: "->duat_provisorio",
		},
		{
			
			Tag:      "licenca",
			Label:    "Licenca de Construção",
			DataType: "->licenca",
		},
		{
			
			Tag:      "titulo",
			Label:    "Titulo de DUAT",
			DataType: "->titulo_duat",
		},
		{
			
			Tag:      "registro",
			Label:    "registro de propriedade",
			DataType: "->registro",
		},
	},
}
