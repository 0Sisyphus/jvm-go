package main

import (
	"fmt"
	"jvm-go/classfile"
	"jvm-go/classpath"
	"jvm-go/cmd"
	"strings"
)

func main() {
	//myCmd := cmd.ParsecCmd()
	myCmd := myTestCmd()
	if myCmd.VersionFlag {
		fmt.Println("Version: 0.0.1")
	} else if myCmd.HelpFlag || myCmd.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(myCmd)
	}

}

func myTestCmd() *cmd.Cmd {
	return &cmd.Cmd{
		XjreOption: "/Library/Java/JavaVirtualMachines/jdk-1.8.jdk/Contents/Home/jre",
		Class:      "java.lang.String",
		Args:       []string{},
	}
}

func startJVM(cmd *cmd.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v \n", cmd.CpOption, cmd.Class, cmd.Args)
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.Class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count:%v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags:0x%x\n", cf.AccessFlags())
	fmt.Printf("this class:%s\n", cf.ClassName())
	fmt.Printf("super class:%s\n", cf.SuperClassName())
	fmt.Printf("interfaces count:%v\n", len(cf.InterfaceNames()))
	fmt.Printf("fields count:%v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("field name:%s \n", f.Name())
	}
	fmt.Printf("method count:%v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("method name:%s descriptor: %s\n", m.Name(), m.Descriptor())
	}

}
