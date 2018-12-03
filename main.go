package main

import (
	"strconv"
	"strings"
)

func main() {

}

type Token struct {
	op  string
	val float64
}

type Parser struct {
}

func NewParser() *Parser { return &Parser{} }
func (p *Parser) Evaluate(s string) (float64, error) {
	s = strings.Replace(s, " ", "", -1)
	n, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return n, nil
	}

	plus := strings.Index(s, "+")
	if plus > -1 {
		left := s[:plus]
		right := s[plus+1:]

		lv, err := p.Evaluate(left)
		if err != nil {
			return 0.0, err
		}
		rv, err := p.Evaluate(right)
		if err != nil {
			return 0.0, err
		}

		return lv + rv, nil

	}

	minus := strings.Index(s, "-")
	if minus > -1 {
		left := s[:minus]
		right := s[minus+1:]

		lv, err := p.Evaluate(left)
		if err != nil {
			return 0.0, err
		}
		rv, err := p.Evaluate(right)
		if err != nil {
			return 0.0, err
		}

		return lv - rv, nil

	}

	mult := strings.Index(s, "*")
	div := strings.Index(s, "/")
	if div > mult {
		left := s[:div]
		right := s[div+1:]

		lv, err := p.Evaluate(left)
		if err != nil {
			return 0.0, err
		}
		rv, err := p.Evaluate(right)
		if err != nil {
			return 0.0, err
		}

		return lv / rv, nil

	} else if mult > div {
		left := s[:mult]
		right := s[mult+1:]

		lv, err := p.Evaluate(left)
		if err != nil {
			return 0.0, err
		}
		rv, err := p.Evaluate(right)
		if err != nil {
			return 0.0, err
		}

		return lv * rv, nil
	}
	return n, err
}
