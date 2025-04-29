package storage

import (
	"binstruct/bins"
	"binstruct/file"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type BinsStorage interface {
	Read() (*bins.BinList, error)
	Write(content []byte) error
}

type FileStorage struct {
	FilePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{
		FilePath: filePath,
	}
}

func (storage *FileStorage) Write(content []byte) error {
	fileName := storage.FilePath
	if filepath.Ext(fileName) == "" {
		fileName = fileName + ".json"
	}

	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("ошибка при создании файла")
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return errors.New("ошибка при записи файла")
	}
	return nil
}

func (storage *FileStorage) Read() (*bins.BinList, error) {
	binList := bins.NewBinList()
	content, err := file.ReadFromFile(storage.FilePath)
	if err != nil {
		return nil, errors.New("произошла ошибка во время чтения")
	}

	if len(content) == 0 {
		return binList, nil
	}

	err = json.Unmarshal(content, binList)
	if err != nil {
		return nil, err
	}

	return binList, nil
}
