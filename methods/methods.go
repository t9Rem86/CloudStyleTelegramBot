package methods

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"time"
	_ "time"
)

func InsertDebit(conn pgx.Conn, idBalance int, idType int, idModel int, idBrand int, price float32, count int, delivery float32) error {
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO debit(id_debit, id_balance, id_type, id_brand, id_model, price, count, delivery, date) 
		 VALUES (default, $1, $2, $3, $4, $5, $6, $7, $8)`,
		idBalance, idType, idBrand, idModel, price, count, delivery, time.Now(),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка вставки операции по дебету: %v\n", err)
	}
	return err
}

func InsertCredit(conn pgx.Conn, idBalance int, idType int, idModel int, idBrand int, price float32, count int, delivery float32) error {
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO credit(id_credit, id_balance, id_type, id_brand, id_model, price, count, date) 
		 VALUES (default, $1, $2, $3, $4, $5, $6, $7)`,
		idBalance, idType, idBrand, idModel, price, count, time.Now(),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка вставки операции по кредиту: %v\n", err)
	}
	return err
}
