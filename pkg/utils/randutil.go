package utils

import "golang.org/x/exp/rand"

type Rand struct {
	Seed  uint64
	Start int64
	End   int64
}

func CreateRandFactory(seed uint64, start int64, end int64) *Rand {
	return &Rand{Seed: seed, Start: start, End: end}
}

func (r *Rand) GetRandInt() int64 {
	rand.Seed(r.Seed)
	num := rand.Int63n(r.End-r.Start+1) + r.Start
	return num
}
