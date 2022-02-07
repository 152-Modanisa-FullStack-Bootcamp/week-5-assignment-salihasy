package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y // sum is overflow
	if sum < x { // if sum < x, overflow is true ( because sum can never be less than x and y)
		return sum, true
	} else {
		return sum, false
	}
}

func CeilNumber(f float64) float64 {
	i, float := math.Modf(f) //  i = integer part of the passed parameter, float = fractional part of the passed parameter as a float64 number.
	// if f = 22.85 , i=22 , f= 0.8500000000000014
	decimalValAsInt := math.Ceil(float * 100) // float number converted to integer => will be 85
	switch {
	case decimalValAsInt > 75:
		{
			roundedValue := math.Ceil(f) // if fractional part is bigger than 75, f will be rounded up to its greater integer. 22.85 => 23
			return roundedValue          // will be return this
		}
	case decimalValAsInt > 50:
		{
			roundedValue := i + 0.75 // if fractional part is bigger than 50, we will sum 0.75 and i which is the integer part of the number ( we founded -> i=22) 22.75
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
		return f // if fractional part is equal to 0, sayı yuvarlanmayacak
	}
}

func AlphabetSoup(s string) string {
	a := strings.Split(s, "")    // we separate all the letters
	sort.Strings(a)              // and we sort them (order by asc)
	joins := strings.Join(a, "") // and we join the letters which we separated
	return joins
}

func StringMask(s string, n uint) string {
	a := strings.Split(s, "") // we seperate all the letters in s
	length := len(a)          // length a
	nint := int(n)            // n converted to int for use on if loop

	if nint >= length { // if n is bigger than my array length or equal to my array length
		for i := 0; i < length; i++ {
			a[i] = "*" // all array element will be *
		}
	} else {
		for i := nint; i < length; i++ { // if n is less than my array length, After n elements there will be stars
			a[i] = "*"
		}
	}
	w := strings.Join(a, "") // and join the letters
	return w
}

func WordSplit(arr [2]string) string {
	find := []string{}                    // a slice to discard any elements we find
	myArr := arr[0]                       // array in which we will search for words
	words := arr[1]                       // words
	wordsArr := strings.Split(words, ",") // words seperated with "," and it been array.
	for _, v := range wordsArr {
		if strings.Contains(myArr, v) { // if myArr contains element words
			el := []string{v}          // i am assigning the element to the array for give as the first parameter inside append
			find = append(el, find...) // The last element I found (in the first index) is merged with the elements I found earlier.
			// "range" return sequentially in "words", but it does not return sequentially in the string(MyArray). If we're not do this, whatever the elements are, orders by "words".
			myArr = strings.ReplaceAll(myArr, v, "") //I deleted the element I found from the array I was searching for, because later I will check if there is an element in this array that I did not search for
		}
	}
	if len(find) == 0 || len(myArr) > 0 { // if MyArray contains elements that are not element of words or if the element count is 0 => return not possible
		return "not possible"
	} else {
		return strings.Join(find, ",") // otherwise it will combine and return the words I find.
	}
}

// !!! I found the answer to this question on the internet, i understand that and made some changes
func VariadicSet(i ...interface{}) []interface{} {
	tempMap := make(map[interface{}]bool) // created map, types => [key:interface, value:bool]
	var uniqueValues []interface{}        //the array where we will add the previously not added elements one by one.
	for k := 0; k < len(i); k++ {
		if _, val := tempMap[i[k]]; !val { //if element does not exist in map =>
			tempMap[i[k]] = true                      // key will be add with "true" value
			uniqueValues = append(uniqueValues, i[k]) //and will be add to uniqueValues
		}
	}
	return uniqueValues //returned
}
