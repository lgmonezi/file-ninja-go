package folders

// type fileIdentifier struct {
// 	root      string
// 	tree      *Folder
// 	count     atomic.Int32
// 	mutex     sync.Mutex
// 	waitGroup sync.WaitGroup
// 	errors    []*ScanError
// 	scanning  bool
// }

// func (s *fileIdentifier) identifyFiles() {
// 	filesys := os.DirFS(s.root)
// 	s.scanning = true
// 	s.identifyFilesInFS(filesys, ".", s.root, nil)
// 	s.waitGroup.Wait()
// 	s.scanning = false
// }

// func (s *fileIdentifier) identifyFilesInFS(filesys fs.FS, fileName string, dirPath string, parent *Folder) {
// 	s.waitGroup.Add(1)
// 	go func() {
// 		var folder Folder
// 		defer func() { s.waitGroup.Done() }()
// 		s.count.Add(1)
// 		entries, err := fs.ReadDir(filesys, fileName)
// 		if err != nil {
// 			s.appendError(fileName, err)
// 			return
// 		}
// 		for _, entry := range entries {
// 			if entry.IsDir() {
// 				subFileSys, err := fs.Sub(filesys, entry.Name())
// 				if err != nil {
// 					s.appendError(fileName, err)
// 					continue
// 				}
// 				s.identifyFilesInFS(subFileSys, ".", path.Join(dirPath, entry.Name()))
// 				continue
// 			}
// 			s.count.Add(1)
// 		}
// 	}()
// }

// func (s *fileIdentifier) Count() int {
// 	return s.count
// }

// func (s *fileIdentifier) Errors() []error {
// 	return s.errors
// }

// func (s *fileIdentifier) appendError(path string, err error) {
// 	scanError := ScanError{path, err}
// 	s.mutex.Lock()
// 	s.errors = append(s.errors, &scanError)
// 	s.mutex.Unlock()
// }
