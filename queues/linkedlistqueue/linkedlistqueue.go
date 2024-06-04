package linkedlistqueue


type Node[T any] struct{
	var val T
	var next *Node
} 

type Queue[T any] struct{
	var head *Node
	var tail *Node
}