package main

import "fmt"
 
func main() {
   c := color("Red")		// I don't follow why it's color(). Why not c:= "Red" 							// The line "c := "Red" " doesn't work
 
   fmt.Println(c.describe("is an awesome color"))
}
 
type color string
 
func (c color) describe(description string) (string) {
   return string(c) + " " + description
}