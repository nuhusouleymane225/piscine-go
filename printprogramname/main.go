package main

import (
	"fmt"
	"os"
	"strings"
	)

func main(){
	arguments:=os.Args
	arr:=strings.Split(arguments[0],"/")
	fmt.Printf("%v\n",arr[len(arr)-1])
}
