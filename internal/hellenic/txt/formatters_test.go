package txt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_boolToYesNoFlag(t *testing.T) {
	tests := []struct {
		name  string
		value bool
		want  string
	}{
		{"Y", true, "Y"},
		{"N", false, "N"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, boolToYesNoFlag(tt.value))
		})
	}
}

func Test_limit(t *testing.T) {
	type args struct {
		value string
		lim   uint16
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"cut 35", args{"2/4 Letnikovskaya str., Moscow, 115114", 35}, "2/4 Letnikovskaya str., Moscow, 115"},
		{"cut more 1", args{"abcdef", 3}, "abc"},
		{"cut more 2", args{"abcdef", 10}, "abcdef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, limit(tt.args.value, tt.args.lim))
		})
	}
}

func Test_underscore(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"has value", "some value", "some value"},
		{"no value", "", "_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, underscore(tt.value))
		})
	}
}

func Test_lPadSp(t *testing.T) {
	tests := []struct {
		name  string
		value string
		lim   int
		want  string
	}{
		{"empty pad 5", "", 5, "     "},
		{"less pad 5", "abc", 5, "  abc"},
		{"eq pad 5", "abcde", 5, "abcde"},
		{"more pad 5", "abcdefgh", 5, "abcdefgh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lPadSp(tt.value, tt.lim))
		})
	}
}

func Test_rPadSp(t *testing.T) {
	tests := []struct {
		name  string
		value string
		lim   int
		want  string
	}{
		{"empty pad 5", "", 5, "     "},
		{"less pad 5", "abc", 5, "abc  "},
		{"eq pad 5", "abcde", 5, "abcde"},
		{"more pad 5", "abcdefgh", 5, "abcdefgh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rPadSp(tt.value, tt.lim))
		})
	}
}
