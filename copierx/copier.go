package copierx

import (
	"github.com/jinzhu/copier"
)

func Copy(to, from interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		panic(err)
	}
}
