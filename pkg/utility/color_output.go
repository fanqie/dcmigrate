package utility

import (
	"fmt"
	"github.com/fatih/color"
)

func ErrPrintf(str string, values ...interface{}) {
	_, err := color.New(color.FgHiRed).Printf(fmt.Sprintf("[Error]%s\n", str), values...)
	if err != nil {
		panic(err)
	}
}
func ErrPrint(str string) {
	_, err := color.New(color.FgHiRed).Print(fmt.Sprintf("[Error]%s\n", str))
	if err != nil {
		panic(err)
	}
}
func SuccessPrintf(str string, values ...interface{}) {
	_, err := color.New(color.FgHiGreen).Printf(fmt.Sprintf("[Success]%s\n", str), values...)
	if err != nil {
		panic(err)
	}
}
func SuccessPrint(str string) {
	_, err := color.New(color.FgHiGreen).Print(fmt.Sprintf("[Success]%s\n", str))
	if err != nil {
		panic(err)
	}
}
func InfoPrintf(str string, values ...interface{}) {
	_, err := color.New(color.FgHiBlue).Printf(fmt.Sprintf("[Info]%s\n", str), values...)
	if err != nil {
		panic(err)
	}
}
func InfoPrint(str string) {
	_, err := color.New(color.FgHiBlue).Printf(fmt.Sprintf("[Info]%s\n", str))
	if err != nil {
		panic(err)
	}
}
func WarningPrintf(str string, values ...interface{}) {
	_, err := color.New(color.FgHiYellow).Printf(fmt.Sprintf("[Warning]%s\n", str), values...)
	if err != nil {
		panic(err)
	}
}
func WarningPrint(str string) {
	_, err := color.New(color.FgHiYellow).Printf(fmt.Sprintf("[Warning]%s\n", str))
	if err != nil {
		panic(err)
	}
}
