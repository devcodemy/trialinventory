package main

import (
	"log"
	"time"
	"trial/data"
	"trial/tmp"
	"trial/tokentype"
	"trial/usertype"
)

func main() {
	// log.Printf("Some new %s", time.Date().Now())
	log.Printf("Some new %s", time.Now())
	tk := data.Token{
		Type:  tokentype.PaaS,
		Name:  "Test token",
		Value: "SOME_TOKEN_HASH",
	}

	tk2 := data.Token{
		Type:  tokentype.Cluster,
		Name:  "Test token for Cluster",
		Value: "SOME_TOKEN_HASH",
	}

	admin := usertype.User{
		UserId:       0,
		Name:         "Adam",
		Login:        "admin",
		Email:        "admin@company.com",
		PasswordHash: "SOME_PASS_HASH",
		Role:         usertype.Admin,
	}
	user1 := usertype.User{
		UserId:       1,
		Name:         "Alice",
		Login:        "alice",
		Email:        "alice@company.com",
		PasswordHash: "SOME_PASS_HASH",
		Role:         usertype.ReadOnly,
	}

	log.Printf("Paas: %+v", tk)
	log.Printf("Admin user: %+v", admin)
	log.Printf("ReadOnly user: %+v", user1)

	log.Printf("Paas: %+v", tk2)

	be := tmp.Behavior{
		Name:   "Bob",
		Action: tmp.SomeWhere,
	}
	log.Printf("TMP: %+v", be)
}
