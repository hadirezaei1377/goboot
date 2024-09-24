stack by slice:

func push(stack []int, value int) []int {
    return append(stack, value)
}

func pop(stack []int) ([]int, int) {
    length := len(stack)
    element := stack[length-1]
    return stack[:length-1], element
}

func main() {
    var stack []int

    stack = push(stack, 10)
    stack = push(stack, 20)
    stack = push(stack, 30)

    fmt.Println("stack:", stack)


    stack, popped := pop(stack)
    fmt.Println("deleted value:", popped)
    fmt.Println("stack:", stack)
}
