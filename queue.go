package main

import "fmt"

//Queue represents a  queue  that  holds  a slice 

type Queue struct {
	items [] int
}


//Enqueue adds a value  at  the end 

func (q *Queue)Enqueue(i int){
	q.items = append(q.items,i)
}

func main(){
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Enqueue(400)
	myQueue.Enqueue(500)
	myQueue.Enqueue(600)
	myQueue.Enqueue(700)
	fmt.Println(myQueue)

}