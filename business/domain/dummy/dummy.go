package dummy

import (
	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/utils/common"
)

func GetDummyList() ([]model.Dummy, common.Error) {
	return getDummyListSQL()
}

func GetDummy(id string) (model.Dummy, common.Error) {
	return getDummySQL(id)
}

func CreateDummy(payload model.CreateDummyRequest) common.Error {
	return createDummySQL(payload)
}

func UpdateDummy(payload model.UpdateDummyRequest) common.Error {
	return updateDummySQL(payload)
}

func DeleteDummy(id string) common.Error {
	return deleteDummySQL(id)
}
