package main

import "fmt"

func TranspileToGo(c JClass) string {
    return fmt.Sprintf(`package main
import "fmt"

func main() {
    fmt.Println(%s)
}
`, c.Body)
}
