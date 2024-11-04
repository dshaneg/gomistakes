// Demonstrates that the error #32 in the book is now outdated.
package main

import "fmt"

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		s.m[customer.ID] = &customer
	}
}

func main() {
	store := Store{m: make(map[string]*Customer)}

	store.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})

	for k, c := range store.m {
		fmt.Printf("key=%v, value=%v\n", k, c)
	}
}
