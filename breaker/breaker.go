package breaker

type Smasher interface {
	Smash() (Broken, error)
}

type Broken bool

type Breaker struct {
	Smasher Smasher
}

func (b Breaker) Break() (Broken, error) {
	return b.Smasher.Smash()
}
