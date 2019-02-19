package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/matsune/jc"
	"golang.org/x/crypto/ssh/terminal"
)

type conf struct {
	Key    string
	Number string
	String string
	Bool   string
	Null   string
}

var version = "1.0"

func main() {
	for _, v := range os.Args {
		if v == "-v" || v == "--version" {
			fmt.Printf("jc version %s\n", version)
			os.Exit(0)
		}
	}
	if terminal.IsTerminal(0) {
		return
	}

	// read from pipe
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	usr, err := user.Current()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	c := conf{
		Key:    "",
		Number: "34",
		String: "33",
		Bool:   "31",
		Null:   "36",
	}

	confPath := filepath.Join(usr.HomeDir, ".jc.conf")
	if _, err = os.Stat(confPath); !os.IsNotExist(err) {
		_, err = toml.DecodeFile(confPath, &c)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

	keys, err := splitAttributes(c.Key)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	keyColor := color.New(keys...)

	nums, err := splitAttributes(c.Number)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	numColor := color.New(nums...)

	strings, err := splitAttributes(c.String)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	strColor := color.New(strings...)

	bools, err := splitAttributes(c.Bool)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	boolColor := color.New(bools...)

	nulls, err := splitAttributes(c.Null)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	nullColor := color.New(nulls...)

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

func splitAttributes(str string) ([]color.Attribute, error) {
	if len(str) == 0 {
		return []color.Attribute{}, nil
	}
	numStrs := strings.Split(str, ",")
	attrs := make([]color.Attribute, len(numStrs))
	for i, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		attrs[i] = color.Attribute(num)
	}
	return attrs, nil
}
