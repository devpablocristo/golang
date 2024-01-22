package main

import (
	"encoding/json"
	"fmt"

	"testdto/dto"
)

func main() {

	jsonData := []byte(`{
		"id": "layout123",
		"status": "active",
		"variants": [
			{
				"id": "variant1",
				"units": ["unit1", "unit2"],
				"support": [
					{
						"sites": ["site1", "site2"],
						"clients": [
							{"id": 1},
							{"id": 2}
						]
					}
				],
				"status": "enabled"
			}
		]
	}`)

	// Decodificar los datos JSON en la estructura LayoutsResponse
	var layoutResponse dto.LayoutsResponse
	err := json.Unmarshal(jsonData, &layoutResponse)
	if err != nil {
		fmt.Printf("Error al decodificar el JSON: %v\n", err)
		return
	}

	// Traducir la estructura LayoutsResponse a la estructura entity.LayoutsResponse
	entityLayoutsResponse, err := layoutResponse.TranslateLayouts2Domain()
	if err != nil {
		fmt.Printf("Error al traducir la estructura LayoutsResponse: %v\n", err)
		return
	}

	// Imprimir la estructura entity.LayoutsResponse en formato JSON
	entityLayoutsResponseJSON, err := json.MarshalIndent(entityLayoutsResponse, "", "  ")
	if err != nil {
		fmt.Printf("Error al codificar la estructura entity.LayoutsResponse en JSON: %v\n", err)
		return
	}

	fmt.Println("entity.LayoutsResponse JSON:")
	fmt.Println(string(entityLayoutsResponseJSON))
}
