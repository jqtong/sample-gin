//Package utils tool bag
package utils

import (
	"errors"
	"fmt"
	"math"
	"nucarf.com/store_service/api/conf/initialize"
	"os"
	"reflect"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
	"nucarf.com/store_service/api/conf"
)

//ParseTimeDimension 根据时间 计算type类型 7天以内为 h, 大于七天小于365天为 d, 大于364天为 m
func ParseTimeDimension(start, end time.Time) string {
	var t string

	diffTime := end.Unix() - start.Unix()
	t = "h"

	if diffTime > 7*24*3600 {
		t = "d"
	}

	if diffTime > 364*24*3600 {
		t = "m"
	}

	return t
}

//MkDir 判断路径是否存在
func MkDir(path string) (err error) {
	_, err = os.Stat(path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		return
	}
	return
}

//Recover 异常处理
func Recover() {
	err := recover()
	conf.AppLog.WithFields(logrus.Fields{
		"error": err,
	}).Warn()
}

//ErrorLog error log
func ErrorLog(err error) {
	conf.AppLog.WithFields(logrus.Fields{
		"error": err,
	}).Warn()
}

//GetPercentWithPrecision eleme 最大饥饿算法
/**
 * @param {Array.<number>} valueList a list of all data 一列数据
 * @param {number} idx index of the data to be processed in valueList 索引值（数组下标）
 * @param {number} precision integer number showing digits of precision 精度值
 * @return {number} percent ranging from 0 to 100 返回百分比从0到100
 * eg. GetPercentWithPrecision([]int{3, 3, 3}, 0, 2)   // 33.34
 */
func GetPercentWithPrecision(valueList []int, idx int, precision int) (res float64) {
	if idx >= len(valueList) {
		res = 0
		return
	}
	var digits = math.Pow10(precision)
	var total = 0
	for _, value := range valueList {
		total += value
	}
	if total == 0 {
		return
	}
	var votesPerQuota = make([]float64, len(valueList))
	for i, value := range valueList {
		votesPerQuota[i] = float64(value) / float64(total) * digits * 100
	}

	var targetSeats = digits * 100
	var currentSeatsSum int
	var seats = make([]int, len(valueList))
	var remainder = make([]float64, len(valueList))
	for i, votes := range votesPerQuota {
		seats[i] = int(votes)
		currentSeatsSum += int(votes)
		remainder[i] = votes - float64(seats[i])
	}

	for currentSeatsSum < int(targetSeats) {
		var max = float64(0)
		var maxID = 0
		for i, remainVal := range remainder {
			if remainVal > max {
				max = remainVal
				maxID = i
			}
		}
		seats[maxID]++
		remainder[maxID] = 0
		currentSeatsSum++
	}

	res = float64(seats[idx]) / digits
	return
}

//GetRefererHost get referer host
func GetRefererHost(referer string) (host string, err error) {
	reg := regexp.MustCompile(`(http://|https://)?([^/]*)`)
	if reg == nil {
		err = errors.New("regexp error")
		return
	}
	result := reg.FindAllStringSubmatch(referer, -1)
	host = result[0][2]
	return
}

// CallUserFuncArray call func
func CallUserFuncArray(obj interface{}, name string, param []interface{}) (interface{}, bool) {
	t := reflect.TypeOf(obj)
	_func, ok := t.MethodByName(name)
	if !ok {
		fmt.Println("方法不存在")
		return nil, false
	}
	_param := make([]reflect.Value, len(param)+1)
	_param[0] = reflect.ValueOf(obj)
	for key, value := range param {
		_param[key+1] = reflect.ValueOf(value)
	}
	res := _func.Func.Call(_param)
	return res, false
}

//GetDebugOrProd return the debug or prod key
func GetDebugOrProd(key string) string {
	if initialize.ServerConf.RunMode == "debug" {
		key = "debug_" + key
	}

	return key
}
