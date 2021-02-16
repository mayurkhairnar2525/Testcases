package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main (){
	fmt.Println(Converter("This is some ####sample string !@#4 you"))
}
func Converter(input string)(string,error){
	if len (input)==7 {
return "",fmt.Errorf("Input level is 7 which is not forbidden")
	}
splitter := func (c rune )bool {
	return! unicode.IsLetter(c)&&!unicode.IsNumber(c)
}
splits:= strings.FieldsFunc(input,splitter)
out:= strings.Join(splits,"")
return out,nil
}

