package database

import "strings"

func GetCleanFolder() (string, error) {
	mapper := func(scan ScannerFunction) string {
		var cleanFolder string
		scan(cleanFolder)
		return cleanFolder
	}
	sql := "SELECT value FROM tbl_config WHERE name = 'clean_folder'"
	results, err := QueryDB(sql, mapper)
	if err != nil {
		return "", err
	}
	return results[0], nil
}

func GetDirtyRootFolders() ([]*TblDirtyFiles, error) {
	mapper := func(scan ScannerFunction) *TblDirtyFiles {
		var row TblDirtyFiles
		scan(&row.Id, &row.Name, &row.Path, &row.Type, &row.Clean, &row.Md5sum, &row.ParentFolderId, &row.Mtime)
		return &row
	}
	sqlSb := strings.Builder{}
	sqlSb.WriteString("SELECT id, name, path, type, clean, md5sum, parent_folder_id ")
	sqlSb.WriteString("FROM tbl_dirty_files WHERE parent_folder_id IS NULL")
	results, err := QueryDB(sqlSb.String(), mapper)
	if err != nil {
		return nil, err
	}
	return results, nil
}
