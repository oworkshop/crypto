package crypto_test

import (
	"fmt"
	"testing"

	"github.com/oworkshop/crypto"
)

func TestUUID(T *testing.T) {
	fmt.Println(crypto.NewDateUUID())
	fmt.Println(crypto.NewDateUUID())
	fmt.Println(crypto.NewDateUUID())
	fmt.Println(crypto.NewDateUUID())
	fmt.Println(crypto.NewDateUUID())
	fmt.Println(crypto.NewUUID())
	fmt.Println(crypto.NewUUID())
	fmt.Println(crypto.NewUUID())
	fmt.Println(crypto.NewUUID())
}
