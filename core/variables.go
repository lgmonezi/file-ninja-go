package core

import "errors"

var ScanningInProgress = errors.New("scanning in progress")
var NotADirectoryErr = errors.New("the file is not a directory")
