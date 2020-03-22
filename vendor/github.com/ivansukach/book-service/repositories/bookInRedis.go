package repositories

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"log"
)

func New(client *redis.Client) Repository {
	return &bookRepository{client: client}
}

type bookRepository struct {
	client *redis.Client
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
func (br *bookRepository) Create(book *Book) error {
	str, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return br.client.Set(book.Id, str, 0).Err()

}
func (br *bookRepository) Read(id string) (*Book, error) {
	str, err := br.client.Get(id).Result()
	book := Book{}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(str), &book)
	return &book, err
}
func (br *bookRepository) Update(book *Book) error {
	str, err := json.Marshal(book)
	if err != nil {
		return err
	}
	log.Println(str)
	return br.client.Set(book.Id, str, 0).Err()
}

func (br *bookRepository) Delete(id string) (err error) {
	_, err = br.client.Del(id).Result()
	return
}
