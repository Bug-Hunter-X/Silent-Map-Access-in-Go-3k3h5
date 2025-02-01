# Silent Map Access in Go

This example demonstrates a common, yet subtle issue in Go: accessing a missing key in a map.  Go's maps will not throw an error when accessing a nonexistent key, instead returning the zero value for the type. 

This can lead to hard-to-debug issues where the program continues execution with an unexpected value, rather than immediately revealing an error condition.

**bug.go** contains a simple map access which demonstrates the behavior.  **bugSolution.go** illustrates a solution for handling this issue.