package model

import "time"

type ProductInfo struct {
	Name	string `json:"weburl,omitempty"`
	Image string `json:"image,omitempty"`
	Description string `json:"description,omitempty"`
	Price string `json:"price,omitempty"`
	TotalReview string `json:"totalreview,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

