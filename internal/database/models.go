package database

import "database/sql"

type TblConfig struct {
	Name  string
	Value string
}

type TblDirtyFiles struct {
	Id             uint64
	Name           string
	Path           string
	Type           string
	Clean          bool
	Md5sum         sql.NullString
	ParentFolderId sql.NullInt64
	Mtime          string
}
