package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fanta()
	k := anything ( 500 )
	fmt.Println(k)
	strings()
	intergers()
	boolean()
}
func fanta()  {
	fmt.Println("my name is CBoy")
}
func anything (population int) (int){

	population = population *2

	return population
}
func strings() {
	a:= "Once Upon "
	b:= "A Time "
	c:= "In Hollywood"
	d:= a + b + c
	fmt.Println(d)

}
func intergers(){
	a:= 5
	peopleinlagos:= 2500
	years_stayed:= 10
	total_people_in_lagos:= (peopleinlagos +years_stayed)*a
	fmt.Println(total_people_in_lagos)
}
func  boolean(){
	senior := false
	ade:= 25
	tobi:= 50
	if tobi > ade*3 {
		senior = false
	}
	fmt.Println(senior)

}