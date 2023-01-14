package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// for upper case strings.ToUpper(a)  ---- for lower strings.ToLower(varName)
	a := "UndersStasdnbadaAsdf"
	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToLower(a))
	// strings.HasPrefix ---> searches the string from the beginning
	// strings.HasSuffix ----> searches the string from the end
	// strings.Contains ----> searches any where in the string
	// strings.Count ----> Count how many times the string appears in the string
	b := "Hello world"
	fmt.Println(strings.HasSuffix(b, "world"))
	fmt.Println(strings.HasPrefix(b, "Hello"))
	fmt.Println(strings.Contains(b, "w"))
	fmt.Println(strings.Contains(b, "w "))
	fmt.Println(strings.Count(b, "l"))
	c := []string{"sharks", "crabs", "whales"}
	fmt.Println(strings.Join(c, "_"))
	fmt.Println(strings.Join(c, ", "))
	w := "hello World"
	x := "world hello"
	fmt.Println(strings.ToUpper(w)) //entire strings to upper case
	fmt.Println(strings.ToLower(w)) // entire strings to lower case
	fmt.Println(strings.Title(w))   // capatilize the first letter
	fmt.Println(strings.Compare(x, w))
	fmt.Println(strings.HasPrefix(w, "hell"))
	var sb strings.Builder
	fmt.Println("This is a string builder: ", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteString("Hello")
	fmt.Println("This is a string builder: ", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Grow(100)
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Reset()
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Write([]byte{65, 65, 65})
	fmt.Println("This is a string builder: ", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	h := "123"
	fmt.Printf("%T type: %v\n", h, h)
	y, _ := strconv.Atoi(h)
	fmt.Printf("%T type: %v\n", y, y)

}
