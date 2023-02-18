package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/HomayoonAlimohammadi/splitshare/member"
)

var saveSuffix string = ".json"

type Event struct {
	mu      sync.Mutex
	Members []*member.Member
}

func New() *Event {
	return &Event{}
}

// Given the amount and the spender, calculate other members shares
// and adds it to their debts to the spender
func (e *Event) MemberSpent(spender string, amount float64) {
	share := amount / float64(len(e.Members))
	for _, mem := range e.Members {
		if mem.Name != spender {
			mem.Borrow(spender, share)
		}
	}
}

func (e *Event) SaveToDisk(fileName string) error {

	e.mu.Lock()
	defer e.mu.Unlock()

	// open file
	f, err := os.Create(fileName + saveSuffix)
	if err != nil {
		return err
	}
	defer f.Close()

	// marshal (encode) the content
	b, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	_, err = io.Copy(f, r)

	return err
}

func (e *Event) LoadFromDisk(fileName string) (*Event, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// open file
	f, err := os.Open(fileName + saveSuffix)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// decode the content
	err = json.NewDecoder(f).Decode(e)

	return e, err
}

// Logs the result of the event.
//
// Only use when the event is ready and the records are completely given.
func (e *Event) Describe() {
	fmt.Println("##############################")
	fmt.Println("Event Description:")
	for _, m := range e.Members {
		fmt.Println("-------------------")
		fmt.Println(m.Name)
		for otherName := range m.Debts {
			otherMem, found := e.getMemberByName(otherName)
			if !found {
				fmt.Printf("failed to get member %s \n", otherName)
			}
			debt := m.Owes(otherMem)
			if debt > 0 {
				fmt.Printf("Should Give %f to `%s` \n", debt, otherName)
			}
		}
		fmt.Println("-------------------")
	}
	fmt.Println("##############################")
}

// Checks if the event contains a member given his/her name
func (e *Event) HasMember(name string) bool {
	for _, m := range e.Members {
		if m.Name == name {
			return true
		}
	}
	return false
}

// Returns a member of the event given his/her name. Also alerts if the member was
// not found.
func (e *Event) getMemberByName(name string) (*member.Member, bool) {
	for _, m := range e.Members {
		if m.Name == name {
			return m, true
		}
	}
	return nil, false
}
