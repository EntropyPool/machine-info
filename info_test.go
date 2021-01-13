package machinfo

import (
	"testing"
)

func TestReadMachineInfo(t *testing.T) {
	info, err := ReadMachineInfo()
	if nil != err {
		t.Errorf("CANNOT read machine info by [%v]", err)
	}
	if nil == info {
		t.Errorf("CANNOT read machine info by [return nil]")
	}
}
