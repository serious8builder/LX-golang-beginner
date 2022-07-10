package main

import "fmt"

type Report interface {
	Report() string
}

type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct{}

func (s *EmailSender) Send(r *Report) {
	fmt.Println("Send Email...")
	fmt.Println((*r).Report())
}

type Point struct {
	X, Y int64
}

//func (p *Point) String() string {
//	return fmt.Sprintf("Point(X=%d, Y=%d)", p.X, p.Y)
//}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	fmt.Printf("Type: %T, Value: %v\n", &p, &p)
}
