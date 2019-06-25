package piscine

import "strconv"
import "fmt"

func ConvertBase(nbr, baseFrom, baseTo string) string{
	return PrintNbrBase1(AtoiBase(nbr,baseFrom),baseTo)
}

func PrintNbrBase1(n int,base string) string{//convertir un int dans une base et retourne le string de celui-ci
	str:=""
	if len(base)<2 || !uniquealphaandnosigne(base){
		fmt.Print("NV")	
	}else if base=="0123456789"{
		return strconv.Itoa(n)
	}else{
		if n<0{
			n=-n
			str=Concat(str,"-")
		}
		ch:=basen(n,len(base))
		for _,val:= range ch{
			num,_:=strconv.Atoi(string(val))
			str=Concat(str,string(base[num]))
		}
		return str
	}
	return str
}

