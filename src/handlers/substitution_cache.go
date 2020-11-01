package handlers

import (
	"errors"
	"fmt"

	"github.com/NathanBak/nptl/src/runes"
)

const substitutionRune rune = '\uEEEE'

type SubstitutionCache struct {
	queue []runes.Runes
}

func (c *SubstitutionCache) ReplaceAll(rs runes.Runes) (runes.Runes, error) {
	out := runes.Runes{}

	for _, r := range rs {
		if r == substitutionRune {
			if len(c.queue) < 1 {
				return out, errors.New("cache does not contain substitution value")
			}
			out = append(out, c.queue[0]...)
			c.queue[0] = nil
			c.queue = c.queue[1:]
		} else {
			out = append(out, r)
		}
	}

	if len(c.queue) > 0 {
		return out, errors.New("cache contains extra values")
	}

	return out, nil
}

func (c *SubstitutionCache) Substitute(rs, search, replace runes.Runes) (runes.Runes, error) {
	out := runes.Runes{}

	if len(search) < 1 {
		return out, errors.New("search Runes cannot be empty")
	}

	found := false
	for i := 0; i < len(rs); i++ {

		fmt.Println(rs.String())
		for k := 0; k < i; k++ {
			fmt.Print(" ")
		}
		fmt.Println(search.String())
		fmt.Println("")

		match := true
		if !found {
			for j, s := range search {
				if s != rs[i+j] {
					match = false
					break
				}
			}
		} else {
			match = false
		}

		if match {
			out = append(out, substitutionRune)
			c.queue = append(c.queue, replace)
			i += len(search) - 1
			found = true
		} else {
			out = append(out, rs[i])
		}
	}

	if !found {
		return out, fmt.Errorf("unable to find [%v] in [%v]", search, rs)
	}
	return out, nil
}

func (c *SubstitutionCache) SubstituteStrings(rs runes.Runes, search, replace string) (runes.Runes, error) {
	s, err := runes.FromString(search)
	if err != nil {
		return nil, err
	}
	r, err := runes.FromString(replace)
	if err != nil {
		return nil, err
	}
	return c.Substitute(rs, s, r)
}
