package bmc

import (
	"common/stringx"
	"common/system"
)

type IBmcInfo interface {
	GetSMBmcInfo() error
}

type VendorInfo struct {
	ProductModel string
	SN           string
	BiosVer      string
	BiosDate     string
	SysMac       []string
	BmcInfo      IBmcInfo
}

func (e *VendorInfo) GetVendorInfo() (err error) {
	dmiStr, err := system.ExecShellLinux("dmidecode")
	if err != nil {
		return err
	}
	e.ProductModel = stringx.SearchSplitStringColumnFirst(dmiStr, ".*Product Name.*", ":", 2)
	e.BiosVer = stringx.SearchSplitStringColumnFirst(dmiStr, ".*Version.*", ":", 2)
	e.BiosDate = stringx.SearchSplitStringColumnFirst(dmiStr, ".*Release Date.*", ":", 2)
}
