package crypto_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oworkshop/crypto"
)

func TestSnappy(T *testing.T) {
	fmt.Println("TestSnappy")
	text := strings.Repeat("12345678901234567890", 300)
	ctext := crypto.SnappyCompress(text)
	fmt.Println(ctext)
	dtext, _ := crypto.SnappyDecompress(ctext)
	fmt.Println(dtext)
}
