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

func TestReadPublicIPAddr(t *testing.T) {
	addr, err := ReadPublicIPAddr()
	if nil != err {
		t.Errorf("CANNOT read public ip by [%v]", err)
	}
	if nil == addr {
		t.Errorf("CANNOT read public ip by [return nil]")
	}
}

func TestReadSysinfo(t *testing.T) {
	info, err := ReadMemoryInfo()
	if nil != err {
		t.Errorf("CANNOT read memory info by [%v]", err)
	}
	if nil == info {
		t.Errorf("CANNOT read memory info by [return nil]")
	}
}
