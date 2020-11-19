package main

import "fmt"


//stack represents stack that hold a slice 

type Stack struct{
	items  []int
}

//push will add a  value  at  the end

func (s * Stack) Push(i int){
	s.items = append(s.items,i)
}

//pop  will remove  a value  at the end
//and returns  the removed value
func (s * Stack) Pop()int {
	l := len(s.items)-1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main(){
	mystack := Stack{}
	fmt.Println(mystack)
	mystack.Push(100)
	mystack.Push(200)
	mystack.Push(300)
	mystack.Push(400)
	fmt.Println(mystack)

	mystack.Pop()
	fmt.Println(mystack)
}

