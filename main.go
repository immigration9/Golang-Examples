package main

import "fmt"

func main() {
	// var card string = "Ace of Spaces"
	// card := "Ace of Spades"
	// card := newCard()

	// cards := []string{newCard(), newCard()}
	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "Six of Spades")
	cards := newDeck()
	fmt.Println(cards.toString())
	// Go Compiler 가 type을 detect하도록
	// 새로운 변수 만들 때만 :=을 사용한다. reassign할 때는 필요 없다.
	hand, remainingCards := deal(cards, 5)

	// fmt.Println(cards)
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards.print()
	fmt.Println("==========")
	hand.print()
	fmt.Println("-----")
	remainingCards.print()

	cards.saveToFile("my_cards")

	fmt.Println(newDeckFromFile("my_cards"))

}

// needs to have return type
func newCard() string {
	return "Five of Diamonds"
}
