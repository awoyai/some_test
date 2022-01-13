package elliptic_curve

import (
	"math/big"
)

var Point0 = Point{x: big.NewInt(0), y: big.NewInt(0)}

type EllipticCurve struct {
	A, B, order *big.Int
}

func NewEllipticCurve(A, B, order *big.Int) EllipticCurve {
	return EllipticCurve{
		A:     A,
		B:     B,
		order: order,
	}
}

type Point struct {
	x, y *big.Int
}

func (p *Point) Equal(point Point) bool {
	return point.x.Cmp(p.x) == 0 && point.y.Cmp(p.y) == 0
}

func (e *EllipticCurve) Add(a, b Point) Point {
	switch {
	case e.Is0(a):
		return b
	case e.Is0(b):
		return a
	case e.MirrorByX(a, b):
		return Point0
	}

	x := big.NewInt(1)
	y := big.NewInt(1)
	k := e.CalculateK(a, b)

	x = x.Mul(x, pow(k, 2)).Sub(x, a.x).Sub(x, b.x).Mod(x, e.order)
	y = y.Mul(new(big.Int).Sub(a.x, x), k).Sub(y, a.y).Mod(y, e.order)
	return Point{x: x, y: y}
}

func (e *EllipticCurve) Mul(a Point, k int) Point {
	 res := Point0
	for i := 0; i < k; i++ {
		res = e.Add(res, a)
	}
	return res
}

func (e *EllipticCurve) Is0(p Point) bool {
	return p.x.Cmp(Point0.x) == 0 && p.y.Cmp(Point0.y) == 0
}

func (e *EllipticCurve) MirrorByX(a, b Point) bool {
	mirrorY := new(big.Int).Mod(new(big.Int).Mul(b.y, big.NewInt(-1)), e.order)
	return (a.x.Cmp(b.x) == 0) && (a.y.Cmp(mirrorY)) == 0
}

func (e *EllipticCurve) CalculateK(a, b Point) *big.Int {
	k := new(big.Rat)
	if a.Equal(b) {
		num := big.NewInt(3)
		num.Mul(num, pow(a.x, 2)).Add(num, e.A)
		denom := new(big.Int).Mul(big.NewInt(2), a.y)
		k.SetFrac(num, denom)
	} else {
		k.SetFrac(new(big.Int).Sub(b.y, a.y), new(big.Int).Sub(b.x, a.x))
	}
	return Mod(k, e.order)
}

func Mod(rat *big.Rat, p *big.Int) *big.Int {
	if rat.Denom().Int64() == 1 {
		return new(big.Int).Mod(rat.Num(), p)
	}
	res := big.NewInt(1)
	fastPow := FastPow(rat.Denom(), new(big.Int).Sub(p, big.NewInt(2)), p)
	res.Mod(rat.Num(), p).Mul(res, fastPow).Mod(res, p)
	return res
}

func (e *EllipticCurve) OnCurve(p Point) bool {
	// x^3 + ax + b
	x3 := pow(p.x, 3)
	x3AndB := new(big.Int).Add(x3, e.B)
	y := new(big.Int).Add(x3AndB, new(big.Int).Mul(e.A, p.x))
	modY := new(big.Int).Mod(y, e.order)
	y2 := new(big.Int).Mod(pow(p.y, 2), e.order)
	return modY.Cmp(y2) == 0
}

func pow(a *big.Int, p int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(0); i < p; i++ {
		res.Mul(res, a)
	}
	return res
}

func FastPow(a, b, p *big.Int) *big.Int {
	var res = big.NewInt(1)
	a.Mod(a, p)
	for b.Int64() != 0 {
		if b.Int64()&1 == 1 {
			res.Mul(res, a).Mod(res, p)
		}
		b = b.Rsh(b, 1)
		a.Mul(a, a).Mod(a, p)
	}
	return res
}
