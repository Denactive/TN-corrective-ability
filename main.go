package main

import (
	"fmt"
	"html/template"
	"math/bits"
	"net/http"
)

const n = 15
const k = 11
const inVector = 83  // 000.0101.0011b
const genPolynomial = 19  // 10011b

func powBinary(n uint64) uint64 {
	res := uint64(1)
	for i := uint64(1); i <= n; i++ {
		res <<= 1
	}
	return res
}

func getBinaryLength(digit uint64) uint64 {
	bitsNum := uint64(0)
	for ; digit / 2 != 0; digit /= 2 {
		bitsNum++
	}
	bitsNum++
	return bitsNum
}

func IntToBytes(digit uint64) []byte {
	var res []byte
	for i := powBinary(getBinaryLength(digit) - 1); i > 0; i /= 2 {
		res = append(res, byte(digit / i))
		digit %= i
	}
	return res
}
func factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func OperationO(a, b uint64) (uint64, uint64) {
	if a < b {
		return 0, a
	}

	var integer uint64
	aBytes := IntToBytes(a)
	bLen := getBinaryLength(b)
	var cur uint64

	aBytesPos := uint64(0)
	for ; aBytesPos < bLen; aBytesPos++ {
		cur <<= 1
		cur += uint64(aBytes[aBytesPos])
	}

	for ; aBytesPos <= uint64(len(aBytes)); aBytesPos++ {
		firstBitInCur := cur / powBinary(bLen - 1)
		integer <<= 1
		integer += firstBitInCur

		if firstBitInCur == 1 {
			cur ^= b
		}
		if aBytesPos == uint64(len(aBytes)) {
			break
		}

		cur <<= 1
		cur += uint64(aBytes[aBytesPos])
	}

	return integer, cur
}

func getAllErrorsByClasses(n uint64) [][]uint64 {
	errorClasses := make([][]uint64, n + 1)
	for i := uint64(1); i <= n; i++ {
		size := factorial(n) / factorial(n-i) / factorial(i)
		errorClasses[i] = make([]uint64, 0, size)
	}

	for i := uint64(1); i < powBinary(n); i++ {
		class := bits.OnesCount64(i)
		errorClasses[class] = append(errorClasses[class], i)
	}
	return  errorClasses
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	errorClasses := getAllErrorsByClasses(n)
	errorClassesBinStr := make([][]string, n + 1)
	for i := uint64(1); i <= n; i++ {
		errorClassesBinStr[i] = make([]string, len(errorClasses[i]))
		for j := uint64(0); j < uint64(len(errorClasses[i])); j++ {
			errorClassesBinStr[i][j] = fmt.Sprintf("%b",errorClasses[i][j])
		}
	}

	tmpl, _ := template.ParseFiles("./templates/errors.html")
	tmpl.Execute(w, errorClassesBinStr)
}

func getSymptomErrorTable(n, genPolynomial uint64) map[uint64] uint64 {
	errorMap := make(map[uint64] uint64, powBinary(n))
	for i := uint64(1); i < powBinary(n); i++ {
		_, symptom := OperationO(i, genPolynomial)
		errorMap[symptom] = i
	}
	return errorMap
}

func getSymptomErrorTableString(n, genPolynomial uint64) map[string] string {
	errorMap := make(map[string] string, powBinary(n))
	for i := uint64(1); i < powBinary(n); i++ {
		_, symptom := OperationO(i, genPolynomial)
		errorMap[fmt.Sprintf("%b",symptom)] = fmt.Sprintf("%b",i)
	}
	//fmt.Println(errorMap)
	return errorMap
}

func SymptomPage(w http.ResponseWriter, r *http.Request) {
	errorMap := getSymptomErrorTableString(n, genPolynomial)

	tmpl, _ := template.ParseFiles("./templates/symptoms.html")
	tmpl.Execute(w, errorMap)
}

func getSymptomArray(n, genPolynomial uint64) []string {
	symptomArray := make([]string, powBinary(n) + 1)
	for i := uint64(1); i < powBinary(n); i++ {
		_, symptom := OperationO(i, genPolynomial)
		symptomArray[i] = fmt.Sprintf("%b", symptom)
	}
	return symptomArray
}

func SymptomPageArray(w http.ResponseWriter, r *http.Request) {
	errorMap := getSymptomArray(n, genPolynomial)

	tmpl, _ := template.ParseFiles("./templates/symptomsArray.html")
	tmpl.Execute(w, errorMap)
}

func main() {

	http.HandleFunc("/", ErrorPage)
	http.HandleFunc("/symptoms/", SymptomPage)
	http.HandleFunc("/symptoms/array", SymptomPageArray)
	http.ListenAndServe(":8080", nil)

}
