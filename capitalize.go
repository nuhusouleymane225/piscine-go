package piscine

func Capitalize(s string) string{
	cbstr:=[]rune(s)
	if cbstr[0]>=97 && cbstr[0]<=122{//on met le premier caractère en majuscule s'il est alphanumérique
		cbstr[0]-=32
	}
	for i,_:=range cbstr{	
		if !((cbstr[i]>=48 && cbstr[i]<=57) || (cbstr[i]>=65 && cbstr[i]<=90) || (cbstr[i]>=97 && cbstr[i]<=122)){/*on regarde si le caractère 
		n'est pas alphanumérique, si c'est le cas on met le prochain caractère alphanumerique en majuscule*/
			if i<len(cbstr)-1{
				if cbstr[i+1]>=97 && cbstr[i+1]<=122{
					cbstr[i+1]-=32
				}
			}
		}else{//si le caractère est alphanumerique on regarde s'il ne se trouve pas en debut de mot, si c'est le cas on le met en miniscule
			if i> 0{
				if ((cbstr[i-1]>=48 && cbstr[i-1]<=57) || (cbstr[i-1]>=65 && cbstr[i-1]<=90) || (cbstr[i-1]>=97 && cbstr[i-1]<=122)){
					if cbstr[i]>=65 && cbstr[i]<=90{
						cbstr[i]+=32
					}
				}
			}
		}
	}
	return string(cbstr)
}
