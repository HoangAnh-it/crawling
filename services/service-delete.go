package services

import (
	"crawling/config"
	"crawling/utils"
)

func DeleteById(id string) {
	err := config.DBController.DeleteById(id)
	utils.CatchError(err)
}

func DeleteData(dataModel utils.DataModel) {
	err := config.DBController.Delete(dataModel)
	utils.CatchError(err)
}
