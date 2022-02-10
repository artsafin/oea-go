package excel

import "fmt"

type floater interface {
	AsFloat64() float64
}

type numberCell struct {
	value   float64
	comment string
}

func (r numberCell) AsFloat64() float64 {
	return r.value
}
func (r numberCell) Comment() string {
	return r.comment
}

type cell struct {
	value   interface{}
	comment string
}

func (r cell) String() string {
	return fmt.Sprint(r.value)
}

func (r cell) Comment() string {
	return r.comment
}

//func (r cell) StyleID() int {
//	return r.styleID
//}
