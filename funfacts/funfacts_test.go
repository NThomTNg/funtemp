package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input
    want
	}
}


for _, tc := range tests {
  got := GetFunFacts(tc.input)
  if !reflect.DeepEqual(tc.want, got) {
    t.Errorf("expected: %v, got: %v", tc.want, got)
  }
}
/*
*
	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/


	// Her må du legge inn korrekte testverdier
	//tests := []test{
	//  {input: , want: },
	//}
