
type Node struct {
    Value int
    Next  *Node
}

func main() {

    node1 := &Node{Value: 10}
    node2 := &Node{Value: 20}
    node3 := &Node{Value: 30}

   
    node1.Next = node2
    node2.Next = node3

    current := node1
    for current != nil {
        fmt.Println(current.Value)
        current = current.Next
    }
}
