package id_generator

// type idGenerator interface {
// 	GenerateId() int32
// }

type idGenerator struct {
	counter     int
	generate    chan struct{}
	current     chan struct{}
	getGenerate chan int
	getCurrent  chan int
}

func (g *idGenerator) generator() int {
	for {
		select {
		case <-g.generate:
			g.counter += 1
			g.getGenerate <- g.counter
		case <-g.current:
			g.getCurrent <- g.counter
		}
	}
}

func (g *idGenerator) Generate() int {
	g.generate <- struct{}{}
	return <- g.getGenerate
}

func (g *idGenerator) Current() int {
	g.current <- struct{}{}
	return <- g.getCurrent
}

func newIdGenerator() *idGenerator {
	g := idGenerator{
		generate:    make(chan struct{}),
		current:     make(chan struct{}),
		getGenerate: make(chan int),
		getCurrent:  make(chan int),
	}
	go g.generator()
	return &g
}

// Generator generator
var Generator = newIdGenerator()
