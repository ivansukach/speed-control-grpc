package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	StartAccessHour   int
	StartAccessMinute int
	StartAccessSecond int
	EndAccessHour     int
	EndAccessMinute   int
	EndAccessSecond   int
}

func Load() (cfg Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg.StartAccessHour, err = strconv.Atoi(os.Getenv("START_ACCESS_HOUR"))
	if err != nil {
		log.Error(err)
	}
	cfg.StartAccessMinute, err = strconv.Atoi(os.Getenv("START_ACCESS_MINUTE"))
	if err != nil {
		log.Error(err)
	}
	cfg.StartAccessSecond, err = strconv.Atoi(os.Getenv("START_ACCESS_SECOND"))
	if err != nil {
		log.Error(err)
	}
	cfg.EndAccessHour, err = strconv.Atoi(os.Getenv("END_ACCESS_HOUR"))
	if err != nil {
		log.Error(err)
	}
	cfg.EndAccessMinute, err = strconv.Atoi(os.Getenv("END_ACCESS_MINUTE"))
	if err != nil {
		log.Error(err)
	}
	cfg.EndAccessSecond, err = strconv.Atoi(os.Getenv("END_ACCESS_SECOND"))
	if err != nil {
		log.Error(err)
	}
	log.Println("start access hour: ", cfg.StartAccessHour)
	log.Println("start access minute: ", cfg.StartAccessMinute)
	log.Println("start access second: ", cfg.StartAccessSecond)

	log.Println("end access hour: ", cfg.EndAccessHour)
	log.Println("end access minute: ", cfg.EndAccessMinute)
	log.Println("end access second: ", cfg.EndAccessSecond)
	return
}
