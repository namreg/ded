package ded

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

//ErrEmailMalformed error returns when email is not valid.
var ErrEmailMalformed = errors.New("ded: email malformed")

//IsDisposableEmail checks whether the given email is disposable.
func IsDisposableEmail(email string) (bool, error) {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 {
		return false, ErrEmailMalformed
	}
	return IsDisposableDomain(parts[1])
}

//IsDisposableDomain checks whether the given domain is disposable.
func IsDisposableDomain(domain string) (bool, error) {
	domains.once.Do(func() { domains.loadFromFile("domains.txt") })
	if domains.err != nil {
		return false, domains.err
	}
	domain = strings.TrimSpace(domain)
	return domains.has(normalizeString(domain)), nil
}

func normalizeString(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

var domains = new(collection)

type collection struct {
	items map[string]struct{}
	err   error
	once  sync.Once
}

func (c *collection) has(item string) bool {
	_, ok := c.items[item]
	return ok
}

func (c *collection) loadFromFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		c.err = fmt.Errorf("ded: could not open %q file: %v", filename, err)
		return
	}
	c.items = make(map[string]struct{})

	s := bufio.NewScanner(f)
	for s.Scan() {
		domain := normalizeString(s.Text())
		if domain == "" {
			continue
		}
		c.items[domain] = struct{}{}
	}

	if err := s.Err(); err != nil {
		c.err = fmt.Errorf("ded: scanner error: %v", err)
	}
}
