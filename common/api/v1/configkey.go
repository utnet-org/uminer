package v1

import (
	"regexp"
	"uminer/common/errors"
	"uminer/common/utils"
)

const (
	ConfigKeyTypeInput    = "input"
	ConfigKeyTypeRadio    = "radio"
	ConfigKeyTypeCheckBox = "checkbox"
)

func (k *ConfigKey) ValidateValue(v string) error {
	if k.Required && v == "" {
		return errors.Errorf(nil, errors.ErrorConfigValueValidateFailed)
	}

	if (k.Type == ConfigKeyTypeRadio || k.Type == ConfigKeyTypeCheckBox) && v != "" &&
		!utils.StringSliceContainsValue(k.Options, v) {
		return errors.Errorf(nil, errors.ErrorConfigValueValidateFailed)
	}

	if k.Regexp != "" {
		matched, err := regexp.MatchString(k.Regexp, v)
		if err != nil || !matched {
			return errors.Errorf(nil, errors.ErrorConfigValueValidateFailed)
		}
	}

	return nil
}
