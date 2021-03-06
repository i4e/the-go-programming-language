package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "true")
		} else {
			fmt.Fprintf(buf, "false")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Complex64, reflect.Complex128:
		z := v.Complex()
		fmt.Fprintf(buf, "#C(%v %v)", real(z), imag(z))

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('{')
		indent++
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				buf.Write(bytes.Repeat([]byte{' '}, indent))
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		indent++
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				buf.Write(bytes.Repeat([]byte{' '}, indent))
			}
			n, _ := fmt.Fprintf(buf, "{%q: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+n); err != nil {
				return err
			}
			buf.WriteByte('}')
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte('\n')
				buf.Write(bytes.Repeat([]byte{' '}, indent))
			}
			buf.WriteByte('{')
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.WriteByte('}')
		}
		buf.WriteByte('}')

	case reflect.Interface:
		buf.WriteByte('{')
		indent++
		n, _ := buf.WriteString(strconv.Quote(v.Elem().Type().String()))
		indent += n
		buf.WriteString(": ")
		indent++
		encode(buf, v.Elem(), indent)
		buf.WriteByte('}')

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %q", v.Type())
	}
	return nil
}
