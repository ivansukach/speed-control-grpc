package repositories

type Book struct {
	Id            string `db:"id"`
	Title         string `db:"title"`
	Author        string `db:"author"`
	Genre         string `db:"genre"`
	Edition       string `db:"edition"`
	NumberOfPages int32  `db:"numberofpages"`
	Year          int32  `db:"year"`
	Amount        int32  `db:"amount"`
	IsPopular     bool   `db:"ispopular"`
	InStock       bool   `db:"instock"`
}
type Repository interface {
	Create(book *Book) error
	Read(id string) (*Book, error)
	Update(book *Book) error
	Delete(id string) error
}
