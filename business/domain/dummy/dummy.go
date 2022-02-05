package dummy

import (
	"github.com/alvinarthas/myrepo/business/model"
)

func GetDummyList() ([]model.Dummy, error) {
	return getDummyListSQL()
}
