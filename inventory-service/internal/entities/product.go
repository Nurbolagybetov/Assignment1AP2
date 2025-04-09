package entities

type Product struct {
	ID       string  `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Stock    int     `json:"stock" bson:"stock"`
	Price    float64 `json:"price" bson:"price"`
}
