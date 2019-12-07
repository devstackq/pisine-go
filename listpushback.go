package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

//  List need for - know where - head and tail
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {

	node := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node

}

// func (n *Node) InsertEnd(newValue interface{}) {
// 	//if head empty
// 	if n == nil {
// 		fmt.Println("Empty")
// 		os.Exit(1)
// 	}

// 	newNode := new(Node)

// 	newNode.element = newValue

// 	newNode.next = n.next

// 	n.next = newNode
// }

// //read
// func (n *Node) Read() {
// 	fmt.Println(n.element)
// }

// //write
// func (n *Node) Write(value interface{}) {
// 	n.element = value
// }

// // 1
// type Node struct {
// 	next    *Node
// 	element interface{}
// }

// func (n *Node) InsertEnd(newValue interface{}) {
// 	//if head empty
// 	if n == nil {
// 		fmt.Println("Empty")
// 		os.Exit(1)
// 	}
// 	// create new class
// 	newNode := new(Node)
// 	// write value from - new va;ue -> inside newNode
// 	newNode.element = newValue
// 	// recursion -> call
// 	newNode.next = n.next

// 	n.next = newNode
// }
// func (n *Node) PrintList() {

// 	for n != nil {
// 		fmt.Println(n.element)
// 		n = n.next
// 	}
// }

// func main() {

// 	head := new(Node)

// 	head.element = 43

// 	head.InsertEnd(99)
// 	//head.next.Write(10)
// 	head.PrintList()
// }

// func newNode(element interface{}) *Node {
// 	node := new(Node)
// 	node.element = element
// 	return node
// }

// func appendEnd(head *Node) (element interface{}) {

// 	node := newNode(element)
// 	if head == null {
// 		head = node

// 	} else {
// 		head.next = node
// 		node.previous = tail
// 	}
// 	previous = node
// }

// func (list *LinkedList) AddToEnd(element interface{}) {
// 	node := newNode(element)
// 	if list.isEmpty() {
// 		list.head = node
// 	} else {
// 		list.tail.next = node
// 		node.previous = list.tail
// 	}
// 	list.tail = node
// }
