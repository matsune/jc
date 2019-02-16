package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/matsune/jc"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if !terminal.IsTerminal(0) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		keyColor := color.New(color.FgCyan)
		numColor := color.New(color.BgBlue, color.Bold)
		strColor := color.New(color.FgWhite, color.BgBlack)
		boolColor := color.New(color.FgHiBlue)
		nullColor := color.New(color.FgHiBlue, color.BgHiRed)
		j := jc.New(
			jc.KeyColor(keyColor),
			jc.NumberColor(numColor),
			jc.StringColor(strColor),
			jc.BoolColor(boolColor),
			jc.NullColor(nullColor),
		)
		if err := j.Colorize(string(b)); err != nil {
			panic(err)
		}
	}
}
