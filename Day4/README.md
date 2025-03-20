# Day 4

## Lab - Golang modules

Let's create 3 modules in your home directory
```
cd ~
mkdir addition subtraction main
```

Let's create the first golang module called addition
```
cd ~/addition
go mod init addition
cat go.mod
touch addition.go
```

Update the addition.go file with the below content
<pre>
package addition

func Add( x float32, y float32 ) float64 {
  return float64( x + y )
}
</pre>

Let's create the second golang module called subtraction
```
cd ~/subtraction
go mod init subtraction
cat go.mod
touch subtraction.go
```

Update the subtraction.go file with the below content
<pre>
package subtraction

func Subtract( x float32, y float32 ) float64 {
  return float64(x-y)
}
</pre>

Let's create the main module
```
cd ~/main
go mod init main
cat go.mod
touch main.go
```

Update the main.go with the below content
<pre>
package main

import (
   "fmt"
   "addition"
   "subtraction"
)
  
func main( ) {
   x := 500.7
   y := 200.5
   result1 := addition.Add( x, y )
   result2 := subtraction.Subtract( x, y )

   fmt.Println ( "The sum of ", x, " and ", y, " is ", result1 )
   fmt.Println ( "The difference of ", x, " and ", y, " is ", result2 )
}
</pre>

Running the above application
```
cd ~/main
go mod edit --replace addition=../addition
go mod edit --replace subtraction=../subtraction
go mod tidy
go run ./main.go
```

Expected output
