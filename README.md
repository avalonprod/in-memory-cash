## Fast and simple library written in go for hashing.
___
cash - is a simple and fast library written in go and is used to hash any data.

Download: `go get -u github.com/avalonprod/in-memory-cash`.

#### Methods
> Set 
```go
cash.Set(key, data)
```
+ indicator under which the value data will be written
___
> Get

```go

result, found := cash.Get(key)
```
+ key identifier by which to get the value
+ as a result, the result is returned if nothing is found, zero is returned
+ в found returns true if the element is found and false if no such element exists
___
> Delete
```go
err := cash.Delete(key)
```
+ key id at address need to delete value
+ returns an error if the element is not found
___

### Example

```go
package main

import "log"
import (
cash: "github.com/avalonprod/in-memory-cash"
      "log"
      "fmt"
)

func main() {
	cash.Set("user", "John")

	value, found := cash.Get("user")
	if !found {
		fmt.Print("Element is not found.")
	}
	err := cash.Delete("user")
	if err != nil {
		log.Fatal("Error element is not found")
	}
}
```
#### <a href="https://github.com/avalonprod/in-memory-cash/blob/main/LICENSE">License</a> 
