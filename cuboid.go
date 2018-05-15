package main

import "fmt"

// drawing te cuboid function
// 3 variables: drawX, drawY, drawZ
func cuboidDraw(drawX, drawY, drawZ int) {
	fmt.Printf("Cuboid %d %d %d: \n", drawX, drawY, drawZ)
	cubeLine(drawY+1, drawX, 0, "+-")
	for i := 1; i <= drawY; i++ {
		cubeLine(drawY-i+1, drawX, i-1, "/ |")
	}
	cubeLine(0, drawX, drawY, "+-|")
	for i := 4*drawZ - drawY - 2; i > 0; i-- {
		cubeLine(0, drawX, drawY, "| |")
	}
	cubeLine(0, drawX, drawY, "| +")
	for i := 1; i <= drawY; i++ {
		cubeLine(0, drawX, drawY-i, "| /")
	}
	cubeLine(0, drawX, 0, "+-\n")
}

// cubeLine function

func cubeLine(n, drawX, drawY int, cubeDraw string) {
	fmt.Printf("%*s", n+1, cubeDraw[:1])
	for d := 9*drawX - 1; d > 0; d-- {
		fmt.Print(cubeDraw[1:2])
	}
	fmt.Print(cubeDraw[:1])
	fmt.Printf("%*s\n", drawY+1, cubeDraw[2:])
}

// main function, where the function will ask the user
// to give inputs for dimensions of the cuboid
// l = length (drawX)
// w = width (drawY)
// h = height (drawZ)

func main() {
	fmt.Println("Enter 3 dimensions of Cuboid : ")
	var l, w, h int
	fmt.Scanln(&l)
	fmt.Scanln(&w)
	fmt.Scanln(&h)
	cuboidDraw(l, w, h)
}
