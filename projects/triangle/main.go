package main
import "fmt"
import "math"
func main(){

  var a, b float64
  var c float64
  fmt.Scan(&a)
  fmt.Scan(&b)
  c = a * a + b * b
  fmt.Println(math.Sqrt(c))
}
