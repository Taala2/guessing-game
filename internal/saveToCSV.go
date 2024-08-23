package internal

import (
	"os"
	"encoding/csv"
	"fmt"
	"strconv"
)

type GameRecord struct {
    Result   string 
    Attempts int   
}

func saveToCSV(records []GameRecord, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("не удалось создать файл: %v", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    headers := []string{"Result", "Attempts"}
    if err := writer.Write(headers); err != nil {
        return fmt.Errorf("не удалось записать заголовки: %v", err)
    }


    for _, record := range records {
        row := []string{record.Result, strconv.Itoa(record.Attempts)}
        if err := writer.Write(row); err != nil {
            return fmt.Errorf("не удалось записать данные: %v", err)
        }
    }

    return nil
}
