package fa

import "github.com/cayleygraph/quad/voc"

func init() {
	voc.RegisterPrefix(Prefix, NS)
}

const (
	NS     = "http://namespaces.lgmonezi.net/file-attributes"
	Prefix = "fa:"
)

const (
	Name = Prefix + "Name"
)
