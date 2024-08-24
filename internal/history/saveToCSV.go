package history

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

var Headers = []string{"Result", "Attempts"}

func SaveToCSV(records []GameRecord, filename string) error {

    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    info, err := file.Stat()
    if err != nil {
        return fmt.Errorf("failed to get file info: %v", err)
    }

    if info.Size() == 0 {
        if err := writer.Write(Headers); err != nil {
            return fmt.Errorf("failed to write headers: %v", err)
        }
    }


    for _, record := range records {
        row := []string{record.Result, strconv.Itoa(record.Attempts)}
        if err := writer.Write(row); err != nil {
            return fmt.Errorf("failed to write row %v: %v", row, err)
        }
    }

    return nil
}

func ReadCSV(filepath string) ([]GameRecord, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.FieldsPerRecord = len(Headers)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("failed to read CSV: %v", err)
    }


    var gameRecords []GameRecord
    for _, record := range records[1:] {
        attempts, err := strconv.Atoi(record[1])
        if err != nil {
            return nil, fmt.Errorf("failed to convert attempts to int: %v", err)
        }
        gameRecords = append(gameRecords, GameRecord{
            Result:   record[0],
            Attempts: attempts,
        })
    }

    return gameRecords, nil
}
