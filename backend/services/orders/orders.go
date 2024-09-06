package orders

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type Orders struct {
	Id         int
	Created_on time.Time
	Price      float64
	Category   string
}

func GetAllOrders(ctx *sqlx.DB) []Orders {
	res := []Orders{}

	err := ctx.Select(&res, "SELECT * FROM orders")

	if err != nil {
		log.Fatal("Error querying DB: " + err.Error())
	}

	return res
}
