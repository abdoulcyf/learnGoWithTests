package integers


import (
  "testing"
  "fmt")

  var _ fmt.Formatter// for debugging only,delete when done


func TestAdder(t *testing.T){
  //---------------------------------------------
  t.Run("add two integers", func(t *testing.T){
    got := Adder(1, 3)
    want := 4

    assertCorrectMessage(t, got, want)
  })

  //---------------------------------------------------
  t.Run("add even integers", func(t *testing.T){
    got := Adder(2, 3, 5, 4, 6)
    want := 12

    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want int) {
  t.Helper()
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}


// func ExampleAdder() {
 
//   sum := Adder(1, 5)
//   fmt.Println(sum)
//   // Output: 6
// }
