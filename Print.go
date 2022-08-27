package Main

//you need to import fmt
import (
	"fmt"
)

func main() {
	//then you can use:
	fmt.Print("this message will be printed without a linefeed at the end")
	//or
	fmt.Println("this message will be printed with a linefeed a the end")
	//or										   the \n is a line feed
	fmt.Printf("You can also use %s formatings with Printf \n (PS : this message won't end with a line feed)", "multiples")

	fmt.Print("Hello World")

}

//Package Main
// //import (
// 	fmt
// ) so we cant print messages

// func main() {

// }
