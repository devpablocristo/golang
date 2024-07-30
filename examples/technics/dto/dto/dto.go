package dto

import (
	entity "testdto/entities"
)

// LayoutsResponse from the service
type LayoutsResponse struct {
	ID       string    `json:"id"`
	Status   string    `json:"status"`
	Variants []Variant `json:"variants"`
}

// Variant nested responses
type Variant struct {
	ID      string    `json:"id"`
	Units   []string  `json:"units"`
	Support []Support `json:"support"`
	Status  string    `json:"status"`
}

type Support struct {
	Sites   []string `json:"sites"`
	Clients []Client `json:"clients"`
}

type Client struct {
	ID uint `json:"id"`
}

func (layout *LayoutsResponse) TranslateLayouts2Domain() (entity.LayoutsResponse, error) {
	var lay entity.LayoutsResponse

	// layout.Variants => []Variant{}
	resVariants := make([]entity.Variant, len(layout.Variants))

	for i, v := range layout.Variants {
		var resSupport []entity.Support

		for _, s := range v.Support {
			var resClients []entity.Client
			for _, c := range s.Clients {
				resClients = append(resClients, entity.Client{ID: c.ID})
			}

			resSupport = append(resSupport, entity.Support{
				Sites:   s.Sites,
				Clients: resClients,
			})
		}

		resVariants[i] = entity.Variant{
			ID:      v.ID,
			Units:   v.Units,
			Support: resSupport,
			Status:  v.Status,
		}
	}

	lay.ID = layout.ID
	lay.Status = layout.Status
	lay.Variants = resVariants

	return lay, nil
}
