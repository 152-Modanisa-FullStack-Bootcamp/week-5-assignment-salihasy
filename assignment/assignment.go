package assignment

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y
	if sum < x {
		return sum, true
	} else {
		return sum, false
	}
}

func CeilNumber(f float64) float64 {
	i, float := math.Modf(f)
	decimalValAsInt := math.Ceil(float * 100)

	switch {
	case decimalValAsInt > 75:
		{
			roundedValue := math.Ceil(f)
			return roundedValue
		}
	case decimalValAsInt > 50:
		{
			roundedValue := i + 0.75
			return roundedValue
		}
	case decimalValAsInt > 25:
		{
			roundedValue := i + 0.50
			return roundedValue
		}
	case decimalValAsInt > 0:
		{
			roundedValue := i + 0.25
			return roundedValue
		}
	default:
		return f
	}
	fmt.Println(decimalValAsInt)
	return decimalValAsInt
}

func AlphabetSoup(s string) string {
	a := strings.Split(s, "")
	sort.Strings(a)
	joins := strings.Join(a, "")
	return joins
}

func StringMask(s string, n uint) string {
	a := strings.Split(s, "")
	length := len(a)
	nint := int(n)

	if nint >= length || nint == 0 {
		for i := 0; i < length; i++ {
			a[i] = "*"
		}
	} else {
		for i := nint; i < length; i++ {
			a[i] = "*"
		}
	}
	w := strings.Join(a, "")
	return w
}

func WordSplit(arr [2]string) string {
	find := []string{}
	myArr := arr[0]
	words := arr[1]
	wordsArr := strings.Split(words, ",")
	for _, v := range wordsArr {
		if strings.Contains(myArr, v) {
			el := []string{v}
			find = append(el, find...)
			myArr = strings.ReplaceAll(myArr, v, "")
		}
	}
	if len(find) == 0 || len(myArr) > 0 {
		return "not possible"
	} else {
		return strings.Join(find, ",")
	}
}

func VariadicSet(i ...interface{}) []interface{} {
	tempMap := make(map[interface{}]bool)
	var uniqueValues []interface{}
	for k := 0; k < len(i); k++ {
		if _, val := tempMap[i[k]]; !val {
			tempMap[i[k]] = true
			uniqueValues = append(uniqueValues, i[k])
		}
	}
	return uniqueValues
}
