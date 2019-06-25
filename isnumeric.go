package piscine

func IsNumeric(s string) bool{
	cbstr:=[]rune(s)
	for i,_:=range cbstr{	
		if !(cbstr[i]>=48 && cbstr[i]<=57){
			return false
		}
	}
	
	return true
}
