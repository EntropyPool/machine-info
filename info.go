package machinfo

import (
	"errors"
	"github.com/EntropyPool/machine-spec"
)

type MachineInfo struct {
}

func ReadMachineInfo() (*MachineInfo, error) {
	spec, err := machspec.ReadMachineSpec()
	if nil != err {
		return nil, err
	}
	if nil == spec {
		return nil, errors.New("read machine info: cannot read machine spec")
	}
	return nil, nil
}
