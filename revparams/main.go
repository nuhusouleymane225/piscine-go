package main

import (
	"fmt"
	"os"
	)

func main(){
	args:=os.Args
	if len(args)> 1{
		for i:=len(args)-1;i>0;i--{
			fmt.Printf("%v\n",args[i])
		}
		
	}
}
