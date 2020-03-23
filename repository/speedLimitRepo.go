package repository

import (
	"fmt"
	"io"
	"log"
	"os"
)

func New(date *string) Repository {
	file, err := os.Create(*date + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	return &speedLimitRepository{file: file}
}

type speedLimitRepository struct {
	file *os.File
}

func (slr *speedLimitRepository) Create(record *Record) error {
	str := fmt.Sprintf("%s | %s | %f\n", record.Date, record.Number, record.Speed)
	_, err := slr.file.WriteString(str)
	return err

}
func (slr *speedLimitRepository) Listing() (*[]Record, error) {
	data := make([]byte, 64)

	for {
		n, err := slr.file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		} else if err != nil {
			return nil, err
		}
		fmt.Print(string(data[:n]))
	}
	records := make([]Record, 0)
	return &records, nil
}

//func (slr *speedLimitRepository) GetMinMax(date int64) (*Record, error) {
//	data := make([]byte, 64)
//
//	for{
//		n, err := slr.file.Read(data)
//		if err == io.EOF{   // если конец файла
//			break           // выходим из цикла
//		}else if err!=nil{
//			return nil, err
//		}
//		fmt.Print(string(data[:n]))
//	}
//	record:=Record{}
//	return &record, nil
//}
