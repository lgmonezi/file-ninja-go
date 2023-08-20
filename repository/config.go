package repository

import (
	"github.com/lgmonezi/file-ninja-go/db"
	_ "modernc.org/sqlite"
)

type TblConfig struct {
	Name  string
	Value string
}

func GetCleanFolder() (string, error) {
	mapper := func(scan db.ScannerFunction) string {
		var cleanFolder string
		scan(cleanFolder)
		return cleanFolder
	}

	sqlQuery := "SELECT value FROM tbl_config WHERE name = 'clean_folder'"

	results, err := db.QueryDB(sqlQuery, mapper)
	if err != nil {
		return "", err
	}
	return results[0], nil
}
