package main

import(
	"fmt"
	"sort"



)
type User struct{
		Nome string
		Idade int

}

type poridade []User
func(p poridade) Len()int{return len(p)}
func(p poridade) Less(i,j int) bool {return p[i].Idade<p[j].Idade}
func(p poridade) Swap(i,j int){p[i],p[j]= p[j],p[i]}



type pornome []User
func(p pornome) Len()int {return len(p)}
func(p pornome) Less(i,j int) bool{return p[i].Nome<p[j].Nome}
func(p pornome) Swap(i,j int){p[i],p[j]= p[j],p[i]}


func main(){
 
	user1:= User{"mateus",12}
	user2:= User{"julio",31}
	user3:=User{"iago",25}

users:=[]User{user1,user2,user3}
fmt.Println("\nLISTA ORIGINAL:",users)
sort.Sort(poridade(users))
fmt.Println("\nLISTA POR IDADE: ",users)
sort.Sort(pornome(users))
fmt.Println("\nLISTA POR NOME: ",users)
}