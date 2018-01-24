package lm

import (
	"fmt"

	"github.com/barnex/fmath"
)

type Vec3 [3]float32

func (v *Vec3) Pointer() *[3]float32            { return (*[3]float32)(v) }
func (v *Vec3) Slice() []float32                { return v[:] }
func (v Vec3) X() float32                       { return v[0] }
func (v Vec3) Y() float32                       { return v[1] }
func (v Vec3) Z() float32                       { return v[2] }
func (v Vec3) XY() (float32, float32)           { return v[0], v[1] }
func (v Vec3) XZ() (float32, float32)           { return v[0], v[2] }
func (v Vec3) YX() (float32, float32)           { return v[1], v[0] }
func (v Vec3) YZ() (float32, float32)           { return v[1], v[2] }
func (v Vec3) ZX() (float32, float32)           { return v[2], v[0] }
func (v Vec3) ZY() (float32, float32)           { return v[2], v[1] }
func (v Vec3) XYZ() (float32, float32, float32) { return v[0], v[1], v[2] }
func (v Vec3) XYVec() Vec2                      { return Vec2{v[0], v[1]} }
func (v Vec3) XZVec() Vec2                      { return Vec2{v[0], v[2]} }
func (v Vec3) YXVec() Vec2                      { return Vec2{v[1], v[0]} }
func (v Vec3) YZVec() Vec2                      { return Vec2{v[1], v[2]} }
func (v Vec3) ZXVec() Vec2                      { return Vec2{v[2], v[0]} }
func (v Vec3) ZYVec() Vec2                      { return Vec2{v[2], v[1]} }

func (v *Vec3) SetX(x float32) { v[0] = x }
func (v *Vec3) SetY(y float32) { v[1] = y }
func (v *Vec3) SetZ(z float32) { v[2] = z }
func (v *Vec3) SetXYZ(x, y, z float32) {
	v[0] = x
	v[1] = y
	v[2] = z
	return
}

func (v Vec3) String() string {
	return fmt.Sprintf("[%f,%f,%f]", v[0], v[1], v[2])
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec3) Mul(v2 Vec3) Vec3 {
	return Vec3{v1[0] * v2[0], v1[1] * v2[1], v1[2] * v2[2]}
}

func (v Vec3) MulF(f float32) Vec3 {
	return Vec3{v[0] * f, v[1] * f, v[2] * f}
}

func (v Vec3) MulMat4x4(m Mat4x4) Vec3 {
	v, w := v.MulMat4x4W(m)
	return v.DivF(w)
}

func (v Vec3) MulMat4x4W(m Mat4x4) (Vec3, float32) {
	return Vec3{
			m[0]*v[0] + m[4]*v[1] + m[8]*v[2] + m[12],
			m[1]*v[0] + m[5]*v[1] + m[9]*v[2] + m[13],
			m[2]*v[0] + m[6]*v[1] + m[10]*v[2] + m[14]},
		m[3]*v[0] + m[7]*v[1] + m[11]*v[2] + m[15]
}

func (v Vec3) MulQuat(q Quat) Vec3 {
	return q.MulQuat(Quat{v[0], v[1], v[2], 0}).MulQuat(q.Conjugate()).XYZVec()
}

func (v1 Vec3) Div(v2 Vec3) Vec3 {
	return Vec3{v1[0] / v2[0], v1[1] / v2[1], v1[2] / v2[2]}
}

func (v Vec3) DivF(f float32) Vec3 {
	return Vec3{v[0] / f, v[1] / f, v[2] / f}
}

func (v1 Vec3) Dot(v2 Vec3) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0]}
}

func (v Vec3) Len() float32 {
	return fmath.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v Vec3) LenSqr() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v Vec3) Norm() Vec3 {
	return v.MulF(1.0 / v.Len())
}

func (v Vec3) AngleTo(v2 Vec3) float32 {
	return fmath.Acos(v.Norm().Dot(v2.Norm()))
}
