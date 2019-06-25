package piscine

import "fmt"
import "strconv"

func PrintNbrBase(n int,base string){//convertir un int dans une base et retourne le string de celui-ci
	if len(base)<2 || !uniquealphaandnosigne(base){
		fmt.Print("NV")	
	}else if base=="0123456789"{
		fmt.Print(n)
	}else{
		str:=""
		if n<0{
			n=-n
			str=Concat(str,"-")
		}
		ch:=basen(n,len(base))
		for _,val:= range ch{
			num,_:=strconv.Atoi(string(val))
			str=Concat(str,string(base[num]))
		}
		fmt.Print(str)
	}
}

func uniquealphaandnosigne(s string) bool{//verifier si une chaine contient que des caractères unique et pas de signe- ou +
	for i,val1:=range s{
		if val1== 43|| val1== 45{//pour les signe - et +
			return false		
		}
		for j:=i+1;j<len(s);j++{
			if string(val1)==string(s[j]){
				return false
			}
		}	
	}
	return true
}

func basen(nb,n int) []string{// pour convertir un nombre dans une base n
	var cr []string
	var c []string
	if nb<0{
		nb=-nb
		c=append(c,"-")
	}
	re:=0
	div:=0
	DivMod(nb, n, &div, &re)
	for div>0{	
		DivMod(nb, n, &div, &re)
		cr=append(cr,strconv.Itoa(re))
		nb=div
	}
	i:=len(cr)-1
	for 0<=i{
		c=append(c,string(cr[i]))
		i--	
	}
	return c
}
