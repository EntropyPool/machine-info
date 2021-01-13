package machinfo

import (
	"bytes"
	"errors"
	"github.com/EntropyPool/machine-spec"
	"github.com/elastic/go-sysinfo"
	"net"
	"os/exec"
	"strings"
)

type MemoryInfo struct {
	PhyMemoryBytes uint64 `json:"phy_mem_bytes"`
	VirMemoryBytes uint64 `json:"vir_mem_bytes"`
	HugepageSize   uint64 `json:"hugepage_size"`
	Hugepages      int    `json:"hugepages"`
}

type EthernetInfo struct {
	MacAddress       net.HardwareAddr `json:"mac"`
	PrivateAddresses []net.Addr       `json:"private_ip"`
}

type MachineInfo struct {
	PublicIPAddr *net.IPAddr           `json:"public_ip"`
	MachineSpec  *machspec.MachineSpec `json:"machine_spec"`
	EthernetInfo []EthernetInfo        `json:"ethernet"`
	MemoryInfo   *MemoryInfo           `json:"memory_info"`
}

func ReadEthernetInfos() ([]EthernetInfo, error) {
	interfaces, err := net.Interfaces()
	if nil != err {
		return nil, err
	}

	infos := make([]EthernetInfo, 0)
	for _, it := range interfaces {
		validInterface := true

		addrs, err := it.Addrs()
		if nil != err {
			return nil, err
		}

		for _, addr := range addrs {
			ipNet, validIPNet := addr.(*net.IPNet)
			if !validIPNet || ipNet.IP.IsLoopback() || nil == ipNet.IP.To4() {
				validInterface = false
			}
		}

		if !validInterface {
			continue
		}

		addresses := make([]net.Addr, 0)
		for _, addr := range addrs {
			addresses = append(addresses, addr)
		}

		infos = append(infos, EthernetInfo{
			MacAddress:       it.HardwareAddr,
			PrivateAddresses: addresses,
		})
	}

	return infos, nil
}

func ReadPublicIPAddr() (*net.IPAddr, error) {
	cmd := exec.Command("dig", "+short", "myip.opendns.com", "@resolver1.opendns.com")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if nil != err {
		return nil, err
	}
	return net.ResolveIPAddr("ip", strings.TrimSpace(out.String()))
}

func ReadMemoryInfo() (*MemoryInfo, error) {
	host, err := sysinfo.Host()
	if nil != err {
		return nil, err
	}
	memory, err := host.Memory()
	if nil != err {
		return nil, err
	}

	var info MemoryInfo
	info.PhyMemoryBytes = memory.Total
	info.VirMemoryBytes = memory.VirtualTotal
	info.HugepageSize = memory.Metrics["Hugepagesize"]
	info.Hugepages = int(memory.Metrics["HugePages_Total"])

	return &info, nil
}

func ReadMachineInfo() (*MachineInfo, error) {
	spec, err := machspec.ReadMachineSpec()
	if nil != err {
		return nil, err
	}
	if nil == spec {
		return nil, errors.New("read machine info: cannot read machine spec")
	}

	pubIPAddr, err := ReadPublicIPAddr()
	if nil != err {
		return nil, err
	}

	ethInfos, err := ReadEthernetInfos()
	if nil != err {
		return nil, err
	}

	memInfo, err := ReadMemoryInfo()
	if nil != err {
		return nil, err
	}

	return &MachineInfo{
		PublicIPAddr: pubIPAddr,
		MachineSpec:  spec,
		EthernetInfo: ethInfos,
		MemoryInfo:   memInfo,
	}, nil
}
