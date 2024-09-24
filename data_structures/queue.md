
func enqueue(queue []int, value int) []int {
    return append(queue, value)
}

func dequeue(queue []int) ([]int, int) {
    element := queue[0]
    return queue[1:], element
}

func main() {
    var queue []int

    queue = enqueue(queue, 10)
    queue = enqueue(queue, 20)
    queue = enqueue(queue, 30)

    fmt.Println("queue:", queue)

    
    queue, dequeued := dequeue(queue)
    fmt.Println("deleted value:", dequeued)
    fmt.Println("value:", queue)
}
