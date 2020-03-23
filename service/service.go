package service

import (
	"github.com/ivansukach/speed-control-grpc/repository"
	log "github.com/sirupsen/logrus"
)

type SpeedLimitService struct {
	r repository.Repository
}

func (sls *SpeedLimitService) Create(record *repository.Record) error {
	return sls.r.Create(record)
}

//func (sls *SpeedLimitService) Update(book *repositories.Book) error {
//	return sls.r.Update(book)
//}
func (sls *SpeedLimitService) GetMinMax(date string) (*repository.Record, *repository.Record, error) {
	records, err := sls.r.Listing()
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	var min float32 = 999.9
	indexOfMin := 0
	var max float32 = 0.00
	indexOfMax := 0
	for index, value := range *records {
		if min > value.Speed {
			min = value.Speed
			indexOfMin = index
		}
		if max < value.Speed {
			max = value.Speed
			indexOfMax = index
		}
	}

	return &(*records)[indexOfMin], &(*records)[indexOfMax], nil
}
func (sls *SpeedLimitService) Listing(date string, speedLimit float32) (*[]repository.Record, error) {
	records, err := sls.r.Listing()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	selectionResult := make([]repository.Record, 0)
	for _, value := range *records {
		if value.Speed > speedLimit {
			selectionResult = append(selectionResult, value)
		}
	}
	return &selectionResult, nil
}

//func (sls *SpeedLimitService) Delete(id string) error {
//	return sls.r.Delete(id)
//}
func New(repo repository.Repository) *SpeedLimitService {
	return &SpeedLimitService{r: repo}
}
