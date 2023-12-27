package main

import ("fmt"
        "bufio"
        "os"
)

func main(){
             welcome:= "welcome to user input";
             fmt.Println(welcome);

            reader := bufio.NewReader(os.Stdin)
            fmt.Println("please enter the rating for our pizza:");

			//comma ok syntax || error ok syntax
			input, _ := reader.ReadString('\n');
			fmt.Println("Thanks for rating", input)
			fmt.Printf("type of vlaue is: %T \n",input)

}
