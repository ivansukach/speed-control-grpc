package repository

type Record struct {
	Date   string
	Number string
	Speed  float32
}
type Repository interface {
	Create(record *Record) error
	//ControlTime()
	//GetMinMax(date int64) (*Record, error)
	Listing(date string) (*[]Record, error)
}
