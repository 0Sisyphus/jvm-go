package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	XjreOption  string
	Class       string
	Args        []string
}

func ParsecCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	fmt.Printf("ParsecCmd classpath:%s class:%s XjreOption:%s args:%v \n", cmd.CpOption, cmd.Class, cmd.XjreOption, cmd.Args)
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] Class [Args...]\n", os.Args[0])
}
