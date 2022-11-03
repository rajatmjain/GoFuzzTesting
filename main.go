package main

func main() {
    
} 

type MyDate struct{
    Day int
    Month int
}
func IsValidDate(date MyDate)(bool,bool){
    if((date.Day>=1 && date.Day<=31) && (date.Month>=1 && date.Month<=12)){
        return true,true
    }else if((date.Day<1 && date.Day>31) && (date.Month>=1 && date.Month<=12)){
        return false,true
    }else if((date.Day>=1 && date.Day<=31) && (date.Month<1 && date.Month>12)){
        return true,false
    }else{
        return false,false
    }  
}

func IsPositiveAndEven(input int)(bool,bool){
    if((input>=0) && (input%2==0)){
        return true,true
    }else if((input<0) && (input%2==0)){
        return false,true
    }else if((input>=0) && (input%2!=0)){
        return true,false
    }else{
        return false,false
    } 
}

