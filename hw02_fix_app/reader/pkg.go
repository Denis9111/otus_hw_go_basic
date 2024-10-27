package reader

import (
	"encoding/json"
	"fmt"
	"github.com/fixme_my_friend/hw02_fix_app/types"
	"io"
	"os"
)

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	defer f.Close() // Не забываем закрыть файл после работы

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Десериализация JSON в срез структур Employee
	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %v", err)
	}

	// Если необходимо ограничить количество записей
	if limit > 0 && len(data) > limit {
		data = data[:limit]
	}

	// Возврат данных
	return data, nil
}
