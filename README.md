# Async Loop

Async Loop is a package that provides functions to range over slices concurrently

## Requirements

This package requires Go 1.21 or higher.

## Usage

This package aims to simplify spawning multiple goroutines, which can help
fetching data from multiple APIs at once, web scraping, load testing websites and web apps by simulating high traffic, etcetera.

### Loop

A commonly used pattern in Go is to iterate over a slice of elements in parallel with a wait group.

The Loop function provides this functionality in an easy to use interface

```go
package main

import (
	"fmt"
	"time"
	"github.com/danielmesquitta/asyncloop"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	squares := make([]int, len(slice))

	asyncloop.Loop(slice, func(i int, val int) {
		// Each iteration runs in a goroutine
		squares[i] = val * val
		// Simulate a long running task
		time.Sleep(time.Second)
	})

	fmt.Println(squares) // [1 4 9 16 25]
}
```

The above task will run in parallel, which means the total operation will only take 1 second,
instead of the 5 it would take otherwise.

⚠️ One thing to be aware of is that each iteration runs in a separate goroutine. Therefore
you'll want to make sure you are performing thread safe operations.

The parallel task won't speed up any compute heavy operations, in that case, you're better off using a normal loop. However, in the event of performing network requests or async tasks, then using asyncloop.Loop will improve performance.

```go
package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/danielmesquitta/asyncloop"
)

func main() {
	colors := []string{"green", "yellow", "blue"}

	results := make([]*http.Response, len(colors))
	asyncloop.Loop(colors, func(i int, color string) {
		_, err := http.Post(
			"http://example.com/colors",
			"text/plain",
			strings.NewReader(color),
		)
		if err != nil {
			log.Println("oops", err)
		}
	})
}
```

### Pool

The pool function is very similar to `asyncloop.Loop`, however it allows to caller to set the spawned goroutines amount with the second argument.

This is useful in the event you want bounded concurrency.

```go
package main

import (
	"time"

	"github.com/danielmesquitta/asyncloop"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	size := 2

	asyncloop.Pool(slice, size, func(i int, val int) bool {
		// Simulate a long running task
		time.Sleep(time.Second)
		// Return true to continue the loop,
		// false to stop iterating
		return true
	})
}
```

In the above example, only 2 elements will be performed at a time.
The return of the iterator function serves to indicate if you want to stop or continue iterating over the slice

⚠️ One thing to be aware is that goroutines already started won't be stopped,
returning false will just stop new goroutines from being spawned

### Batch

The Batch function provides the ability to range over elements in batches. The size of each batch is decided by the given size argument, in which a batch will either be the same size or less than.

```go
package main

import (
	"fmt"

	"github.com/danielmesquitta/asyncloop"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	asyncloop.Batch(slice, 2, func(i int, batch []int) {
		fmt.Println(i, batch)
	})
}
```

The above code will print something like the following output:

```
2 [5]
0 [1 2]
1 [3 4]
```

If a batch size of 0 is passed in, then no iterations of the loop are performed.

⚠️ Note that the order in which it will print may vary since it is running concurrently

### Range

The range function allows you to iterate over a range of integer types.

```go
package main

import (
	"fmt"

	"github.com/danielmesquitta/asyncloop"
)

func main() {
	asyncloop.Range(2, 5, func(i int) {
		fmt.Println(i)
	})
}
```

The above code will print out

```
2
3
4
```

The `asyncloop.Range` method includes the starting value, but excludes the stop value.

⚠️ The output sequence might differ as it is executed concurrently.

### LoopN

This method allows to perform a parallel operation for a given number of times.

For example

```go
package main

import (
	"fmt"

	"github.com/danielmesquitta/asyncloop"
)

func main() {
	asyncloop.LoopN(3, func(i int) {
		fmt.Println(i)
	})
}
```

The above code will print out

```
0
1
2
```

⚠️ Keep in mind that the print order may vary due to concurrent execution.
