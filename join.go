package piscine

func Join(strs []string, sep string) string{
	result:=""
	for i,value:=range strs{
		if i!=len(strs)-1{
			result=concat1(result,value)
			result=concat1(result,sep)
		}else{
			result=concat1(result,value)
		}
	}
	return result	
}

func concat1(str1,str2 string) string{
	result:=make([]byte,len(str1)+len(str2))
	j:=0
	for i,_:=range result{
		if(i<len(str1)){
			result[i]=str1[i]
		}else{
			result[i]=str2[j]
			j++
		}
		
	}
	return string(result)
}

