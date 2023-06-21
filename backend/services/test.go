package services

import "github.com/iltyty/db_connect/backend_go/models"

func GetTestData() (data []models.APITest, err error) {
	data, err = models.GetTestData()
	if len(data) == 0 {
		models.CreateTestData()
		data, err = models.GetTestData()
	}
	return
}
