# Async Loop

Async Loop is a package that provides functions to iterate over slices concurrently.

## Requirements

This package requires Go 1.21 or higher.

## Usage

This package simplifies the process of spawning multiple goroutines, which can be beneficial for fetching data from multiple APIs simultaneously, web scraping, and load testing websites and web applications by simulating high traffic, among other uses.

### Loop

A common pattern in Go is for-looping over a slice of elements in parallel using goroutines with wait groups and/or channels.

The `asyncloop.Loop` function offers this functionality through an easy-to-use interface.

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

The task runs in parallel, so the total operation will only take 1 second instead of the 5 it would take otherwise.

⚠️ Remember that each iteration runs in a separate goroutine, so you should ensure you are performing thread-safe operations.

Parallel tasks will not speed up compute-heavy operations; in such cases, you're better off using a normal loop. However, for network requests or asynchronous tasks, using `asyncloop.Loop` will improve performance.

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

The `asyncloop.Pool` function is similar to `asyncloop.Loop` but allows the caller to set the number of spawned goroutines with the second argument.

This feature is useful when you want bounded concurrency to avoid rate limits, for example.

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

In the above example, only 2 elements are processed at a time.
The return value of the iterator function indicates whether to stop or continue iterating over the slice.

⚠️ Be aware that goroutines already started will not be stopped; returning false will just prevent new goroutines from being spawned.

### Batch

The Batch function allows you to iterate over elements in batches. The size of each batch is determined by the given size argument, where a batch can be the full size or smaller.

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

⚠️ The output order may vary since it runs concurrently. If a batch size of 0 is passed, then no iterations of the loop are performed.

### Range

The Range function allows you to iterate over a range of integer values.

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

This function performs a parallel operation for a specified number of times.

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
