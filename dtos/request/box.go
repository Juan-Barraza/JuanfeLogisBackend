package request

type BoxRequest struct {
	Name       string `json:"name"`
	LocationID uint   `json:"location_id"`
	LabelIDs   []uint `json:"label_ids"`
}
