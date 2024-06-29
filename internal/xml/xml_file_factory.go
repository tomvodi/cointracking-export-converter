package xml

import "github.com/tomvodi/cointracking-export-converter/internal/interfaces"

type xmlFFActory struct {
}

func (x *xmlFFActory) NewXmlFile() interfaces.XmlFile {
	return newXmlFile()
}

func NewXmlFileFactory() interfaces.XmlFileFactory {
	return &xmlFFActory{}
}
