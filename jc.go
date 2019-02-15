package jc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type jc struct {
	indent string
	writer io.Writer
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
	if err := j.writeln(); err != nil {
		return err
	}
	return nil
}

func (j *jc) indentation(depth int) (err error) {
	for i := 0; i < depth; i++ {
		err = j.write(j.indent)
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
		err = j.writef("%v", val)
	case string:
		err = j.writef("%q", val)
	case bool:
		err = j.writef("%v", val)
	case map[string]interface{}:
		err = j.writeln("{")
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

			err = j.writef(`%q: `, k)
			if err != nil {
				return err
			}

			err = j.parse(vv, depth+1)
			if err != nil {
				return err
			}
			if i < count-1 {
				err = j.write(",")
				if err != nil {
					return err
				}
			}
			err = j.writeln()
			if err != nil {
				return err
			}
			i++
		}

		err = j.indentation(depth)
		if err != nil {
			return err
		}

		err = j.write("}")
	case []interface{}:
		err = j.writeln("[")
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
				err = j.write(",")
				if err != nil {
					return err
				}
			}

			err = j.writeln()
			if err != nil {
				return err
			}
		}
		err = j.indentation(depth)
		if err != nil {
			return err
		}

		err = j.write("]")
	default:
		return fmt.Errorf("unknown type: %v", val)
	}

	if err != nil {
		return err
	}
	return nil
}

func (j *jc) _write(s string) error {
	if _, err := j.writer.Write([]byte(s)); err != nil {
		return err
	}
	return nil
}

func (j *jc) write(a ...interface{}) error {
	return j._write(fmt.Sprint(a...))
}

func (j *jc) writef(f string, args ...interface{}) error {
	return j._write(fmt.Sprintf(f, args...))
}

func (j *jc) writeln(a ...interface{}) error {
	return j._write(fmt.Sprintln(a...))
}
