package main

// VARIABLES
// // ! you CANT have initialize variables outside of the function body
// func main() {
// 	// var card string = "Ace of Spades"
// 	// card := "Ace of Spades"; // := used only to initialize a variable the first time
// 							 // to re-assign it use =

// 	card := newCard();

// 	fmt.Println(card);
// }

func newCard() string {
	return "Five of Diamonds";
}

// // SLICES/ ARRAYS
// func main() {
// 	// slice where each element is of type int: []int{}

// 	cards := deck{
// 		"Ace of Diamonds",
// 		newCard(),
// 	}

// 	cards = append(cards, "Meow")

// 	// for i := 0; i < len(cards); i++ {
// 	// 	fmt.Println(cards[i]);
// 	// }

// 	// for i, card := range cards {
// 	// 	fmt.Println(i, card);
// 	// }

	
// 	deck.print(cards);
// 	cards.print();



// }


// func main() {
	
// 	newDeck := newDeck();
// 	newDeck.print();

// }

// func main() {
// 		cards := newDeck();

// 		hand, remainingDeck := deal(cards, 5)
// 		fmt.Println("Hand")
// 		hand.print()

// 		fmt.Println("Remaining Deck")
// 		remainingDeck.print()
	
// 	}

// FROM ARRAY TO BYTE SLICE
// func main() {
// 	gretting := "Hi there!"
// 	fmt.Println([]byte(gretting))
// }


func main() {
	// cards := newDeck();
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.shuffle();

	cards.print();
}