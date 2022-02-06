package dummy

import (
	"github.com/alvinarthas/myrepo/business/model"
)

func GetDummyList() ([]model.Dummy, error) {
	return getDummyListSQL()
}

func CreateDummy(payload model.CreateDummyRequest) error {
	return createDummySQL(payload)
}

func UpdateDummy(payload model.UpdateDummyRequest) error {
	return updateDummySQL(payload)
}
