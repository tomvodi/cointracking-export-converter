package xml

import "github.com/tomvodi/cointracking-export-converter/internal/interfaces"

type FileFactory struct {
}

// nolint: ireturn
// linter exception is ok here, as the XMLFileFactory
// interface returns an interface here
func (x *FileFactory) NewXMLFile() interfaces.XMLFile {
	return newXMLFile()
}

func NewFileFactory() *FileFactory {
	return &FileFactory{}
}
