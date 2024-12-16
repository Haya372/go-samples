package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	args StringList = make(StringList, 0)
)

type StringList []string

func (l *StringList) Set(s string) error {
	*l = append(*l, s)
	return nil
}

func (l *StringList) String() string {
	return "[" + strings.Join([]string(*l), ",") + "]"
}

func init() {
	flag.Var(&args, "args", "list of args")

	flag.Parse()
}

func main() {
	fmt.Printf("args=%v\n", args)
}
