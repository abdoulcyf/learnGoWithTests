package integers



func Adder(num... int) int {
  if len(num) == 2 {
    return num[0] + num[1]
  }

  sum := 0
  for _, n := range num{
    if n % 2 == 0{
      sum += n
    }
    }
   return sum
}