package linux

import (
	"common/stringx"
	"io/ioutil"
	"strings"
)

type CpuHwInfo struct {
	Model     string
	Count     int
	CoreCount int
	Stepping  string
}

func (e *CpuHwInfo) GetCpuHwInfo() (err error) {
	//tmpStr, err := ExecShellLinux("cat /proc/cpuinfo")
	tmp, err := ioutil.ReadFile("/proc/cpuinfo")
	tmpStr := strings.Replace(string(tmp), "\n", "", 1)
	if err != nil {
		return err
	}
	e.Model = stringx.SearchSplitStringColumnFirst(tmpStr, ".*model name.*", ":", 2)
	e.Stepping = stringx.SearchSplitStringColumnFirst(tmpStr, ".*stepping.*", ":", 2)
	countTmp1 := stringx.SearchString(tmpStr, ".*physical id.*")
	countTmp := stringx.UniqStringList(countTmp1)
	e.Count = len(countTmp)
	coreCountTmp1 := stringx.SearchString(tmpStr, ".*processor.*")
	coreCountTmp := stringx.UniqStringList(coreCountTmp1)
	e.CoreCount = len(coreCountTmp)
	return err
}
