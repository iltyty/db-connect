package models

import (
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Test struct {
	gorm.Model

	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type APITest struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

const InitDataSize = 100
const deltaTimestamp = 1000

var now = time.Now().UnixMilli()

func GetTestData() (data []APITest, err error) {
	res := db.Model(&Test{}).Find(&data)
	if res.Error != nil {
		err = res.Error
	}
	return
}

func CreateTestData() {
	var data []Test
	for i := 0; i < InitDataSize; i++ {
		value := 20 + 10*rand.Float64()
		timestamp := now + int64(i*deltaTimestamp)
		data = append(data, Test{Timestamp: timestamp, Value: value})
	}
	db.Create(&data)
}
