package maze

import "math/rand"

// Service interface exposes the generator service
type Service interface {
	Generate(width, height int, seed int64, algorithm Chooser) Maze
}

type service struct{}

// Middleware wraps the Service and performs actions before the implementation is called
type Middleware func(Service) Service

// New returns a Service implementation
func New(middleware ...Middleware) Service {
	svc := Service(service{})

	for _, m := range middleware {
		svc = m(svc)
	}

	return svc
}

func (s service) Generate(width, height int, seed int64, algorithm Chooser) Maze {
	return generate(width, height, rand.New(rand.NewSource(seed)), algorithm)
}
