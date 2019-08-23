package main

import ("fmt")

func main(){
 
	for i:= 1; i<10; i++ {
		if i % 3==0 || i % 5==0{
		fmt.Println(i)
		}
	}
 sayi()
}
func sayi(){
    a:=0
	for i:= 1; i<1000; i++ {
		if i % 3==0 || i % 5==0{
			a=a+i
			fmt.Println(a)
		}
	}
}
