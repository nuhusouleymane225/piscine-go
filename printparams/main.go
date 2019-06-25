package main

import (
	"fmt"
	"os"
	)

func main(){
	args:=os.Args
	if len(args)> 1{
		for i:=1;i<len(args);i++{
			fmt.Printf("%v\n",args[i])
		}
		
	}
}
