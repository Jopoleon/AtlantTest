package mongo

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Jopoleon/AtlantTest/models"
	"github.com/Jopoleon/AtlantTest/server/proto_server"
)

const ProductsCollection = "products"

func (m *MongoDB) GetProducts(ctx context.Context, filter *proto_server.ListRequest) ([]models.Product, error) {
	var products []models.Product
	skips := filter.Limit * (filter.Page - 1)
	sortOptions := options.Find().
		SetSort(bson.D{{strings.ToLower(filter.Sort.String()), int(filter.Order.Number())}}).
		SetSkip(int64(skips)).
		SetLimit(int64(filter.Limit))

	cur, err := m.productDB.Collection(ProductsCollection).Find(ctx, bson.D{}, sortOptions)
	if err != nil {
		m.Logger.Error(err)
		return products, err
	}
	for cur.Next(ctx) {
		p := models.Product{}
		err := cur.Decode(&p)
		if err != nil {
			m.Logger.Error(err)
			return products, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (m *MongoDB) SaveProducts(ctx context.Context, ps []models.Product) error {
	col := m.productDB.Collection(ProductsCollection)
	for _, p := range ps {
		p.LastUpdated = time.Now().String()
		filter := bson.D{{"name", p.Name}}
		opts := options.Update().SetUpsert(true)
		cur, err := m.productDB.Collection(ProductsCollection).Find(ctx, filter)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				update := bson.D{
					{"$set", bson.D{
						{"name", p.Name},
						{"last_price", p.LastPrice},
						{"last_updated", p.LastUpdated},
						{"price_changes", 0},
					}},
					{"$inc", bson.D{{"price_changes", 1}}},
				}
				//res, err := col.InsertOne(ctx, p)
				_, err := col.UpdateOne(ctx, filter, update, opts)
				if err != nil {
					m.Logger.Error(err)
					return err
				}
				continue
			}
			m.Logger.Error(err)
			return err
		}
		var oldProduct models.Product
		for cur.Next(ctx) {
			err = cur.Decode(&oldProduct)
			if err != nil {
				m.Logger.Error(err)
				return err
			}
		}
		if oldProduct.LastPrice != p.LastPrice {
			update := bson.D{
				{"$set", bson.D{
					{"name", p.Name},
					{"last_price", p.LastPrice},
					{"last_updated", p.LastUpdated},
				}},
				{"$inc", bson.D{{"price_changes", 1}}},
			}
			//res, err := col.InsertOne(ctx, p)
			_, err := col.UpdateOne(ctx, filter, update, opts)
			if err != nil {
				m.Logger.Error(err)
				return err
			}
		}
	}

	return nil
}
