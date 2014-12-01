package montecarlo

import "sync"

func Launch(wg *sync.WaitGroup, bank_score *Score, player_score *Score) {

  deck := &Deck{}
  deck.Initialize()
  // Shuffle the deck
  deck.Shuffle()

  player_1 := &Player{}
  player_2 := &Player{}
  bank := &Bank{}

  player_1.Initialize()
  player_2.Initialize()
  bank.Initialize()

  bank.cards[0] = deck.cards[0]
  player_1.cards[0] = deck.cards[1]
  player_2.cards[0] = deck.cards[2]
  bank.cards[1] = deck.cards[3]
  player_1.cards[1] = deck.cards[4]
  player_2.cards[1] = deck.cards[5]
  bank.cards[2] = deck.cards[6]

  bank.score = bank.Calculate()
  player_1.score = player_1.Calculate()
  player_2.score = player_2.Calculate()

  if bank.score >= player_1.score || bank.score >= player_2.score {
    bank_score.Add(1)
  } else {
    player_score.Add(1)
  }

  wg.Done()
  return
}