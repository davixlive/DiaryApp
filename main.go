package main

import (
	"errors"
	"os"
	"time"
)

func createDay() {
	time := time.Now()
	_, err := os.Create(time.Format("2006/01/02"))
	if err != nil {
		panic(err)
	}
}
func existsDay(day string) (bool, error) {
	_, err := os.Open(day)
	if errors.Is(err, os.ErrNotExist) {
		return false, errors.New("file not exists error")
	} else {
		return true, err
	}
}
func modifyDay(day string) {
	exist, _ := existsDay(day)
	if exist {

	}
}
func main() {

}
