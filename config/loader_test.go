package config

import (
	"fmt"
	env "restfullApi/util/environment"
	"testing"
)

type test struct {
	name string
	input
	output string
	hasErr bool
}
type input struct {
	key string
	ext string
	enc bool
}

func TestNewEncrypted(t *testing.T) {
	tests := []test{
		{
			name:   "encrypted load test",
			input:  input{"", "json", true},
			output: "",
			hasErr: false,
		},
	}

	for _, tst := range tests {
		_, err := New("AnKoloft@~delNazok!12345", "json", true)
		if tst.hasErr && err == nil {
			t.Errorf("test %s failed, expected error but got nil", tst.name)
		}
	}
}

func TestNewPlane(t *testing.T) {
	tests := []test{
		{
			name:   "plane load test-1",
			input:  input{"kfjbrewjfv", "json", false},
			output: "",
			hasErr: true,
		},
		{
			name:   "plane load test-2",
			input:  input{"", "txt", false},
			output: "",
			hasErr: true,
		},
		{
			name:   "plane load test-3",
			input:  input{"", "json", false},
			output: "",
			hasErr: false,
		},
	}
	for _, tst := range tests {
		_, err := New(tst.input.key, tst.input.ext, tst.input.enc)
		if tst.hasErr && err == nil {
			t.Errorf("test %s failed, expected error but got nil", tst.name)
		}
	}
}

func TestFile_DbConfig(t *testing.T) {
	_ = env.SetEnv("environment", "local")
	tests := []test{
		{
			name:   "test get database config",
			input:  input{},
			output: "",
			hasErr: false,
		},
	}
	l, err := New("", "json", false)
	if err != nil {
		t.Errorf("test failed, expected error but got nil")
	}
	for _, tst := range tests {
		_, cErr := l.DbConfig()
		if cErr != nil {
			t.Errorf("test %s failed, expected error but got nil", tst.name)
		}
	}
}
func TestFile_GetRootAdmin(t *testing.T) {
	tests := []test{
		{
			name:   "test loader",
			input:  input{"", "json", true},
			output: "",
			hasErr: false,
		},
	}
	for _, tst := range tests {
		l, err := New("", "", false)
		if tst.hasErr && err == nil {
			t.Errorf("test %s failed, expected error but got nil", tst.name)
		}
		fmt.Println(l.DbConfig())
	}
}
