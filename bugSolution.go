func main() {
    var m = make(map[string]int)
    val, ok := m["b"]
    if ok {
        fmt.Println("Value found:", val)
    } else {
        fmt.Println("Key not found") // Handle missing key gracefully
    }
} 