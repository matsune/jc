package jc

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	json "github.com/matsune/go-json"
)

type jc struct {
	indent    string
	writer    io.Writer
	keyColor  *color.Color
	numColor  *color.Color
	strColor  *color.Color
	boolColor *color.Color
	nullColor *color.Color
}

func New(opts ...Option) *jc {
	j := &jc{
		indent:    "\t",
		writer:    os.Stdout,
		numColor:  color.New(color.Attribute(34)),
		strColor:  color.New(color.Attribute(33)),
		boolColor: color.New(color.Attribute(31)),
		nullColor: color.New(color.Attribute(36)),
	}
	for _, opt := range opts {
		opt(j)
	}
	return j
}

func (j *jc) Colorize(str string) error {
	v, err := json.Parse(str)
	if err != nil {
		return err
	}
	j.walk(v, 0)
	return nil
}

func (j *jc) indentation(depth int) error {
	var err error
	for i := 0; i < depth; i++ {
		err = j.write(nil, j.indent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (j *jc) walk(v json.Value, nest int) error {
	var err error
	switch v := v.(type) {
	case *json.ObjectValue:
		if err = j.writeln(nil, "{"); err != nil {
			return err
		}

		nest++
		for i, kv := range v.KeyValues {
			if err = j.indentation(nest); err != nil {
				return err
			}

			if err = j.writef(j.keyColor, "%q: ", kv.Key); err != nil {
				return err
			}

			if err = j.walk(kv.Value, nest); err != nil {
				return err
			}

			if i < len(v.KeyValues)-1 {
				err = j.writeln(nil, ",")
			} else {
				err = j.writeln(nil)
			}
			if err != nil {
				return err
			}
		}
		nest--

		if err = j.indentation(nest); err != nil {
			return err
		}
		err = j.write(nil, "}")
	case *json.ArrayValue:
		j.writeln(nil, "[")
		nest++
		for i, vv := range v.Values {
			if err = j.indentation(nest); err != nil {
				return err
			}
			if err = j.walk(vv, nest); err != nil {
				return err
			}
			if i < len(v.Values)-1 {
				if err = j.write(nil, ","); err != nil {
					return err
				}
			}
			if err = j.writeln(nil); err != nil {
				return err
			}
		}
		nest--
		if err = j.indentation(nest); err != nil {
			return err
		}
		err = j.write(nil, "]")
	case *json.IntValue:
		err = j.write(j.numColor, v)
	case *json.FloatValue:
		err = j.write(j.numColor, v)
	case *json.BoolValue:
		err = j.write(j.boolColor, v)
	case *json.NullValue:
		err = j.write(j.nullColor, v)
	case *json.StringValue:
		err = j.write(j.strColor, v)
	}
	return err
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
