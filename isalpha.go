package piscine

func IsAlpha(s string) bool{
	cbstr:=[]rune(s)
	for i,_:=range cbstr{	
		if !((cbstr[i]>=48 && cbstr[i]<=57) || (cbstr[i]>=65 && cbstr[i]<=90) || (cbstr[i]>=97 && cbstr[i]<=122)){
			return false
		}
	}
	
	return true
}
