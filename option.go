package jc

import (
	"io"

	"github.com/fatih/color"
)

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

func KeyColor(c *color.Color) Option {
	return func(j *jc) {
		j.SetKeyColor(c)
	}
}

func (j *jc) SetKeyColor(c *color.Color) {
	j.keyColor = c
}

func NumberColor(c *color.Color) Option {
	return func(j *jc) {
		j.SetNumberColor(c)
	}
}

func (j *jc) SetNumberColor(c *color.Color) {
	j.numberColor = c
}

func StringColor(c *color.Color) Option {
	return func(j *jc) {
		j.SetStringColor(c)
	}
}

func (j *jc) SetStringColor(c *color.Color) {
	j.stringColor = c
}

func BoolColor(c *color.Color) Option {
	return func(j *jc) {
		j.SetBoolColor(c)
	}
}

func (j *jc) SetBoolColor(c *color.Color) {
	j.boolColor = c
}
