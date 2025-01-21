package main

import (
	"fmt"
	"jvm-go/classpath"
	"jvm-go/cmd"
	"strings"
)

func main() {
	myCmd := cmd.ParsecCmd()
	if myCmd.VersionFlag {
		fmt.Println("Version: 0.0.1")
	} else if myCmd.HelpFlag || myCmd.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(myCmd)
	}
}

func startJVM(cmd *cmd.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v \n", cmd.CpOption, cmd.Class, cmd.Args)
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.Class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
