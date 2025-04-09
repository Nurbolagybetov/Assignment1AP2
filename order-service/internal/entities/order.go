package entities

type Order struct {
	ID         string   `json:"id" bson:"_id"`
	UserID     string   `json:"user_id" bson:"user_id"`
	ProductIDs []string `json:"product_ids" bson:"product_ids"`
	Total      float64  `json:"total" bson:"total"`
	Status     string   `json:"status" bson:"status"`
}
