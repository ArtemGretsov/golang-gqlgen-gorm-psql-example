package calcutil

import "fmt"

func CalculateDifference(cur float64, prev float64) string {
  return fmt.Sprintf("%f", cur / prev * 100 - 100) + "%"
}