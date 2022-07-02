package main

import "fmt"

type NumError struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat, ParseComplex)
	Num  string // the input
	Err  error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
}

func (e *NumError) Error() string {
	return "strconv." + e.Func + ": " + "parsing " + e.Num + ": " + e.Err.Error()
}

func main() {
	var s string = "123"

	n := 0
	for i, ch := range []byte(s) {
		fmt.Printf("%d > ch=%v, value=%v\n", i, ch, ch-'0')
		ch -= '0'
		if ch > 9 {
			fmt.Println("Error")
			return
		}

		n = 10*n + int(ch)
	}
	fmt.Println(n)
}
