package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
var Duat = assets.AssetType{
	Tag:         "duat",
	Label:       "Duat",
	Description: "Duat",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nr_processo",
			Label:    "Nr processo",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "nuit",
			Label:    "NUIT",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			/// Reference to another asset
			Tag:      "info_parcela",
			Label:    "Informação da Terra",
			DataType: "->parcela",
		},
		{
			// String list
			Tag:      "parecer",
			Label:    "Parecer Tecnico",
			DataType: "->parecer",
		},
		{
			// String list
			Tag:      "licensa",
			Label:    "Licensa de Construção",
			DataType: "->licensa",
		},
		{
			// String list
			Tag:      "registro",
			Label:    "registro de propriedade",
			DataType: "->registro",
		},
		{
			// String list
			Tag:      "titulo",
			Label:    "Titulo de DUAT",
			DataType: "->titulo",
		},
	},
}
