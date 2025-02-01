func main() {
    var m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["b"]) // this will print 0, not throw an error
}