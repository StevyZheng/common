package bmc

import (
	"common/stringx"
	"common/system"
	"fmt"
	"strings"
)

type SMBmcSel struct {
	Date    string
	MsgType string
	Msg     string
}

type SMBmcInfo struct {
	BmcVer  string
	BmcDate string
	BmcMac  string
	BmcIp   string
	BmcSel  []SMBmcSel
}

func (e *SMBmcInfo) GetSMBmcInfo() (err error) {
	summaryOut, err := system.ExecShellLinux("./ipmicfg -summary")
	if err != nil {
		return err
	}
	summaryOut = stringx.Strip(summaryOut)
	lines := strings.Split(summaryOut, "\n")
	for index, line := range lines {
		if index > 1 {
			ret := strings.Split(line, ":")
			if len(ret) > 1 {
				value := ret[1]
				switch index {
				case 2:
					e.BmcIp = value
				case 3:
					e.BmcMac = value
				case 4:
					e.BmcVer = value
				case 5:
					e.BmcDate = value
				}
			}
		}
	}
	selOut, err := system.ExecShellLinux("./ipmicfg -sel list")
	if err != nil {
		return err
	}
	lines = strings.Split(selOut, "\n")
	var tmpSelStrList []string
	var tmpStr string
	for index, line := range lines {
		if index%2 == 0 {
			tmpStr = line
		} else {
			tmpStr = fmt.Sprintf("%s%s", tmpStr, line)
			tmpSelStrList = append(tmpSelStrList, tmpStr)
		}
	}
	for _, line := range tmpSelStrList {
		smBmcSel := SMBmcSel{}
		tmpSelList := strings.Split(line, "|")
		if len(tmpSelList) == 4 {
			smBmcSel.Date = tmpSelList[1]
			smBmcSel.MsgType = tmpSelList[2]
			smBmcSel.Msg = tmpSelList[3]
		}
		e.BmcSel = append(e.BmcSel, smBmcSel)
	}
	return err
}
