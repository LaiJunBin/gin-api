package converts

import "strconv"

type StrTo string

func (str StrTo) String() string {
	return string(str)
}

func (str StrTo) Int() (int, error){
	value, err := strconv.Atoi(str.String())
	return value, err
}

func (str StrTo) MustInt() int {
	value, _ := str.Int()
	return value
}

func (str StrTo) UInt32() (uint32, error) {
	value, err := strconv.Atoi(str.String())
	return uint32(value), err
}

func (str StrTo) MustUInt32() uint32 {
	value, _ := str.UInt32()
	return value
}

func (str StrTo) UInt() (uint, error) {
	value, err := strconv.Atoi(str.String())
	return uint(value), err
}

func (str StrTo) MustUInt() uint {
	value, _ := str.UInt()
	return value
}