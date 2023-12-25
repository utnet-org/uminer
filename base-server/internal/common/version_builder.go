package common

import (
	"fmt"
	"strconv"
	"uminer/common/errors"
)

const (
	VERSION_PREFIX byte = 'V'
)

func VersionStrParse(version string) (int64, error) {
	if len(version) <= 0 {
		err := errors.Errorf(nil, errors.ErrorVersionInvalid)
		return 0, err
	}
	if version[0] != VERSION_PREFIX {
		err := errors.Errorf(nil, errors.ErrorVersionInvalid)
		return 0, err
	}

	versionInt, err := strconv.Atoi(version[1:])
	if err != nil {
		err := errors.Errorf(err, errors.ErrorVersionInvalid)
		return 0, err
	}

	return int64(versionInt), nil
}

func VersionStrBuild(version int64) string {
	return fmt.Sprintf("%c%d", VERSION_PREFIX, version)
}
