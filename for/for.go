package main

import "fmt"

func main()  {

  fmt.Println("GOLANG for LOOP")

  i := 1
  for i < 3 {
    fmt.Println(i)
    i++
  } // end first loop


  for j :=1 ; j < 6; j++ {
    fmt.Println(j)
  } // end second loop

  for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
  } // end third loop

}
