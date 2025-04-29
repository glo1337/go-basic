package file

import (
	"errors"
	"os"
	"path"
)

func ReadFromFile(fileName string) ([]byte, error) {
	info, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("файл не существует")
		}
		return nil, errors.New("ошибка получения информации о файле")
	}

	if info.IsDir() {
		return nil, errors.New("ожидался файл, но это папка")
	}

	fileExt := path.Ext(fileName)
	if fileExt != ".json" {
		return nil, errors.New("неверный формат файла")
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("ошибка чтения файла")
	}
	return content, nil
}
