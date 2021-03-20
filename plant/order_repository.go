package main

import (
	"database/sql"
	"fmt"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetNumIntersects(ident string, start string, end string) (string, error) {
	query := "SELECT COUNT(*) FROM porder WHERE ident_order=$1 AND ((start_order>$2 AND start_order<$3) OR (end_order>$2 AND end_order<$3))"
	row := r.db.QueryRow(query, ident, start, end)
	// if err != nil {
	// 	return "5", fmt.Errorf("error querying plants, err: %v", err)
	// }
	var num string
	err := row.Scan(&num)
	if err != nil {
		return "5", fmt.Errorf("could not close rows, err %v", err)
	}
	// err = rows.Close()
	// if err != nil {
	// 	return nil, fmt.Errorf("could not close rows, err %v", err)
	// }
	return num, nil
}
