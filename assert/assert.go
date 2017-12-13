package assert

import (
	"database/sql"
	"reflect"

	"fmt"
)

func Nil(x interface{}) {
	if x != nil {
		panic(fmt.Sprintf("expected nil; got %v", x))
	}
}

func NotNil(x interface{}) {
	if x == nil {
		panic(x)
	}
}

func Equal(x, y interface{}) {
	if x == y {
		return
	}
	yAsXType := reflect.ValueOf(y).Convert(reflect.TypeOf(x)).Interface()
	if !reflect.DeepEqual(x, yAsXType) {
		panic(fmt.Sprintf("%s != %s", x, y))
	}
}

func StrictlyEqual(x, y interface{}) {
	if x != y {
		panic(fmt.Sprintf("%s != %s"))
	}
}

func OneRowAffected(r sql.Result) {
	count, err := r.RowsAffected()
	Nil(err)
	if count != 1 {
		panic(count)
	}
}

func True(b bool) {
	if !b {
		panic(b)
	}
}
