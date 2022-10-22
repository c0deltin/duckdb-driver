package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
)

type StringArray []string

// Scan implements the sql.Scanner interface.
func (a *StringArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []interface{}:
		if len(src) == 0 {
			return nil
		}
		var strArr = make(StringArray, len(src))
		for i, x := range src {
			strArr[i] = x.(string)
		}
		*a = strArr
		return nil
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("cannot convert %T to StringArray", src)
}

// Value implements the driver.Valuer interface.
func (a StringArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+3*n)
		b[0] = '{'

		b = appendArrayQuotedBytes(b, []byte(a[0]))
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedBytes(b, []byte(a[i]))
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

func appendArrayQuotedBytes(b, v []byte) []byte {
	b = append(b, '"')
	for {
		i := bytes.IndexAny(v, `"\`)
		if i < 0 {
			b = append(b, v...)
			break
		}
		if i > 0 {
			b = append(b, v[:i]...)
		}
		b = append(b, '\\', v[i])
		v = v[i+1:]
	}
	return append(b, '"')
}

// Int32Array represents a one-dimensional array of the PostgreSQL integer types.
type Int32Array []int32

// Scan implements the sql.Scanner interface.
func (a *Int32Array) Scan(src interface{}) error {
	switch src := src.(type) {
	case []interface{}:
		if len(src) == 0 {
			return nil
		}
		var intArr = make(Int32Array, len(src))
		for i, x := range src {
			intArr[i] = x.(int32)
		}
		*a = intArr
		return nil
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("cannot convert %T to Int32Array", src)
}

// Value implements the driver.Valuer interface.
func (a Int32Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendInt(b, int64(a[0]), 10)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendInt(b, int64(a[i]), 10)
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}
