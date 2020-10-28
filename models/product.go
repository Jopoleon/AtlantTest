package models

import (
	"github.com/Jopoleon/AtlantTest/server/proto_server"
)

type Product struct {
	Name         string  `bson:"name" json:"name,omitempty"`
	LastPrice    float32 `bson:"last_price" json:"last_price,omitempty"`
	PriceChanges int32   `bson:"price_changes" json:"price_changes,omitempty"`
	LastUpdated  string  `bson:"last_updated" json:"last_updated,omitempty"`
}

func (p *Product) ToGRPCProduct() *proto_server.Product {
	return &proto_server.Product{
		Name:         p.Name,
		LastPrice:    p.LastPrice,
		PriceChanges: p.PriceChanges,
		LastUpdated:  p.LastUpdated,
	}
}

type Products []Product

func (ps Products) ToGRPCProducts() []*proto_server.Product {
	var res []*proto_server.Product
	for _, p := range ps {
		res = append(res, p.ToGRPCProduct())
	}
	return res
}

//
//func FromGRPCProduct(p *proto_server.Product) *Product {
//	return &Product{
//		ID:           p.Id,
//		Name:         p.Name,
//		LastPrice:    p.LastPrice,
//		PriceChanges: p.PriceChanges,
//		LastUpdated:  p.LastUpdated,
//	}
//}
