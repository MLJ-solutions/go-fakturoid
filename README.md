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

## Full Example of creating invoice
```go
package main

import (
	"fmt"
	. "github.com/MLJ-solutions/go-fakturoid"
	"github.com/MLJ-solutions/go-fakturoid/models"
	"log"
	"time"
)

func main() {
	client, _ := New(&Options{
		Creds: NewCredentials("EMAIL", "SLUG", "PRIVATE KEY"),
	})
	
	subject, err := client.CreateSubject(models.Subject{
		Name:           "Name",
		Street:         "Ulice adresa",
		City:           "City",
		Zip:            "12345",
		Country:        "CZ",
		RegistrationNo: "123456789",
		FullName:       "Full name",
		Email:          "developer@mlj.solutions",
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(subject)

	createInvoice := models.Invoice{
		SubjectId:             subject.Id,
		Currency:              models.CurrencyCZK,
		PaymentMethod:         models.PaymentMethodCard,
		Due:                   14,
		IssuedOn:              models.FakturoidDate(time.Now()),
		TaxableFulfillmentDue: models.FakturoidDate(time.Now()),
		Lines:                 []models.InvoiceLine{{Name: "name", Quantity: 1, UnitName: "ks", UnitPrice: 100, VatRate: 15}},
	}

	invoice, err := client.CreateInvoice(createInvoice)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(invoice)
}

```

```go
package main

import (
	go_fakturoid "github.com/MLJ-solutions/go-fakturoid"
	"log"
)

func main() {
	client, _ := go_fakturoid.New(&go_fakturoid.Options{
		Creds: go_fakturoid.NewCredentials("EMAIL", "SLUG", "PRIVATE KEY"),
	})

	log.Println(client)

	account, err := client.Account()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(account)
}
```

## Full Example of getting account setting

```go
package main

import (
	go_fakturoid "github.com/MLJ-solutions/go-fakturoid"
	"log"
)

func main() {
	client, _ := go_fakturoid.New(&go_fakturoid.Options{
		Creds: go_fakturoid.NewCredentials("EMAIL", "SLUG", "PRIVATE KEY"),
	})

	log.Println(client)

	account, err := client.Account()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(account)
}
```
