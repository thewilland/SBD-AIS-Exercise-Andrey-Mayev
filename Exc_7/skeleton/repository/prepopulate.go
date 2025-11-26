package repository

import (
	"context"
	"log/slog"
	"math/rand"
	"ordersystem/model"
	"ordersystem/storage"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

func Prepopulate(db *DatabaseHandler, s3 *minio.Client) error {
	// check if prepopulate has already run once
	var exists bool
	err := db.dbConn.Model(&model.Drink{}).
		Select("count(*) > 0").
		Find(&exists).
		Error
	if err != nil {
		return err
	}
	if exists {
		// don't prepopulate if has already run
		return nil
	}
	slog.Info("Prepopulating database and S3")
	// create drink menu
	drinks := []model.Drink{
		{Name: "Water", Price: 1.80, Description: "Yay water"},
		{Name: "Evil Water", Price: 6.66, Description: "AAAA OH NO"},
		{Name: "Leaf Water Hot", Price: 2.80, Description: "Oh that's tea mmmm AAAAA HOT OH NO"},
	}
	err = db.dbConn.Create(drinks).Error
	if err != nil {
		return err
	}
	// create orders
	var orders []model.Order
	for _, drink := range drinks {
		for i := 0; i < 15; i++ {
			order := model.Order{
				Base: model.Base{
					CreatedAt: time.Now().Add(time.Duration(rand.Intn(30)) * time.Minute),
				},
				Amount:  uint64(rand.Intn(5)),
				DrinkID: drink.ID,
			}
			orders = append(orders, order)
		}
	}
	err = db.dbConn.Create(orders).Error
	if err != nil {
		return err
	}
	// store orders to s3
	for _, order := range orders {
		markdown := order.ToMarkdown()
		receiptReader := strings.NewReader(markdown)
		_, err = s3.PutObject(context.Background(), storage.OrdersBucket, order.GetFilename(), receiptReader, int64(len(markdown)),
			minio.PutObjectOptions{ContentType: "text/markdown"})
		if err != nil {
			return err
		}
	}

	return nil
}
