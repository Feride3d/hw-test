package hw10programoptimization

import (
	"bufio"
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
	res, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(res, domain)
}

type users []User

func getUsers(r io.Reader) (users, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	scanner := bufio.NewScanner(r)
	i := 0
	user := User{}
	result := make(users, 0, 100_000)
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return result, err
		}
		result = append(result, user)
		i++
	}

	return result, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	for _, user := range u {
		if strings.Contains(user.Email, "."+domain) {
			matchDomain := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			num := result[matchDomain]
			num++
			result[matchDomain] = num
		}
	}
	return result, nil
}
