package orders

import (
	"database/sql"
	"log"
	"time"
)

type Orders struct {
	id int
	timestamp time.Time
	price float64
	category string
}

func GetAllOrders(ctx *sql.DB) []Orders {
	var res []Orders

	rows, err := ctx.Query("SELECT * FROM Orders")
	if err != nil {
		log.Fatal("Couldn't quert database: " + err.Error())
	}

	for rows.Next() {
		order := Orders {}
		rows.Scan(&order)
		res = append(res, order)
	}

	return res
}
