package linux

import (
	"common/stringx"
	"common/system"
	"fmt"
	"github.com/jaypipes/ghw"
	"github.com/pkg/errors"
	"log"
	"path/filepath"
	"regexp"
)

type DiskSmart struct {
}

type Disk struct {
	DevName  string
	Wwn      string
	Model    string
	Firmware string
	Serial   string
	Size     uint64
	DevType  string

	PartID string
}

func (d Disk) DiskList() (disks []Disk, err error) {
	block, err := ghw.Block()
	if err != nil {
		return
	}
	var diskT Disk
	for _, disk := range block.Disks {
		flag, err := IsCanUseDisk(disk.Name)
		if err != nil {
			return nil, err
		}
		if flag {
			diskT = Disk{}
			diskT.DevName = disk.Name
			diskT.Model = disk.Model
			diskT.Serial = disk.SerialNumber
			diskT.Wwn = disk.WWN
			diskT.Size = disk.SizeBytes / 1024 / 1024 / 1024
			diskT.DevType = disk.Vendor
			disks = append(disks, diskT)
		}
	}
	return
}

func (d *Disk) MakePartition() (err error) {
	if d.DevName == "" {
		err = errors.Wrap(err, "dev name is nil")
		return
	}
	cmd := fmt.Sprintf("parted /dev/%s -s -- mklabel gpt mkpart primary 1 -1", d.DevName)
	_, err = system.ExecShellLinux(cmd)
	return
}

func (d *Disk) FillPartID() (err error) {
	cmd := fmt.Sprintf("blkid -o value -s PARTUUID /dev/%s1", d.DevName)
	ret, err := system.ExecShellLinux(cmd)
	if err != nil {
		return
	}
	d.PartID = ret
	return
}

func (d Disk) PartProbe() (err error) {
	_, err = system.ExecShellLinux("partprobe")
	return
}

func (d Disk) Print() {
	println(fmt.Sprintf("name: %s\nmodel: %s\nUUID: %s", d.DevName, d.Model, d.PartID))
}

func (d Disk) RemovePartition() (err error) {
	if d.DevName == "" {
		err = errors.Wrap(err, "dev name is nil")
		return
	}
	cmd := fmt.Sprintf("parted /dev/%s -s -- rm 1", d.DevName)
	_, err = system.ExecShellLinux(cmd)
	return
}

func ParseMountL() (osDisk []string, err error) {
	ret, _ := system.ExecShellLinux("mount -l")
	ret1 := stringx.SearchSplitStringColumn(ret, ".+ / .+", " ", 1)
	for _, val := range ret1 {
		flag, err := stringx.MatchStr(val, ".+mapper.+")
		if err != nil {
			return nil, err
		}
		if flag {
			pvRet, _ := system.ExecShellLinux("pvs --noheadings|awk -F'[0-9]' '{print$1}'")
			pvRet = filepath.Base(pvRet)
			osDisk = append(osDisk, pvRet)
		}
		flag, err = stringx.MatchStr(val, "/dev/sd.+")
		if err != nil {
			return nil, err
		}
		if flag {
			reg := regexp.MustCompile("sd[a-z]+")
			ret := reg.FindStringSubmatch(val)
			if len(ret) != 0 {
				osDisk = append(osDisk, ret[0])
			}
		}
	}
	return osDisk, err
}

func GetOSDisk() (devName string) {
	if devNameT, err := system.ExecShellLinux("df|grep -P '/$'|awk '{print$1}'|awk -F'1' '{print$1}'|awk -F'/' '{print$3}'"); err != nil {
		log.Println("error: ", err.Error())
		return ""
	} else {
		return devNameT
	}
}

func IsCanUseDisk(devName string) (b bool, err error) {
	osDevs, err := ParseMountL()
	for _, val := range osDevs {
		if val == "" {
			return false, err
		}
		if val == devName {
			return false, err
		}
		flag, err := stringx.MatchStr(devName, "sd.+")
		if flag {
			return true, err
		} else {
			return false, err
		}
	}
	return false, err
}
