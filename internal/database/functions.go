package database

import (
	"database/sql"
	"github.com/lgmonezi/file-ninja-go/utils"
)

type RowMapper[T any] func(scanner ScannerFunction) T
type ScannerFunction func(dest ...any)

func QueryDB[T any](sql string, mapper RowMapper[T]) ([]T, error) {
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer utils.CloseWarningError(rows)

	var results []T
	for rows.Next() {
		scannerFunction, errPtr := getScannerFunction(rows)
		item := mapper(scannerFunction)
		if *errPtr != nil {
			return nil, *errPtr
		}
		results = append(results, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func getScannerFunction(rows *sql.Rows) (ScannerFunction, *error) {
	var err error
	scanner := func(dest ...any) {
		err = rows.Scan(dest...)
	}
	return scanner, &err
}
