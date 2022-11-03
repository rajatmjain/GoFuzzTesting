package main

func main() {
    
} 

func IsPositiveAndEven(input int)(bool,bool){
    if((input>0) && (input%2==0)){
        return true,true
    }else if((input<0) && (input%2==0)){
        return false,true
    }else if((input>0) && (input%2!=0)){
        return true,false
    }else{
        return false,false
    } 
}


