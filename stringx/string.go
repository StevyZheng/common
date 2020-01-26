package stringx

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//去除首尾的trimStr字符串
func Trim(srcStr string, trimStr string) string {
	regStrTmp := fmt.Sprintf("^%[1]s*|%[1]s*$", trimStr)
	re := regexp.MustCompile(regStrTmp)
	ret := re.ReplaceAllString(srcStr, "")
	return ret
}

//去除首尾的无效字符
func Strip(srcStr string) string {
	if srcStr == "" {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(srcStr, "")
}

//匹配字符串
//regStr: reg字符串
func MatchStr(srcStr string, regStr string) (flag bool, err error) {
	regStrTmp := fmt.Sprintf("(?m:%s)", regStr)
	flag, err = regexp.MatchString(regStrTmp, srcStr)
	return flag, err
}

//获取正则表达式所表示的字符串列表
func SearchString(srcStr string, regStr string) []string {
	regStr1 := fmt.Sprintf("(?m:%s)", regStr)
	re := regexp.MustCompile(regStr1)
	return re.FindAllString(srcStr, -1)
}

//获取正则表达式所表示的字符串，并按splitStr分割后的二维列表
func SearchSplitString(srcStr string, regStr string, splitStr string) [][]string {
	re := SearchString(srcStr, regStr)
	var ret [][]string
	for _, v := range re {
		vRe := strings.Split(v, splitStr)
		ret = append(ret, vRe)
	}
	return ret
}

func SearchSplitStringColumn(srcStr string, regStr string, splitStr string, col int) []string {
	tmp := SearchSplitString(srcStr, regStr, splitStr)
	var ret []string
	for _, v := range tmp {
		ret = append(ret, Trim(v[col-1], " "))
	}
	return ret
}

//获取正则表达式所表示的第一个字符串
func SearchStringFirst(srcStr string, regStr string) string {
	regStr1 := fmt.Sprintf("(?m:%s)", regStr)
	re := regexp.MustCompile(regStr1)
	findStr := re.FindAllString(srcStr, -1)
	if findStr != nil {
		return findStr[0]
	} else {
		return ""
	}
}

//获取正则表达式所表示的字符串，并按splitStr分割后的二维列表的第一列
func SearchSplitStringFirst(srcStr string, regStr string, splitStr string) []string {
	re := SearchStringFirst(srcStr, regStr)
	if re == "" {
		return nil
	}
	var ret []string
	ret = strings.Split(re, splitStr)
	return ret
}

func SearchSplitStringColumnFirst(srcStr string, regStr string, splitStr string, col int) string {
	tmp := SearchSplitStringFirst(srcStr, regStr, splitStr)
	if tmp == nil {
		return ""
	}
	return Trim(tmp[col-1], " ")
}

func UniqStringList(strList []string) []string {
	var newArr []string
	sort.Strings(strList)
	for i := 0; i < len(strList); i++ {
		repeat := false
		for j := i + 1; j < len(strList); j++ {
			if strList[i] == strList[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, strList[i])
		}
	}
	return newArr
}

func StrToInt(src string) (ret int, err error) {
	ret, err = strconv.Atoi(src)
	if err != nil {
		return -1, err
	}
	return ret, err
}

func StrToInt64(src string) (ret int64, err error) {
	ret, err = strconv.ParseInt(src, 10, 64)
	if err != nil {
		return -1, err
	}
	return ret, err
}

func IntToStr(src int) string {
	tmp := strconv.Itoa(src)
	return tmp
}

func Int64ToStr(src int64) string {
	tmp := strconv.FormatInt(src, 10)
	return tmp
}
