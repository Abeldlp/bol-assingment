package custom_type

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type IntArray []int

func (a *IntArray) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("unexpected type for IntArray: %T", value)
	}

	vals := strings.Split(str[1:len(str)-1], ",")
	ints := []int{}
	for _, v := range vals {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		ints = append(ints, i)
	}
	*a = ints
	return nil
}

func (a IntArray) Value() (driver.Value, error) {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, val := range a {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(val))
	}
	buf.WriteString("}")

	return buf.String(), nil
}
