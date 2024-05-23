package main

import (
	"fmt"

	"github.com/ichiban/prolog"
)

type Quest struct {
	F int64
}

func Query(p *prolog.Interpreter, query string) {
	sols, err := p.Query(query)
	if err != nil {
		panic(err)
	}
	defer sols.Close()

	var q Quest

	fmt.Printf("? %s -->\n", query)
	for sols.Next() {
		if err := sols.Scan(&q); err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", q.F)
	}

	if err := sols.Err(); err != nil {
		panic(err)
	}
}

func main() {

	p := prolog.New(nil, nil)

	if err := p.Exec(`
		/** factorial in Prolog */
		fact(0, 1).
		fact(N, F) :-
			N > 0,
			N1 is N - 1,
			fact(N1, F1),
			F is N * F1.
	`); err != nil {
		panic(err)
	}

	Query(p, `fact(5, F).`)
	Query(p, `fact(6, F).`)

}
