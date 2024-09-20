1. YAGNI (You Aren't Gonna Need It)

type Task struct {
    Title string
    Done  bool
}

func addTask(title string) Task {
    return Task{Title: title, Done: false}
}

2. DRY (Don't Repeat Yourself)

func calculateTax(amount float64) float64 {
    return amount * 0.15
}

func main() {
    price := 100.0
    tax := calculateTax(price)
    fmt.Println("Tax:", tax)
}

3. KISS (Keep It Simple, Stupid)

func bubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

func main() {
    numbers := []int{5, 2, 8, 3, 1}
    sorted := bubbleSort(numbers)
    fmt.Println("Sorted:", sorted)
}

4. Separation of Concerns

package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}





