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
			Tag:      "nr_despacho",
			Label:    "Nr despacho",
			DataType: "string",
		},

		//Verificação do técnico da SG
		{
			Tag:      "nome_tec_sg",
			Label:    "Nome do Tecnico da SG",
			DataType: "string",
		},
		{
			Tag:      "aprovacao_sg",
			Label:    "Resposta de Aprovacao da SG",
			DataType: "string",
		},
		//Verificação do técnico da R.ACL
		{
			Tag:      "nome_tec_racl",
			Label:    "Nome do Tecnico da RACL",
			DataType: "string",
		},
		{
			Tag:      "aprovacao_racl",
			Label:    "Resposta de Aprovacao da RACL",
			DataType: "string",
		},
		//Verificação do técnico da R.DUAT
		{
			Tag:      "nome_tec_rduat",
			Label:    "Nome do Tecnico da RDUAT",
			DataType: "string",
		},
		{
			Tag:      "aprovacao_rduat",
			Label:    "Resposta de Aprovacao da RDUAT",
			DataType: "string",
		},
			//parecer do director 
		{
			Tag:      "nome_diretor",
			Label:    "Nome do Director",
			DataType: "string",
		},
		{
			Tag:      "descricao",
			Label:    "Descricao do parecer",
			DataType: "string",
		},
		 
		{
			Tag:      "url_despacho",
			Label:    "URL Despacho",
			DataType: "string",
		},
		{
			Tag:      "url_parecer",
			Label:    "URL Despacho",
			DataType: "string",
		},
		 
		 
		
	},
}
