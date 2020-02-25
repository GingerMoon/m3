package main

import "github.com/GingerMoon/m2"

// go.mod:
//require (
//    github.com/GingerMoon/m1 v0.0.2
//    github.com/GingerMoon/m2 v0.0.2
//)
//func main() {
//    m1.Print() // 2
//    m2.Print() // 2
//}

// go.mod:
//require (
//    github.com/GingerMoon/m1 v0.0.1
//    github.com/GingerMoon/m2 v0.0.2
//)
// go build will change go.mod: github.com/GingerMoon/m1 v0.0.1 -> github.com/GingerMoon/m1 v0.0.2
//func main() {
//   m1.Print() // 2
//   m2.Print() // 2
//}

// go.mod:
//require (
//    github.com/GingerMoon/m1 v0.0.1
//    github.com/GingerMoon/m2 v0.0.2
//)
// go build will change go.mod: github.com/GingerMoon/m1 v0.0.1 -> github.com/GingerMoon/m1 v0.0.2
func main() {
    m2.Print() // 2
}
