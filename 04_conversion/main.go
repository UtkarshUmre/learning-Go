package main

import ("fmt"
         "bufio"
         "os"
         "strconv"
         "strings"
        )

func main () {
            fmt.Println("welcome to our pizza app")
            fmt.Println("please rate our pizza between 1 to 5")


            
            reader := bufio.NewReader(os.Stdin)
            input, _ := reader.ReadString('\n')
            fmt.Println("thanks for rating",input)

            numRating , err := strconv.ParseFloat(strings.TrimSpace(input), 64)

            if err != nil{
                fmt.Println(err)
            }else{
                fmt.Println("added one to your rating",numRating + 1)
            }

            // ratingPlusOne = input + 1
            // fmt.Println(ratingPlusOne)

            
}

