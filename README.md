# go-fakturoid

## Connecting to API

```go
package main

import (
	"fmt"
	go_fakturoid "github.com/MLJ-solutions/go-fakturoid"
)

func main() {
	client, _ := go_fakturoid.New(&go_fakturoid.Options{
		Creds: go_fakturoid.NewCredentials("EMAIL", "SLUG", "PRIVATE KEY"),
	})

	fmt.Println(client)
}
```

## Full Example of getting account setting

```go
package main

import (
	"fmt"
	go_fakturoid "github.com/MLJ-solutions/go-fakturoid"
	"log"
)

func main() {
	client, _ := go_fakturoid.New(&go_fakturoid.Options{
		Creds: go_fakturoid.NewCredentials("EMAIL", "SLUG", "PRIVATE KEY"),
	})

	fmt.Println(client)

	account, err := client.Account()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(account)
}
```
