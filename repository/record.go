package repository

type Record struct {
	Date   int64
	Number string
	Speed  float32
}
type Repository interface {
	Create(record *Record) error
	//GetMinMax(date int64) (*Record, error)
	Listing() (*[]Record, error)
}
