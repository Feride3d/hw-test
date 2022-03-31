package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result, err := countDomains(r, "."+domain)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return result, nil
}

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	user := &User{}
	scanner := bufio.NewScanner(r)
	result := make(DomainStat)
	for scanner.Scan() {
		*user = User{}
		if errors.Is(nil, io.EOF) {
			break
		}
		if err := jsoniter.Unmarshal(scanner.Bytes(), user); err != nil {
			return nil, err
		}
		if strings.Contains(user.Email, "."+domain) {
			result[strings.ToLower(strings.Split(user.Email, "@")[1])]++
		}
	}
	return result, nil
}
