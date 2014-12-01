package montecarlo

type Bank struct {
  cards []Card
  score int64
}

func (b *Bank) Initialize() {
  b.cards = make([]Card, 3)
}

func (b *Bank) Calculate() (x int64) {
  x = 0
  x += b.cards[0].Val
  x += b.cards[1].Val
  x += b.cards[2].Val
  if b.cards[0].Color == b.cards[1].Color && b.cards[0].Color == b.cards[2].Color {
    x += 35
  }
  if b.cards[0].Val == b.cards[1].Val && b.cards[0].Val == b.cards[2].Val {
    x += 999
  }
  return
}