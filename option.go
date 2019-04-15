package jc

import (
	"io"

	"github.com/fatih/color"
)

type Option func(*JC)

func Indent(i string) Option {
	return func(j *JC) {
		j.SetIndent(i)
	}
}

func (j *JC) SetIndent(i string) {
	j.indent = i
}

func Writer(w io.Writer) Option {
	return func(j *JC) {
		j.SetWriter(w)
	}
}

func (j *JC) SetWriter(w io.Writer) {
	j.writer = w
}

func KeyColor(c *color.Color) Option {
	return func(j *JC) {
		j.SetKeyColor(c)
	}
}

func (j *JC) SetKeyColor(c *color.Color) {
	j.keyColor = c
}

func NumberColor(c *color.Color) Option {
	return func(j *JC) {
		j.SetNumberColor(c)
	}
}

func (j *JC) SetNumberColor(c *color.Color) {
	j.numColor = c
}

func StringColor(c *color.Color) Option {
	return func(j *JC) {
		j.SetStringColor(c)
	}
}

func (j *JC) SetStringColor(c *color.Color) {
	j.strColor = c
}

func BoolColor(c *color.Color) Option {
	return func(j *JC) {
		j.SetBoolColor(c)
	}
}

func (j *JC) SetBoolColor(c *color.Color) {
	j.boolColor = c
}

func NullColor(c *color.Color) Option {
	return func(j *JC) {
		j.SetNullColor(c)
	}
}

func (j *JC) SetNullColor(c *color.Color) {
	j.nullColor = c
}
