package jc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type jc struct {
	indent string
	depth  int
	io.Writer
}

func New(opts ...Option) *jc {
	j := &jc{
		indent: "\t",
		depth:  0,
		Writer: os.Stdout,
	}
	for _, opt := range opts {
		opt(j)
	}
	return j
}

func (j *jc) reset() {
	j.depth = 0
}

func (j *jc) Colorize(str string) error {
	var v interface{}
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return err
	}
	j.reset()
	if err := j.parse(v); err != nil {
		return err
	}
	if err := j.writeln(); err != nil {
		return err
	}
	return nil
}

func (j *jc) indentation() (err error) {
	for i := 0; i < j.depth; i++ {
		err = j.write(j.indent)
		if err != nil {
			return err
		}
	}
	return
}

func (j *jc) parse(v interface{}) (err error) {
	switch val := v.(type) {
	case map[string]interface{}:
		err = j.writeln("{")
		if err != nil {
			return err
		}
		j.depth++

		count := len(val)
		i := 0

		for k, vv := range val {
			err = j.indentation()
			if err != nil {
				return err
			}

			err = j.writef(`"%s": `, k)
			if err != nil {
				return err
			}

			err = j.parse(vv)
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

		j.depth--
		err = j.indentation()
		if err != nil {
			return err
		}

		err = j.write("}")
	case float64:
		err = j.writef("%v", val)
	case string:
		err = j.writef("\"%s\"", val)
	case bool:
		err = j.writef("%v", val)
	case []interface{}:
		err = j.writeln("[")
		if err != nil {
			return err
		}
		j.depth++
		for i, vv := range val {
			err = j.indentation()
			if err != nil {
				return err
			}

			err = j.parse(vv)
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
		j.depth--
		err = j.indentation()
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
	return
}

func (j *jc) _write(s string) error {
	if _, err := j.Write([]byte(s)); err != nil {
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
