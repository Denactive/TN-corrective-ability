package main

import (
	//"bytes"
	"fmt"
	"html/template"
	//"log"
	"math/bits"
	"net/http"
	"strconv"
)

const n = 15
const k = 11
const inVector = "00001010011"

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func getBinaryLength(digit int) int {
	bitsNum := 0
	for digitCopy := digit; digitCopy / 2 != 0; digitCopy /= 2 {
		bitsNum++
	}
	bitsNum++
	return bitsNum
}

func powBinary(n uint64) uint64 {
	res := uint64(1)
	for i := uint64(1); i <= n; i++ {
		res <<= 1
	}
	return res
}

//func shiftLeftPositiveBit(digit, pos int) int {
//	bitsNum := getBinaryLength(digit)
//
//	if pos < 0 || pos > bitsNum {
//		return digit
//	}
//
//	return digit + powBinary(pos)
//}
//
//func bitsToString(digit int) string {
//	var res string
//	for i := powBinary(getBinaryLength(digit) - 1); digit != 0 && i > 0; i /= 2 {
//		res += fmt.Sprintf("%d", digit / i)
//		digit %= i
//	}
//	return res
//}
//
//func bitsToBytes(digit int) []byte {
//	var res []byte
//	for i := powBinary(getBinaryLength(digit) - 1); digit != 0 && i > 0; i /= 2 {
//		res = append(res, byte(digit / i))
//		digit %= i
//	}
//	return res
//}


func factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func getAllErrorsByClasses(n uint64) [][]uint64 {
	errorClasses := make([][]uint64, n)
	for i := uint64(1); i <= n; i++ {
		size := factorial(n) / factorial(n-i) / factorial(i)
		errorClasses[i-1] = make([]uint64, 0, size)
	}

	for i := uint64(1); i < powBinary(n); i++ {
		class := bits.OnesCount64(i) - 1
		errorClasses[class] = append(errorClasses[class], i)
	}
	return  errorClasses
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home page</h1>")

	errorClasses := getAllErrorsByClasses(n)
	//for i := uint64(1); i <= n; i++ {
	//	fmt.Fprintf(w,"class %d: %d\n", i, len(errorClasses[i - 1]))
	//}

	//t := template.Must(template.New("").Parse(`<table>{{range .}}<tr><td>{{.}}</td></tr>{{end}}</table>`))
	//var body bytes.Buffer
	//if err := t.Execute(w, errorClasses); err != nil {
	//	log.Fatal(err)
	//}
	//title := "User Info Page"
	//p := &Page{Title: title, Body: body.Bytes()}
	//renderTemplate(w, "view", p)

	// template & error
	tmpl, _ := template.ParseFiles("./home_page.html")
	tmpl.Execute(w, errorClasses[0])
	//tmpl.Execute(w)
}

func main() {

	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)

	//base := append(make([]byte, n-m), bitsToBytes(powBinary(m) - 1)...)
	//fmt.Println(iteratePermutations(base))
	//fmt.Printf("\t%v\n", base)
}
