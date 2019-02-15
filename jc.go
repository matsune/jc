package jc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

type jc struct {
	indent      string
	writer      io.Writer
	keyColor    *color.Color
	numberColor *color.Color
	stringColor *color.Color
	boolColor   *color.Color
}

func New(opts ...Option) *jc {
	j := &jc{
		indent: "\t",
		writer: os.Stdout,
	}
	for _, opt := range opts {
		opt(j)
	}
	return j
}

func (j *jc) Colorize(str string) error {
	var v interface{}
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return err
	}
	if err := j.parse(v, 0); err != nil {
		return err
	}
	if err := j.writeln(nil); err != nil {
		return err
	}
	return nil
}

func (j *jc) indentation(depth int) (err error) {
	for i := 0; i < depth; i++ {
		err = j.write(nil, j.indent)
		if err != nil {
			return err
		}
	}
	return
}

func (j *jc) parse(v interface{}, depth int) error {
	var err error
	switch val := v.(type) {
	case float64:
		err = j.writef(j.numberColor, "%v", val)
	case string:
		err = j.writef(j.stringColor, "%q", val)
	case bool:
		err = j.writef(j.boolColor, "%v", val)
	case map[string]interface{}:
		err = j.writeln(nil, "{")
		if err != nil {
			return err
		}

		count := len(val)
		i := 0

		for k, vv := range val {
			err = j.indentation(depth + 1)
			if err != nil {
				return err
			}

			j.writef(j.keyColor, "%q", k)
			j.write(nil, ": ")

			if err != nil {
				return err
			}

			err = j.parse(vv, depth+1)
			if err != nil {
				return err
			}
			if i < count-1 {
				err = j.write(nil, ",")
				if err != nil {
					return err
				}
			}
			err = j.writeln(nil)
			if err != nil {
				return err
			}
			i++
		}

		err = j.indentation(depth)
		if err != nil {
			return err
		}

		err = j.write(nil, "}")
	case []interface{}:
		err = j.writeln(nil, "[")
		if err != nil {
			return err
		}

		for i, vv := range val {
			err = j.indentation(depth + 1)
			if err != nil {
				return err
			}

			err = j.parse(vv, depth+1)
			if err != nil {
				return err
			}

			if i < len(val)-1 {
				err = j.write(nil, ",")
				if err != nil {
					return err
				}
			}

			err = j.writeln(nil)
			if err != nil {
				return err
			}
		}
		err = j.indentation(depth)
		if err != nil {
			return err
		}

		err = j.write(nil, "]")
	default:
		return fmt.Errorf("unknown type: %v", val)
	}

	if err != nil {
		return err
	}
	return nil
}

func (j *jc) write(c *color.Color, a ...interface{}) error {
	var err error
	if c != nil {
		_, err = c.Fprint(j.writer, a...)
	} else {
		_, err = fmt.Fprint(j.writer, a...)
	}
	return err
}

func (j *jc) writef(c *color.Color, f string, a ...interface{}) error {
	var err error
	if c != nil {
		_, err = c.Fprintf(j.writer, f, a...)
	} else {
		_, err = fmt.Fprintf(j.writer, f, a...)
	}
	return err
}

func (j *jc) writeln(c *color.Color, a ...interface{}) error {
	var err error
	if c != nil {
		_, err = c.Fprintln(j.writer, a...)
	} else {
		_, err = fmt.Fprintln(j.writer, a...)
	}
	return err
}
