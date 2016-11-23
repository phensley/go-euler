package euler

import (
	"fmt"
	"log"
	"sort"
	"time"
)

var (
	// Debug emits debugging output
	Debug = false

	// Verbose emits more output
	Verbose = false

	// Mapping of problem id to solution
	solutions = make(map[string]*solution)
)

type solution struct {
	id          string
	description string
	solve       func(*Context)
}

func sortSolutions() []string {
	ids := []string{}
	for id := range solutions {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

// ShowSolutions displays a list of all solutions registered
func ShowSolutions() {
	fmt.Println("Solutions:")
	for _, id := range sortSolutions() {
		fmt.Printf("  %s  %s\n", id, solutions[id].description)
	}
}

// Register a solution for a given problem id
func Register(id, description string, solve func(ctx *Context)) {
	if _, exists := solutions[id]; exists {
		log.Fatalf("Attempt to register solution '%s' more than once!\n", id)
	}
	solutions[id] = &solution{id, description, solve}
}

// Solve a given problem by id
func Solve(id string) *Context {
	ctx := &Context{id: id}
	solution, exists := solutions[id]
	if !exists {
		return ctx
	}

	known := answers[id]
	if known != "" {
		known = RevealAnswer(id, known)
	}
	ctx.exists = true
	ctx.description = solution.description
	ctx.known = known

	start := time.Now()
	solution.solve(ctx)
	ctx.elapsed = time.Now().Sub(start)
	return ctx
}

// SolveAll executes all solutions
func SolveAll() <-chan *Context {
	ch := make(chan *Context)
	go func() {
		solutions := sortSolutions()
		for _, id := range solutions {
			ch <- Solve(id)
		}
		close(ch)
	}()
	return ch
}
