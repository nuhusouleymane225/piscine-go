package piscine

func IsLower(s string) bool{
	cbstr:=[]rune(s)
	for i,_:=range cbstr{	
		if !(cbstr[i]>=97 && cbstr[i]<=122){
			return false
		}
	}
	
	return true
}
