package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/OneOfOne/xxhash"
	"github.com/cayleygraph/cayley"
	_ "github.com/cayleygraph/cayley/graph/sql/sqlite"
	"github.com/cayleygraph/quad"
	"github.com/lgmonezi/file-ninja-go/database/graph/vocabulary/fa"
	_ "github.com/lgmonezi/file-ninja-go/database/graph/vocabulary/fa"
)

var client *cayley.Handle

type fileSerialization struct {
	name       string
	_type      string
	extension  string
	hash       string
	compressed bool
}

func init() {
	var err error
	client, err = cayley.NewGraph("sqlite", "graph.db", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := getFile()
	if err != nil {
		log.Fatal("Error:", err)
	}
	_, err = hashFile(file)
	f := fileSerialization{name: file.Name()}
	client.AddQuadSet(createFileQuad(&f))
	fmt.Println(*file)
}

func createFileQuad(f *fileSerialization) []quad.Quad {
	file := quad.RandomBlankNode()
	quads := []quad.Quad{}
	quads = append(quads, quad.Make(file, quad.IRI(fa.Name), f.name, nil))
	return quads
}

func getFile() (file *os.File, err error) {
	const path = "/mnt/volume/Books/IT/Programming/Go/The Go Programming Language - Alan A. A. Donovan, Brian W. Kernighan - Addison-Wesley Professional Computing - (2015).epub"
	file, err = os.Open(path)
	return
}

func hashFile(file *os.File) (sum string, err error) {
	hash := xxhash.New64()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}
	sum = fmt.Sprintf("%x", hash.Sum64())
	return
}
