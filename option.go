package jc

import "io"

type Option func(*jc)

func Indent(i string) Option {
	return func(j *jc) {
		j.SetIndent(i)
	}
}

func (j *jc) SetIndent(i string) {
	j.indent = i
}

func Writer(w io.Writer) Option {
	return func(j *jc) {
		j.SetWriter(w)
	}
}

func (j *jc) SetWriter(w io.Writer) {
	j.writer = w
}
