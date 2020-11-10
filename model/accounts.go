package model

import (
	"time"

	"github.com/textileio/go-threads/core/thread"
)

type Member struct {
	Key      thread.PubKey
	Username string
	Role     Role
}

type Account struct {
	Type             AccountType
	Key              thread.PubKey
	Secret           thread.Identity
	Name             string
	Username         string
	Email            string
	Token            thread.Token
	Members          []Member
	BucketsTotalSize int64
	CreatedAt        time.Time
	PowInfo          *PowInfo
}

type AccountType int

type Role int

const (
	Dev AccountType = iota
	Org
)

const (
	OrgOwner Role = iota
	OrgMember
)

func (r Role) String() (s string) {
	switch r {
	case OrgOwner:
		s = "owner"
	case OrgMember:
		s = "member"
	}
	return
}
