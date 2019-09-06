package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []card

func newDeck() deck {
	cards := deck{}

	colors := []string{"Red", "Green", "Yellow", "Blue"}
	faces := make([]int, 10)
	for i := 0; i <= 9; i++ {
		faces[i] = i
	}

	for _, color := range colors {
		for _, face := range faces {
			c := card{
				cardType: "Regular",
				color:    color,
				value:    face,
			}
			cards = append(cards, c)
			if face != 0 {
				cards = append(cards, c)
			}
		}
	}

	specialTypes := []string{"Reverse", "Skip", "Draw Two"}
	for _, color := range colors {
		for _, cardType := range specialTypes {
			if cardType == "Draw Two" {
				c := card{
					cardType: cardType,
					color:    color,
					value:    2,
				}
				cards = append(cards, c)
				cards = append(cards, c)
			} else {
				c := card{
					cardType: cardType,
					color:    color,
				}
				cards = append(cards, c)
				cards = append(cards, c)
			}
		}
	}

	wildcards := []string{"Draw Four", "Wildcard"}
	for _, cardType := range wildcards {
		c := card{
			cardType: cardType,
		}
		for i := 0; i < 4; i++ {
			cards = append(cards, c)
		}
	}

	return cards
}

func (d deck) length() int {
	for i, c := range d {
		if c.cardType == "" {
			return i
		}
	}
	return len(d)
}

func (d deck) print(valids []int) {
	isValid := []bool{}
	i := 0
	for _, v := range valids {
		for ; i < v; i++ {
			isValid = append(isValid, false)
		}
		isValid = append(isValid, true)
		i++
	}
	for len(isValid) < d.length() {
		isValid = append(isValid, false)
	}

	for i, c := range d {
		if c.cardType == "" {
			return
		}
		fmt.Printf("%d. ", i+1)
		c.print()
		if isValid[i] {
			fmt.Println(" (*)")
		} else {
			fmt.Println()
		}
	}
}
func (d deck) getTop() card {
	return d[d.length()-1]
}

func (d deck) printTop() {
	d.getTop().print()
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	r := rand.New(src)
	for i := 0; i < d.length(); i++ {
		newIndex := r.Intn(d.length())
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}

func (d deck) deal(n int) deck {
	hand := make(deck, 108)
	l := d.length()

	for i := 0; i < n; i++ {
		hand[i], d[l-1-i] = d[l-1-i], card{}
	}
	return hand
}

func (d deck) draw(n int, od deck) {
	l, ol := d.length(), od.length()

	for i := 0; i < n; i++ {
		d[l+i], od[ol-1-i] = od[ol-1-i], card{}
	}
}

func (d deck) play(i int, dp deck) card {
	playedCard := d[i]
	dp[dp.length()] = playedCard

	for j := i; j < d.length(); j++ {
		d[j] = d[j+1]
	}
	return playedCard
}
