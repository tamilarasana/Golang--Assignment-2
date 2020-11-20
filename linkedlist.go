package main


import "fmt"

type node struct{

	data int
	next *node
}

func insert	(head*node, data int )*node{
	n := &node{data:data}

	if  head == nil{
		return n 
	}else{
		n.next = head
		return n
	}
}

func Printlist(head *node){
	for head != nil {
		fmt.Println(head.data, "->")
		head =head.next
	}

	fmt.Println(nil)
}

func main(){
	var link *node

	link = insert(link,1)
	link = insert(link,2)
	link = insert(link,3)
	link = insert(link,4)
	link = insert(link,5)
	link = insert(link,6)
	link = insert(link,7)
	link = insert(link,8)
	link = insert(link,9)
	link = insert(link,10)

	Printlist(link)
}