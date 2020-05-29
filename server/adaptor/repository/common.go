package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func ErrorToIsExist(err error, resource string, args ...interface{}) (bool, error) {
	resource = fmt.Sprintf(resource, args...)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return true, errors.Wrapf(err, "failed to get %s", resource)
	}
	return true, nil
}
