package cell

import (
	"github.com/zhujf1989/gotable/util"
)

type Data struct {
	value  string
	length int
}

func CreateData(value string) *Data {
	d := new(Data)
	d.value = value
	d.length = util.Length(value)
	return d
}

func CreateEmptyData() *Data {
	return CreateData("")
}

func (d *Data) String() string {
	return d.value
}

func (d *Data) Length() int {
	return d.length
}

func (d *Data) Original() string {
	return d.String()
}
