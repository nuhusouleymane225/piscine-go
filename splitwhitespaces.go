package piscine

func SplitWhiteSpaces(str string) []string{
	str=Concat(str," ")
	var result []string
	word:=""
	for index,caract:=range str{
		if caract!=9 && caract!=10 && caract!=32{
			word=Concat(word,string(caract))
		}
		if (caract==9 || caract==10 || caract==32) && index>0{
			if str[index-1]!=9 && str[index-1]!=10 && str[index-1]!=32 {
				result=append(result,word)
				word=""
			}
			
		}
		
	}
	return result
}

