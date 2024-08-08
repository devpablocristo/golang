package shared

type ItemReq struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}
