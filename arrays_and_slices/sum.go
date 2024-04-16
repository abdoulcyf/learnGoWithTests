package arrays_and_slices


func Sum(arr []int)int{
  sum := 0

  for _, n := range arr {
    sum += n
  }
  return sum
}