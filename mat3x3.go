package lm

import (
	"fmt"
	"github.com/void6/math32"
)

type Mat3x3 [9]float32

func (m *Mat3x3) Pointer() *[9]float32 { return (*[9]float32)(m) }
func (m *Mat3x3) Slice() []float32     { return m[:] }

func (m *Mat3x3) String() string {
	return fmt.Sprintf("[%f,%f,%f,\n %f,%f,%f,\n %f,%f,%f,]", m[0], m[3], m[6], m[1], m[4], m[7], m[2], m[5], m[8])
}

func Mat3x3Identity() Mat3x3 {
	return Mat3x3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1}
}

func Mat3x3Translate(v Vec2) Mat3x3 {
	return Mat3x3{
		1, 0, 0,
		0, 1, 0,
		v[0], v[1], 1}
}

func Mat3x3Scale(v Vec2) Mat3x3 {
	return Mat3x3{
		v[0], 0, 0,
		0, v[1], 0,
		0, 0, 1}
}

func Mat3x3Rotate(radians float32) Mat3x3 {
	c := math32.Cos(radians)
	s := math32.Sin(radians)
	return Mat3x3{
		s, -c, 0,
		c, s, 0,
		0, 0, 1}
}

func (m1 Mat3x3) Mul(m2 Mat3x3) Mat3x3 {
	return Mat3x3{
		m1[0]*m2[0] + m1[1]*m2[3] + m1[2]*m2[6],
		m1[0]*m2[1] + m1[1]*m2[4] + m1[2]*m2[7],
		m1[0]*m2[2] + m1[1]*m2[5] + m1[2]*m2[8],
		m1[3]*m2[0] + m1[4]*m2[3] + m1[5]*m2[6],
		m1[3]*m2[1] + m1[4]*m2[4] + m1[5]*m2[7],
		m1[3]*m2[2] + m1[4]*m2[5] + m1[5]*m2[8],
		m1[6]*m2[0] + m1[7]*m2[3] + m1[8]*m2[6],
		m1[6]*m2[1] + m1[7]*m2[4] + m1[8]*m2[7],
		m1[6]*m2[2] + m1[7]*m2[5] + m1[8]*m2[8]}
}

/*
func (m Mat3x3) Invert() Mat3x3 {
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

	return Mat3x3{
		(m[5]*c[5] - m[6]*c[4] + m[7]*c[3]) * identity,
		(-m[4]*c[5] + m[6]*c[2] - m[7]*c[1]) * identity,
		(m[4]*c[4] - m[5]*c[2] + m[7]*c[0]) * identity,
		(-m[4]*c[3] + m[5]*c[1] - m[6]*c[0]) * identity,
		(-m[1]*c[5] + m[2]*c[4] - m[3]*c[3]) * identity,
		(m[0]*c[5] - m[2]*c[2] + m[3]*c[1]) * identity,
		(-m[0]*c[4] + m[1]*c[2] - m[3]*c[0]) * identity,
		(m[0]*c[3] - m[1]*c[1] + m[2]*c[0]) * identity,
		(m[13]*s[5] - m[14]*s[4] + m[15]*s[3]) * identity,
		(-m[12]*s[5] + m[14]*s[2] - m[15]*s[1]) * identity,
		(m[12]*s[4] - m[13]*s[2] + m[15]*s[0]) * identity,
		(-m[12]*s[3] + m[13]*s[1] - m[14]*s[0]) * identity,
		(-m[9]*s[5] + m[10]*s[4] - m[11]*s[3]) * identity,
		(m[8]*s[5] - m[10]*s[2] + m[11]*s[1]) * identity,
		(-m[8]*s[4] + m[9]*s[2] - m[11]*s[0]) * identity,
		(m[8]*s[3] - m[9]*s[1] + m[10]*s[0]) * identity}
}
*/
