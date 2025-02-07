package main

import (
	"fmt"
	"jvm-go/classfile"
	"jvm-go/classpath"
	"jvm-go/cmd"
	"jvm-go/rtda"
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
		//startJVMTest(myCmd)
	}

}

func myTestCmd() *cmd.Cmd {
	return &cmd.Cmd{
		XjreOption: "/Library/Java/JavaVirtualMachines/jdk-1.8.jdk/Contents/Home/jre",
		Class:      "GuessTest",
		Args:       []string{},
	}
}

func startJVM(cmd *cmd.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v \n", cmd.CpOption, cmd.Class, cmd.Args)
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpreter(mainMethod)
	} else {
		fmt.Println("Main method not found in class %s\n", cmd.Class)
	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
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

func startJVMTest(cmd *cmd.Cmd) {
	testLocalVarAndOperandStack()
}

// testLocalVarAndOperandStack 测试局部变量表和操作数栈
func testLocalVarAndOperandStack() {

	frame := rtda.NewFrame(rtda.NewThread(), 100, 100)
	testLocalVar(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVar(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
