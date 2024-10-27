package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer f.Close() // Не забываем закрыть файл после работы

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Десериализация JSON в срез структур Employee
	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %w", err)
	}

	// Если необходимо ограничить количество записей
	if limit > 0 && len(data) > limit {
		data = data[:limit]
	}

	// Возврат данных
	return data, nil
}
