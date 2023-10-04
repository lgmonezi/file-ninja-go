package folders

// type FileScanner struct {
// 	root string
// 	fileIdentifier
// }

// func NewFileScanner(root string) (*FileScanner, error) {
// 	info, err := os.Stat(root)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open file: %w", err)
// 	}
// 	if !info.IsDir() {
// 		return nil, core.NotADirectoryErr
// 	}
// 	if !filepath.IsAbs(root) {
// 		root, err = filepath.Abs(root)
// 	}
// 	if err != nil {
// 		return nil, fmt.Errorf("couldn't retrieve absolute path of %s: %w", root, err)
// 	}
// 	return &FileScanner{root: root}, nil
// }
