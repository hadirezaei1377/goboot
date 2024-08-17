Format Verbs

%d

func main() {
    age := ۳۰
    fmt.Printf("سن شما: %d\n", age)
}
 
output:
سن شما: ۳۰

%t

func main() {
    isValid := true
    fmt.Printf("مقدار صحیح است؟ %t\n", isValid)
}

output:
مقدار صحیح است؟ true

%v

func main() {
    name := "علی"
    age := ۳۰
    isStudent := false
    fmt.Printf("نام: %v، سن: %v، دانشجو؟ %v\n", name, age, isStudent)
}

output:
نام: علی، سن: ۳۰، دانشجو؟ false

%s

func main() {
    city := "تهران"
    fmt.Printf("شهر شما: %s\n", city)
}

output:
شهر شما: تهران

%f

func main() {
    price := ۲۵۰۰.۵۰
    fmt.Printf("قیمت: %f\n", price)
}

output:
قیمت: ۲۵۰۰.۵۰۰۰۰۰

%p

func main() {
    var x int = ۱۰
    fmt.Printf("آدرس حافظه x: %p\n", &x)
}

output:
آدرس حافظه x: 0x...

