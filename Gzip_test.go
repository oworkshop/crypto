package crypto_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oworkshop/crypto"
)

func TestGzip(T *testing.T) {
	fmt.Println("TestGzip")
	text := strings.Repeat("12345678901234567890", 300)
	ctext, err := crypto.GzipCompress(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ctext)
	dtext, err := crypto.GzipDecompress(ctext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dtext)
}
