package lm

import (
	"fmt"
	"github.com/void6/math32"
)

type Mat4x4 [16]float32

func (m *Mat4x4) Pointer() *[16]float32 { return (*[16]float32)(m) }
func (m *Mat4x4) Slice() []float32      { return m[:] }

func (m *Mat4x4) String() string {
	return fmt.Sprintf("[%f,%f,%f,%f,\n %f,%f,%f,%f,\n %f,%f,%f,%f,\n %f,%f,%f,%f]",
		m[0], m[4], m[8], m[12], m[1], m[5], m[9], m[13], m[2], m[6], m[10], m[14], m[3], m[7], m[11], m[15])
}

func Mat4x4Identity() Mat4x4 {
	return Mat4x4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

func Mat4x4Translate(v Vec3) Mat4x4 {
	return Mat4x4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		v[0], v[1], v[2], 1}
}

func Mat4x4Scale(v Vec3) Mat4x4 {
	return Mat4x4{
		v[0], 0, 0, 0,
		0, v[1], 0, 0,
		0, 0, v[2], 0,
		0, 0, 0, 1}
}

func Mat4x4Rotate(axis Vec3, radians float32) Mat4x4 {
	x, y, z := axis.Norm().XYZ()
	c := math32.Cos(radians)
	s := math32.Sin(radians)

	return Mat4x4{
		x*x*(1-c) + c, y*x*(1-c) + z*s, x*z*(1-c) - y*s, 0,
		x*y*(1-c) - z*s, y*y*(1-c) + c, y*z*(1-c) + x*s, 0,
		x*z*(1-c) + y*s, y*z*(1-c) - x*s, z*z*(1-c) + c, 0,
		0, 0, 0, 1}
}

func Mat4x4LookAt(eye, center, up Vec3) Mat4x4 {
	f := center.Sub(eye).Norm()
	s := f.Cross(up.Norm())
	u := s.Cross(f)
	return Mat4x4{
		s[0], u[0], -f[0], 0,
		s[1], u[1], -f[1], 0,
		s[2], u[2], -f[2], 0,
		-s.Dot(eye), -u.Dot(eye), f.Dot(eye), 1}
}

func Mat4x4Frustum(left, right, bottom, top, zNear, zFar float32) Mat4x4 {
	width := right - left
	height := top - bottom
	depth := zFar - zNear
	return Mat4x4{
		(zNear * 2.0) / width, 0, 0, 0,
		0, (zNear * 2.0) / height, 0, 0,
		(left + right) / width, (bottom + top) / height, -(zNear + zFar) / depth, -1,
		0, 0, -(zNear * zFar * 2.0) / depth, 0}
}

func Mat4x4Perspective(fovY, aspect, zNear, zFar float32) Mat4x4 {
	f := float32(1.0 / math32.Tan(fovY/2.0))
	d := zNear - zFar
	return Mat4x4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (zFar + zNear) / d, -1,
		0, 0, (2 * zFar * zNear) / d, 0}
}

func (m1 Mat4x4) Mul(m2 Mat4x4) Mat4x4 {
	return Mat4x4{
		m1[0]*m2[0] + m1[1]*m2[4] + m1[2]*m2[8] + m1[3]*m2[12],
		m1[0]*m2[1] + m1[1]*m2[5] + m1[2]*m2[9] + m1[3]*m2[13],
		m1[0]*m2[2] + m1[1]*m2[6] + m1[2]*m2[10] + m1[3]*m2[14],
		m1[0]*m2[3] + m1[1]*m2[7] + m1[2]*m2[11] + m1[3]*m2[15],
		m1[4]*m2[0] + m1[5]*m2[4] + m1[6]*m2[8] + m1[7]*m2[12],
		m1[4]*m2[1] + m1[5]*m2[5] + m1[6]*m2[9] + m1[7]*m2[13],
		m1[4]*m2[2] + m1[5]*m2[6] + m1[6]*m2[10] + m1[7]*m2[14],
		m1[4]*m2[3] + m1[5]*m2[7] + m1[6]*m2[11] + m1[7]*m2[15],
		m1[8]*m2[0] + m1[9]*m2[4] + m1[10]*m2[8] + m1[11]*m2[12],
		m1[8]*m2[1] + m1[9]*m2[5] + m1[10]*m2[9] + m1[11]*m2[13],
		m1[8]*m2[2] + m1[9]*m2[6] + m1[10]*m2[10] + m1[11]*m2[14],
		m1[8]*m2[3] + m1[9]*m2[7] + m1[10]*m2[11] + m1[11]*m2[15],
		m1[12]*m2[0] + m1[13]*m2[4] + m1[14]*m2[8] + m1[15]*m2[12],
		m1[12]*m2[1] + m1[13]*m2[5] + m1[14]*m2[9] + m1[15]*m2[13],
		m1[12]*m2[2] + m1[13]*m2[6] + m1[14]*m2[10] + m1[15]*m2[14],
		m1[12]*m2[3] + m1[13]*m2[7] + m1[14]*m2[11] + m1[15]*m2[15]}
}

func (m Mat4x4) Invert() Mat4x4 {
	var s, c [6]float32
	s[0] = m[0]*m[5] - m[4]*m[1]
	s[1] = m[0]*m[6] - m[4]*m[2]
	s[2] = m[0]*m[7] - m[4]*m[3]
	s[3] = m[1]*m[6] - m[5]*m[2]
	s[4] = m[1]*m[7] - m[5]*m[3]
	s[5] = m[2]*m[7] - m[6]*m[3]

	c[0] = m[8]*m[13] - m[12]*m[9]
	c[1] = m[8]*m[14] - m[12]*m[10]
	c[2] = m[8]*m[15] - m[12]*m[11]
	c[3] = m[9]*m[14] - m[13]*m[10]
	c[4] = m[9]*m[15] - m[13]*m[11]
	c[5] = m[10]*m[15] - m[14]*m[11]

	// assumes it is invertible
	var identity float32 = 1.0 / (s[0]*c[5] - s[1]*c[4] + s[2]*c[3] + s[3]*c[2] - s[4]*c[1] + s[5]*c[0])

	return Mat4x4{
		(m[5]*c[5] - m[6]*c[4] + m[7]*c[3]) * identity,
        (-m[1]*c[5] + m[2]*c[4] - m[3]*c[3]) * identity,
        (m[13]*s[5] - m[14]*s[4] + m[15]*s[3]) * identity,
        (-m[9]*s[5] + m[10]*s[4] - m[11]*s[3]) * identity,
		(-m[4]*c[5] + m[6]*c[2] - m[7]*c[1]) * identity,
        (m[0]*c[5] - m[2]*c[2] + m[3]*c[1]) * identity,
        (-m[12]*s[5] + m[14]*s[2] - m[15]*s[1]) * identity,
        (m[8]*s[5] - m[10]*s[2] + m[11]*s[1]) * identity,
		(m[4]*c[4] - m[5]*c[2] + m[7]*c[0]) * identity,
        (-m[0]*c[4] + m[1]*c[2] - m[3]*c[0]) * identity,
        (m[12]*s[4] - m[13]*s[2] + m[15]*s[0]) * identity,
        (-m[8]*s[4] + m[9]*s[2] - m[11]*s[0]) * identity,
		(-m[4]*c[3] + m[5]*c[1] - m[6]*c[0]) * identity,
		(m[0]*c[3] - m[1]*c[1] + m[2]*c[0]) * identity,
		(-m[12]*s[3] + m[13]*s[1] - m[14]*s[0]) * identity,
		(m[8]*s[3] - m[9]*s[1] + m[10]*s[0]) * identity}
}
