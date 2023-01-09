/*
Create a struct called book having variables author, title,  publisher, price, year of publication and Edition.
Create 2 struct variables and call method update to update book price,  year of publication and Edition.

Call another method display to show updated details of book in a proper reference format
(Example: John Badner,"Learning Go: An Idiomatic Approach to Real World Go Programming", Oreilly, Edition 2, 2021.).
Make another function which takes number of copies as input and calculate total price of book.
Provide 20% discount if the year of publication of ordered book is before 2020.
*/

package main

import "fmt"

type book struct {
	author, title, publisher string
	price, year, edition     int
}

func (b *book) update(price, year, edition int) {
	b.price = price
	b.year = year
	b.edition = edition
}

func (b book) display() {
	fmt.Println("Book Details:")
	fmt.Printf("Author: %s, Title: \"%s\", Publisher: %s, Edition: %d, Year: %d",
		b.author, b.title, b.publisher, b.edition, b.year)
}

func (b book) total(copies int) float64 {
	total := float64(b.price * copies)
	if b.year < 2020 {
		total *= 0.8
	}
	return total
}

func main() {
	b1 := book{
		author:    "John Badner",
		title:     "Learning Go: An Idiomatic Approach to Real World Go Programming",
		publisher: "Oreilly",
		price:     100,
		year:      2015,
		edition:   1,
	}

	b2 := book{
		author:    "Karthik Banjan",
		title:     "Self Help",
		publisher: "Banjan Ltd.",
		price:     100,
		year:      2021,
		edition:   2,
	}

	b1.update(200, 2019, 2)
	b2.update(400, 2024, 4)
	b1.display()
	fmt.Println()
	b2.display()
	fmt.Println()
	fmt.Println("Total Cost of 5 copies of b1:", b1.total(5))
	fmt.Println("Total Cost of 2 copies of b2:", b2.total(5))
}
