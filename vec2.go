package lm

import (
	"fmt"
	"github.com/void6/math32"
)

type Vec2 [2]float32

func (v *Vec2) Pointer() *[2]float32 { return (*[2]float32)(v) }
func (v *Vec2) Slice() []float32     { return v[:] }
func (v Vec2) X() float32            { return v[0] }
func (v Vec2) Y() float32            { return v[1] }
func (v Vec2) XY() (x, y float32) {
	x = v[0]
	y = v[1]
	return
}

func (v *Vec2) SetX(x float32) { v[0] = x }
func (v *Vec2) SetY(y float32) { v[1] = y }
func (v *Vec2) SetXY(x, y float32) {
	v[0] = x
	v[1] = y
	return
}

func (v Vec2) String() string {
	return fmt.Sprintf("[%f,%f]", v[0], v[1])
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec2) Mul(v2 Vec2) Vec2 {
	return Vec2{v1[0] * v2[0], v1[1] * v2[1]}
}

func (v Vec2) MulF(f float32) Vec2 {
	return Vec2{v[0] * f, v[1] * f}
}

func (v Vec2) MulMat3x3(m Mat3x3) (Vec2, float32) {
	return Vec2{
			m[0]*v[0] + m[1]*v[1] + m[2],
			m[3]*v[0] + m[4]*v[1] + m[5]},
		m[6]*v[0] + m[7]*v[1] + m[8]
}

func (v1 Vec2) Div(v2 Vec2) Vec2 {
	return Vec2{v1[0] / v2[0], v1[1] / v2[1]}
}

func (v Vec2) DivF(f float32) Vec2 {
	return Vec2{v[0] / f, v[1] / f}
}

func (v1 Vec2) Dot(v2 Vec2) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v Vec2) Len() float32 {
	return math32.Sqrt(v[0]*v[0] + v[1]*v[1])
}

func (v Vec2) LenSqr() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

func (v Vec2) Norm() Vec2 {
	return v.MulF(1.0 / v.Len())
}

func (v Vec2) Atan2() float32 {
	return math32.Atan2(v[1], v[0])
}

func (v Vec2) AngleTo(v2 Vec2) float32 {
	return math32.Acos(v.Norm().Dot(v2.Norm()))
}
