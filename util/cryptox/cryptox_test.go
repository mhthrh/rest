package cryptox_test

import (
	"log"
	cryptx "restfullApi/util/cryptox"
	"testing"
)

type test struct {
	name   string
	input  string
	output string
}

var (
	key = "XXKoloft@~delNazok!12345"
	c   cryptx.ICrypto
	err error
)

func init() {
	c, err = cryptx.New(key)
	if err != nil {
		log.Fatal(err)
	}
}
func TestEncrypt(t *testing.T) {
	tests := []test{
		{
			name:   "test-1",
			input:  "this is a test",
			output: "PsGGcOGdZ7Cm7ENYRayq2yKGwZ0cp8+hLsim6fgfXhkWRi5RdTSGNq7a",
		},
	}
	for _, tst := range tests {
		r, err := c.Encrypt(tst.input)
		if err != nil {
			t.Error(err)
		}
		if r != tst.output {
			t.Error("output is incorrect")
		}
	}
}

func TestDecrypt(t *testing.T) {
	tests := []test{
		{
			name:   "test-2",
			output: "this is a test",
			input:  "PsGGcOGdZ7Cm7ENYRayq2yKGwZ0cp8+hLsim6fgfXhkWRi5RdTSGNq7a",
		},
	}
	for _, tst := range tests {
		r, err := c.Decrypt(tst.input)
		if err != nil {
			t.Error(err)
		}
		if r != tst.output {
			t.Error("output is incorrect")
		}
	}
}
func TestCrypto_Sha256(t *testing.T) {
	tests := []test{
		{
			name:   "test-3",
			output: "this is a test",
			input:  "49768547e748662cf7883eb00c8129e71261fda61638843184145bea15810d24",
		},
	}
	for _, tst := range tests {
		r := c.Sha256(tst.input)

		if r != tst.output {
			t.Error("output is incorrect")
		}
	}
}
