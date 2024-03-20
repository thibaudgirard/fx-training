package repository

import "math/rand"

type GreetRepository struct {
}

func NewGreetRepository() *GreetRepository {
	return &GreetRepository{}
}

func (r *GreetRepository) Greet() string {
	return []string{
		"Hello world!",
		"Hello you!",
		"Hi!",
	}[rand.Intn(3)]
}
