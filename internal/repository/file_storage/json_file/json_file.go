package json_file

import "C"
import (
	"io"
	"log"
	"os"
)

type JSONFile struct {
	NameFile string
}

func NewJSONFile(nameFile string) *JSONFile {
	return &JSONFile{NameFile: nameFile}
}

type JSONReader interface {
	ReadAll() ([]byte, error)
}

type JSONWriter interface {
	WriteAll([]byte) error
}

func (J *JSONFile) WriteAll(data []byte) error {
	err := error(nil)
	err = os.WriteFile(J.NameFile, data, 0666)
	if err != nil {
		panic(err)
	}
	return err
}

func (J *JSONFile) ReadAll() ([]byte, error) {
	var data []byte
	err := error(nil)
	file := new(os.File)

	file, err = os.Open(J.NameFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err = io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}
	return data, err
}
