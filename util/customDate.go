package util

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// CustomDate - пользовательский тип времени
type CustomDate struct {
	time.Time
}

// MarshalJSON - переопределение метода MarshalJSON для нашего типа
func (cd *CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, cd.Format("2006-01-02"))), nil
}

// UnmarshalJSON - переопределение метода UnmarshalJSON для нашего типа
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// Распаковываем JSON строку и парсим ее как дату
	parsedTime, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}

	// Устанавливаем значение в наш пользовательский тип
	cd.Time = parsedTime
	return nil
}

// Scan - преобразование значения базы данных в пользовательский тип времени
func (cd *CustomDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		cd.Time = v
	case []byte:
		// Если значение представлено в виде байтов, то декодируем JSON строку
		err := json.Unmarshal(v, &cd)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("неподдерживаемый тип для сканирования в CustomDate: %T", value)
	}

	return nil
}

// Value - преобразование пользовательского типа времени в значение, понимаемое базой данных
func (cd CustomDate) Value() (driver.Value, error) {
	return cd.Time, nil
}
