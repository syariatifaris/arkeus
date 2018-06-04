//Package util is used for utility func, like manipulate array, round down, http
package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//JoinArr2String is func for join array to string
func JoinArr2String(array interface{}, separator string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(array), " ", separator, -1), "[]")
}

//InArray return true if element exist in array
func InArray(array interface{}, val interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			//DeepEqual return true if two var(value and type) are exact identical
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

//InArrayStr checks whether a string is in string slice
func InArrayStr(data string, slice []string) bool {
	for _, s := range slice {
		if s == data {
			return true
		}
	}

	return false
}

// GetUniqueInt64Array get unique elements of an int64 array
func GetUniqueInt64Array(inputArray []int64) []int64 {
	mapper := make(map[int64]int64)
	resultArr := make([]int64, 0)
	for i := 0; i < len(inputArray); i++ {
		if _, ok := mapper[inputArray[i]]; !ok {
			mapper[inputArray[i]] = 1
			resultArr = append(resultArr, inputArray[i])
		}
	}
	return resultArr
}

//CompareArray comparing 2 array with same length and type
func CompareArray(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a).Kind() != reflect.Slice || reflect.TypeOf(b).Kind() != reflect.Slice {
		return false, fmt.Errorf("Not An Array")
	}

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false, fmt.Errorf("Not An Array")
	}

	sA := reflect.ValueOf(a)
	sB := reflect.ValueOf(b)

	if sA.Len() == 0 && sB.Len() == 0 {
		return true, nil
	}

	if sA.Len() == 0 || sB.Len() == 0 {
		return false, fmt.Errorf("One of array is nil")
	}

	if sA.Len() != sB.Len() {
		return false, fmt.Errorf("Comparing different length of array")
	}

	for i := 0; i < sA.Len(); i++ {
		if !reflect.DeepEqual(sB.Index(i).Interface(), sA.Index(i).Interface()) {
			return false, fmt.Errorf("Different Element of array")
		}
	}

	return true, nil
}

// SplitStringIntoArrayInt64 split string into array int64
// example : "1,2,3,4" will be []int64{1,2,3,4}
func SplitStringIntoArrayInt64(input string) []int64 {
	res := make([]int64, 0)
	if "" == input {
		return res
	}
	arrStr := strings.Split(input, ",")
	for _, val := range arrStr {
		convertedVal, _ := strconv.ParseInt(strings.Trim(val, " "), 10, 64)
		res = append(res, convertedVal)
	}
	return res
}
