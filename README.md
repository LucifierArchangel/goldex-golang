### GoldEx Integration Go

#### Installation

```bash
go get github.com/LucifierArchangel/goldex-golang-integration
go install github.com/LucifierArchangel/goldex-golang-integration
```

#### Usage Example

##### Get All Requests

```go
package main

import (
    "context"
    "fmt"
    "os"

    GoldEx "github.com/LucifierArchangel/goldex-golang-integration"
)

func main() {
    client := GoldEx.NewClient("<YOUR_API_KEY>")

    ctx := context.Background()

    res, err := client.GetAllRequests(ctx)

    if err != nil {
        os.Exit(1)
    }

    fmt.Println((*res)[0])
}
```

##### Get Request Status By RequestId

```go
package main

import (
    "context"
    "fmt"
    "os"

    GoldEx "github.com/LucifierArchangel/goldex-golang-integration"
)

func main() {
    client := GoldEx.NewClient("<YOUR_API_KEY>")

    ctx := context.Background()

    res, err := client.GetRequestById(ctx, "<YOUR_REQUEST_ID>")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println((*res)[0].Status)
}
```

##### Create Request

```go
package main

import (
    "context"
    "fmt"
    "os"

    GoldEx "github.com/LucifierArchangel/goldex-golang-integration"
)

func main() {
    client := GoldEx.NewClient("<YOUR_API_KEY>")

    ctx := context.Background()

    res, err := client.CreateRequest(ctx, GoldEx.CreateRequest{
		Amount:     1000,
		Currency:   1,
		CardNumber: "Test",
		CardHolder: "Test",
		RequestId:  "Request ID Test",
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res)
}
```
