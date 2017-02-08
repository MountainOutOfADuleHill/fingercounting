
/*
August 22nd, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
*/

package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  fingerValues := [10]int{10, 10, 10, 10, 50, 5, 1, 1, 1, 1}
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("0 for finger down, 1 for finger up. 10 total digits:")
  code, _ := reader.ReadString('\n')
  
  var count int

  for x, num := range code{
    if num == '1'{count += fingerValues[x]}
  }

  fmt.Println("The total count of your entry is ", count)
}
