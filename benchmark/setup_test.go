package benchmark

import (
	"errors"
	"fmt"
	"time"
)

var (
	errExample = errors.New("fail")

	_messages   = fakeMessages(1000)
	_tenInts    = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	_tenStrings = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	_tenTimes   = []time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(2, 0),
		time.Unix(3, 0),
		time.Unix(4, 0),
		time.Unix(5, 0),
		time.Unix(6, 0),
		time.Unix(7, 0),
		time.Unix(8, 0),
		time.Unix(9, 0),
	}
	_oneUser = &user{
		Name:      "Jane Doe",
		Email:     "jane@test.com",
		CreatedAt: time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC),
	}
	_tenUsers = users{
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
	}
)

func fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func getMessage(iter int) string {
	return _messages[iter%1000]
}

type users []*user

type user struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ctxBodyBytes     = 123456789
	ctxRequest       = "GET /icons/ubuntu-logo.png HTTP/1.1"
	ctxTimeElapsedMs = 11.398466
	ctxTime          = time.Now()
	ctxMonths        = []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	ctxFirst10Primes = []int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31}
	ctxErr           = errors.New("failed to open file: /home/dev/new.txt")
)
