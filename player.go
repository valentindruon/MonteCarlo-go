package montecarlo

type Player struct {
  cards []Card
  score int64
}

func (p *Player) Initialize() {
  p.cards = make([]Card, 2)
}

func (p *Player) Calculate() (x int64) {
  x = 0
  x += p.cards[0].Val
  x += p.cards[1].Val
  if p.cards[0].Color == p.cards[1].Color {
    x += 20
  }
  return
}