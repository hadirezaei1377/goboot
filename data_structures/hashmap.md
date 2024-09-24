
func main() {
    dictionary := map[string]string{
        "Go":     "A programming language.",
        "Python": "Another popular language.",
    }

    dictionary["Java"] = "Yet another language."

    fmt.Println("Python:", dictionary["Python"])


    delete(dictionary, "Java")
    fmt.Println("dictionary:", dictionary)
}
