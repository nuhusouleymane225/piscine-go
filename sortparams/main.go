package main

import (
	"fmt"
	"os"
	)
func main(){
	args:=os.Args[1:]
	args=sortascii(args)
	for i:=range args{
		fmt.Println(args[i])
	}	
}

func sortascii(s []string) []string{
	ech:=""
	for i:=0;i<len(s);i++{
		for j:=i+1;j<len(s);j++{
			if []byte(s[j])[0]<[]byte(s[i])[0]{
				ech=s[i]
				s[i]=s[j]
				s[j]=ech
			}
		}
	}
	return s
}
