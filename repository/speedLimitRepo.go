package repository

import (
	"fmt"
	"github.com/ivansukach/speed-control-grpc/config"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func New() Repository {
	cfg := config.Load()
	return &speedLimitRepository{cfg: &cfg}
}

type speedLimitRepository struct {
	cfg *config.Config
}

func (slr *speedLimitRepository) CreateFile(fileName string) *os.File {
	log.Info("Creating new .txt file: " + fileName + " with specified access")
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	//mtime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
	//	slr.cfg.EndAccessHour, slr.cfg.EndAccessMinute, slr.cfg.EndAccessSecond, 0, time.Local)
	mtime := time.Now().Add(1 * time.Minute)
	log.Println("Modification to:", mtime)
	//atime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
	//	slr.cfg.EndAccessHour, slr.cfg.EndAccessMinute, slr.cfg.EndAccessSecond, 0, time.Local)
	atime := time.Now().Add(2 * time.Minute)
	log.Println("Access to:", atime)
	if err := os.Chtimes(fileName, atime, mtime); err != nil {
		log.Fatal(err)
	}
	//Why it doesn't limit access to file?
	return file
}
func (slr *speedLimitRepository) Create(record *Record) error {
	str := fmt.Sprintf("%s | %s | %f\n", record.Date, record.Number, record.Speed)
	onlyDate := strings.Split(record.Date, " ")[0]
	file, err := os.OpenFile(onlyDate+".txt", os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			file = slr.CreateFile(onlyDate + ".txt")
		} else {
			log.Fatal(err)
		}

	}
	_, err = file.WriteString(str)
	defer file.Close()
	return err
}
func (slr *speedLimitRepository) Listing(date string) (*[]Record, error) {
	data := make([]byte, 64)
	file, err := os.Open(date + ".txt")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	fmt.Println("start")
	fileContent := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		} else if err != nil {
			return nil, err
		}
		fmt.Print(string(data[:n]))
		fileContent += string(data[:n])
	}
	fmt.Println("end")
	separateRecords := strings.Split(fileContent, "\n")
	separateRecords = separateRecords[:len(separateRecords)-1]
	records := make([]Record, 0)
	for index, value := range separateRecords {
		fmt.Println("index:", index, "value:", value)
		fields := strings.Split(value, " | ")
		speed, err := strconv.ParseFloat(fields[2], 32)
		if err != nil {
			return nil, err
		}
		records = append(records, Record{Date: fields[0], Number: fields[1],
			Speed: float32(speed)})
	}

	return &records, nil
}
