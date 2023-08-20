package repository

import (
	"database/sql"
	"github.com/lgmonezi/file-ninja-go/db"
)

type TblDirtyFiles struct {
	Id             uint64
	Name           string
	Path           string
	Type           string
	Clean          bool
	Md5sum         sql.NullString
	ParentFolderId sql.NullInt64
}

func GetDirtyRootFolders() ([]*TblDirtyFiles, error) {
	mapper := func(scan db.ScannerFunction) *TblDirtyFiles {
		var row TblDirtyFiles
		scan(&row.Id, &row.Name, &row.Path, &row.Type, &row.Clean, &row.Md5sum, &row.ParentFolderId)
		return &row
	}

	sqlQuery := `SELECT 
					id, name, path, type, clean, md5sum, parent_folder_id 
				 FROM tbl_dirty_files
				 WHERE parent_folder_id IS NULL`

	results, err := db.QueryDB(sqlQuery, mapper)
	if err != nil {
		return nil, err
	}
	return results, nil
}
