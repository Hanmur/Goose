package convert

//类型转换

import "strconv"

//StrTo string转换类型
type StrTo string

//String 将StrTo转换为String类型
func (s StrTo) String() string {
	return string(s)
}

//Int 将StrTo转换为Int类型
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

//MustInt 将StrTo强制转换为Int类型
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

//UInt32 将StrTo转换为Uint32类型
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

//MustUInt32 将StrTo强制转换为Uint32类型
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
