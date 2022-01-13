package elliptic_curve

import (
	"fmt"
	"math/big"
	"testing"
)

func TestINV(t *testing.T) {
	a := big.NewInt(3)
	b := big.NewInt(6)
	p := big.NewInt(97)
	fastPow := FastPow(a,b,p)
	fmt.Printf("%v", fastPow)
}
