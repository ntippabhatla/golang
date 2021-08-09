package main
import (
	"fmt"
)

type test struct {
	count bool
	age int32
	weight float32

}

func main() {
	var t1 test
	fmt.Printf("%+v\n", t1 )
}
