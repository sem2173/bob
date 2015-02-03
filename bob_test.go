package bob

import "testing"

func TestBob(t *testing.T) {
  tests := []struct {
    input string
    expected string
  } {
    {input: "True", expected: "Whatever"},
    {input: "True?", expected: "Sure"},
    {input: "True? ", expected: "Sure"},
    {input: "TRUE ", expected: "Whoa, chill out!"},
    {input: " ", expected: "Fine"},
    {input: "123", expected: "Whatever"},
    {input: "#~@@@", expected: "Whatever"},
    
  }

  for _, testCase := range tests {
    actual := Bob(testCase.input)
    expected := testCase.expected
    if actual != expected {
      t.Errorf("%s != %s", actual, expected)
    }
  }
}
