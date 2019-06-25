package piscine

func IsUpper(s string) bool{
	cbstr:=[]rune(s)
	for i,_:=range cbstr{	
		if !(cbstr[i]>=65 && cbstr[i]<=90){
			return false
		}
	}
	
	return true
}
