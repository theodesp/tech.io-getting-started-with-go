package main
import "testing"
import "fmt"
import "./universe"

func success(b bool) {
  fmt.Printf("TECHIO> success %v\n",b)
}

func msg(channel string, msg string) {
  fmt.Printf("TECHIO> message --channel \"%v\" \"%v\"\n", channel, msg)
}

func TestUniverse(t *testing.T) {
  var s = universe.CountAllStars(800,120,450)
  if s != 1370 {
    t.Error("Expected 1370, got ", s)
    success(false)
    msg("Oops! 🐞", "Have you properly accumulated each value into 'total' 🤔")
  } else {
    msg("Kudos 🌟", "You've passed a test!")
    success(true)
  }
}