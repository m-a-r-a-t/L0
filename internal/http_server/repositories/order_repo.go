package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/m-a-r-a-t/L0/internal/http_server/models"
)

type orderRepo struct {
	db *sql.DB
}

func (o *orderRepo) GetOrderById(id string) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := o.db.QueryRowContext(ctx, "SELECT")
	fmt.Println(row)
	return &models.Order{}, nil
}

func (o *orderRepo) InsertOrders(orders []*models.Order) error {
	var deliveryVals, itemsVals, orderVals, paymentVals []interface{}
	orderStr := `INSERT INTO "Order" (order_uid, track_number, entry, locale,
		internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES`
	deliveryStr := `INSERT INTO "Delivery" (order_uid, name, phone, zip, city, address, region, email) VALUES`
	itemsStr := `INSERT INTO "Items" (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES`

	paymentStr := `INSERT INTO "Payment" (transaction, request_id, currency, provider,
		 amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES`

	quriesMap := map[string]*string{"order": &orderStr, "delivery": &deliveryStr, "items": &itemsStr, "payment": &paymentStr}
	valuesMap := map[string]*[]interface{}{"order": &orderVals, "delivery": &deliveryVals, "items": &itemsVals, "payment": &paymentVals}

	for i, order := range orders {
		orderStr += getValuesFormatPostgreStr(i, 11)
		fmt.Println(getValuesFormatPostgreStr(i, 11))
		orderVals = append(orderVals, order.OrderUid, order.TrackNumber, order.Entry,
			order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService,
			order.Shardkey, order.SmID, order.DateCreated, order.OofShard)

		deliveryStr += getValuesFormatPostgreStr(i, 8)
		deliveryVals = append(deliveryVals, order.OrderUid, order.Delivery.Name,
			order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address,
			order.Delivery.Region, order.Delivery.Email)

		itemsStr += getValuesFormatPostgreStr(i, 12)
		itemsVals = append(itemsVals, order.OrderUid, order.Items.ChrtId, order.Items.TrackNumber,
			order.Items.Price, order.Items.Rid, order.Items.Name, order.Items.Sale, order.Items.Size,
			order.Items.TotalPrice, order.Items.NmId, order.Items.Brand, order.Items.Status)

		paymentStr += getValuesFormatPostgreStr(i, 10)
		paymentVals = append(paymentVals, order.Payment.Transaction, order.Payment.RequestId, order.Payment.Currency,
			order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost,
			order.Payment.GoodsTotal, order.Payment.CustomFee)

	}

	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, key := range []string{"order", "delivery", "items", "payment"} {
		query := quriesMap[key]
		fmt.Println("KEY:", key, ";", "Query", query)
		q := strings.TrimSuffix(*query, ",")
		values := valuesMap[key]
		if _, err := tx.Exec(q, *values...); err != nil {
			fmt.Println("Insert Error ", key, err)
			return err
		}

	}

	return tx.Commit()
}

func NewOrderRepo(db *sql.DB) *orderRepo {
	return &orderRepo{db: db}
}

func getValuesFormatPostgreStr(i int, length int) string {
	var str = "("
	for j := 0; j < length; j++ {
		str += "$" + strconv.Itoa((i*length)+1+j) + ","
	}

	str = strings.TrimSuffix(str, ",") + "),"

	return str
}
