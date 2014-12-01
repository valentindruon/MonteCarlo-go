package montecarlo

import "math/rand"
import "time"

type Deck struct {
  cards []Card
}

func (d *Deck) Initialize() {

  d.cards = make([]Card, 52)
  // Fill the deck
  for i := 1; i <= 13; i++ {
    for j := 1; j <= 4; j++ {
      c := Card{ Val: int64(i), Color: int64(j)}
      d.cards = append(d.cards, c)
    }
  }
}

func (d *Deck) Shuffle() {
    r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
    for i := range d.cards {
        j := r.Intn(len(d.cards))
        d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
    }
}