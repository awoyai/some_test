package elliptic_curve

import "math/big"

const MOD = 23

func RatMod(a *big.Rat) *big.Int {
	p := big.NewInt(MOD)
	return new(big.Int).Mod(new(big.Int).Mul(FastPow(a.Denom(), big.NewInt(MOD-2), p), new(big.Int).Mod(a.Num(), p)), p)
}

