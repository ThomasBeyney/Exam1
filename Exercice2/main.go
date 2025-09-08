package main

import "fmt"

func Lettre(str string){
	j := 0
	for i := 0 ; i <= len(str); i++ {
		if str[i] (str[i] <= 'z' && str[i] >= 'a') || (str[i] <= 'Z' && str[i] >= 'A'){
			j++
		}
		fmt.Print(j)
	}
}

func main(){
	fmt.Println(Lettre("Hello,World!"))
}

//t'embète pas sa marche pas :/