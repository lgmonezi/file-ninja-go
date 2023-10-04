package folders

type File struct {
	fileCommons
}

type Folder struct {
	Files   []*File
	Folders []*Folder
	fileCommons
	Parent *Folder
}

type fileCommons struct {
	name    string
	path    string
	modDate string
	error   error
	Parent  *Folder
}
