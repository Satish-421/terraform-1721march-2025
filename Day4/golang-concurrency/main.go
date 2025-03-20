package main

import (
  "fmt"
  "time"
)

func firstFunction( count int ) {
   for i := 0; i<count; i++ {
	fmt.Println ("First function", i )
	time.Sleep( time.Millisecond * 5 )
   }
}

func secondFunction( ) {
   for i := 0; i<1000; i++ {
	fmt.Println ("Second function", i )
	time.Sleep( time.Millisecond * 5 )
   }
}

func main() {
  fmt.Println ( "Press any key to exit ..." )

  //firstFunction and secondFunction will run at the same time as different threads
  go firstFunction( 1000 )
  go secondFunction( )

  var tmp string
  fmt.Scanln(&tmp )
}
