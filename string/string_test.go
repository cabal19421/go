package string

import "testing"

func Test(t *testing.T) {
  var tests = []struct {
    s, want string
    }{
      {"backward", "drawkcab"},
      {"", ""},
    }
    for _, c := range tests {
      got := Reverse(c.s)
      if got != c.want {
        t.Errorf("in correct reverse")
      }
    }
}