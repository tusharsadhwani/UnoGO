package main

import (
	"fmt"
	"math/rand"
	"os"
)

func validIndices(h deck, dp deck) []int {
	valids := []int{}
	topCard := dp.getTop()
	for i := 0; i < h.length(); i++ {
		c := h[i]

		switch c.cardType {
		case "Draw Four", "Wildcard":
			valids = append(valids, i)
		case "Reverse", "Skip", "Draw Two":
			switch topCard.cardType {
			case "Regular":
				if topCard.color == c.color {
					valids = append(valids, i)
				}
			case "Reverse", "Skip", "Draw Two":
				if topCard.color == c.color {
					valids = append(valids, i)
				}
			case "Draw Four", "Wildcard":
				valids = append(valids, i)
			}
		case "Regular":
			switch topCard.cardType {
			case "Regular":
				if topCard.color == c.color {
					valids = append(valids, i)
				} else if topCard.value == c.value {
					valids = append(valids, i)
				}
			case "Reverse", "Skip", "Draw Two":

				if topCard.color == c.color {
					valids = append(valids, i)
				}
			case "Draw Four", "Wildcard":
				valids = append(valids, i)
			}
		}
	}
	return valids
}

func showUserCards(h deck, valids []int, dp deck) {
	fmt.Println("\nTop Card:")
	dp.printTop()
	fmt.Println()
	// fmt.Println("Discard Pile:")
	// dp.print([]int{})
	// fmt.Println()

	fmt.Println("\nYour cards:")
	h.print(valids)
}

func playUser(hand deck, i int, dp deck) {
	cardIndex := i - 1
	playedCard := hand.play(cardIndex, dp)
	fmt.Print("\nYou played a ")
	playedCard.print()
	fmt.Println()
	// specialEffects(playedCard, 0) //0 is index of user, 1 for comp
	//TODO: IMPLEMENT SPECIAL PLAYED CARD LOGIC
}

func reshuffle(c deck, dp deck) {
	cl := c.length()
	dpl := dp.length()
	for i := 0; i < dpl-1; i++ {
		c[cl+i], dp[i] = dp[i], card{}
	}
	dp[0], dp[dpl-1] = dp[dpl-1], card{}
	c.shuffle()
}

func getRandomValidCard(c deck, h deck, dp deck) int {
	for {
		valids := validIndices(h, dp)
		if len(valids) == 0 {
			fmt.Println("Computer draws a new card...")
			if c.length() < 1 {
				fmt.Println("Empty deck, reshuffling...")
				reshuffle(c, dp)
			}
			// println(c.length())
			h.draw(1, c)
			continue
		}
		randomIndex := valids[rand.Intn(len(valids))]
		return randomIndex
	}
}

func playComp(c deck, h deck, dp deck) {
	index := getRandomValidCard(c, h, dp)
	playedCard := h.play(index, dp)
	fmt.Print("\nComputer plays a ")
	playedCard.print()
	//TODO: IMPLEMENT SPECIAL PLAYED CARD LOGIC
}

func cardIsValid(cardNumber int, valids []int) (valid bool) {
	cardIndex := cardNumber - 1
	for _, validIndex := range valids {
		if cardIndex == validIndex {
			valid = true
			return
		}
	}
	return
}

func userTurn(c deck, h deck, dp deck) {
	for {
		valids := validIndices(h, dp)
		showUserCards(h, valids, dp)

		if len(valids) == 0 {
			fmt.Println("\nLooks like you have no valid cards.")
			fmt.Print("\nPress enter to pick a card: ")
			readInt(os.Stdin)
			if c.length() < 1 {
				fmt.Println("Empty deck, reshuffling...")
				reshuffle(c, dp)
			}
			// println(c.length())
			h.draw(1, c)
			continue
		}

		for {
			fmt.Print("\nChoose: ")
			cardNumber := readInt(os.Stdin)

			if cardNumber <= 0 || cardNumber > h.length() {
				fmt.Println("Index out of range.")
			} else if !cardIsValid(cardNumber, valids) {
				fmt.Println("Invalid choice, choose a valid card.")
			} else {
				playUser(h, cardNumber, dp)
				return
			}
		}
	}

}

func compTurn(c deck, h deck, dp deck) {
	fmt.Println("Computer cards:")
	h.print([]int{})

	playComp(c, h, dp)
}

func newGame() {
	cards := newDeck()
	cards.shuffle()

	userHand := cards.deal(7)
	compHand := cards.deal(7)
	discardPile := cards.deal(1)

	for x := 0; x < 100; x++ {
		userTurn(
			cards, userHand, discardPile,
		)
		// print(len(cards))
		if userHand.length() == 0 {
			println("\nYou won!")
			return
		}
		compTurn(
			cards, compHand, discardPile,
		)
		// print(len(cards))
		if compHand.length() == 0 {
			println("\nComputer won!")
			return
		}
	}
}

func mainMenu() (exit bool) {
	fmt.Println("------------------------------")
	fmt.Println("1. New Game")
	fmt.Println("2. Help")
	fmt.Println("3. Exit")
	for {
		fmt.Print("\nEnter: ")
		choice := readInt(os.Stdin)

		switch choice {
		case 1:
			fmt.Println("\nStarting New Game!:")

			newGame()
			return
		case 2:
			fmt.Println("Help Message")
			return
		case 3:
			exit = true
			return
		default:
			fmt.Println("Try again.")
		}
	}
}

func main() {
	fmt.Println("Welcome to UnoGO!")

	for {
		exit := mainMenu()
		if exit {
			break
		}
	}
}
