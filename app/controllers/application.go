package controllers

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/xDarkicex/SimpleStore/helper"
)

// Application Controller.
type Application helper.Controller

//Index New index function
func (c Application) Index(a helper.RouterArgs) {
	user := Users{
		Name:   "Gentry",
		Age:    23,
		Weight: 189,
	}
	helper.Render(a, map[string]interface{}{
		"user": user,
	})
}

// AllProducts Product Responce
func (c Application) AllProducts(a helper.RouterArgs) {
	products := []Product{
		Product{Name: "Lamp", Price: 800, Tags: []string{"lighting", "Office", "design", "Furniture"}},
		Product{Name: "Mr. Clean", Price: 3, Tags: []string{"cleaning", "kitchen"}},
		Product{Name: "Shampoo", Price: 9, Tags: []string{"hygine", "shower", "soap", "bathroom", "cleaning"}},
		Product{Name: "coffee", Price: 3, Tags: []string{"drink", "Milk", "breakfast", "work"}},
		Product{Name: "Milk", Price: 3, Tags: []string{"dairy", "Milk", "breakfast", "lunch", "drink"}},
		Product{Name: "Desk", Price: 800, Tags: []string{"writing", "Office", "work", "Furniture", "executive"}},
		Product{Name: "Pencil", Price: 2, Tags: []string{"writing", "Office", "work", "school"}},
		Product{Name: "Paper", Price: 5, Tags: []string{"writing", "Office", "work", "school"}},
		Product{Name: "Pen", Price: 8, Tags: []string{"writing", "Office", "work", "school", "executive"}},
	}

	fmt.Fprintln(a.Response, products)
}

// Match Product Responce
func (c Application) Match(a helper.RouterArgs) {
	products := []Product{
		Product{Name: "Lamp", Price: 800, Tags: []string{"lighting", "Office", "design", "Furniture"}},
		Product{Name: "Mr. Clean", Price: 3, Tags: []string{"cleaning", "kitchen"}},
		Product{Name: "Shampoo", Price: 9, Tags: []string{"hygine", "shower", "soap", "bathroom", "cleaning"}},
		Product{Name: "coffee", Price: 3, Tags: []string{"drink", "Milk", "breakfast", "work"}},
		Product{Name: "Milk", Price: 3, Tags: []string{"dairy", "Milk", "breakfast", "lunch", "drink"}},
		Product{Name: "Desk", Price: 800, Tags: []string{"writing", "Office", "work", "Furniture", "executive"}},
		Product{Name: "Pencil", Price: 2, Tags: []string{"writing", "Office", "work", "school"}},
		Product{Name: "Paper", Price: 5, Tags: []string{"writing", "Office", "work", "school"}},
		Product{Name: "Pen", Price: 8, Tags: []string{"writing", "Office", "work", "school", "executive"}},
	}

	sort.Sort(byMatchedTags{
		Tags:     []string{"writing", "executive"},
		Products: products,
	})
	if len(products) >= 3 {
		products = products[0:3]
	}
	fmt.Println(products)
	pJSON, err := json.Marshal(products)

	if err != nil {
		helper.Logger.Println(err)
	}
	fmt.Fprintln(a.Response, string(pJSON))
}

func (p Product) String() string {
	return fmt.Sprintf("\nName: %s \nPrice: %d\nTags: %s\n ", p.Name, p.Price, p.Tags)
}

type byMatchedTags struct {
	Tags     []string
	Products []Product
}

func (a byMatchedTags) Len() int      { return len(a.Products) }
func (a byMatchedTags) Swap(i, j int) { a.Products[i], a.Products[j] = a.Products[j], a.Products[i] }
func (a byMatchedTags) Less(i, j int) bool {
	// Okay so we've got to check i and j for which has fewer matches
	matches := make(map[int]int)
	for _, iteration := range []int{i, j} {
		for _, tag := range a.Products[iteration].Tags {
			//fmt.Println(tag);
			for _, matchTag := range a.Tags {
				if tag == matchTag {
					matches[iteration]++
				}
			}
		}
	}
	return matches[i] > matches[j]
}

// Users ...
type Users struct {
	Name   string
	Age    int
	Weight int
}

// Products is all products
type Products struct {
	Product []Product
}

// Product is details about product
type Product struct {
	Name  string
	Price int
	Tags  []string
}
