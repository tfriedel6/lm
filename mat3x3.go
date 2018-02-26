package lm

import (
	"fmt"

	"github.com/barnex/fmath"
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
	s, c := fmath.Sincos(radians)
	return Mat3x3{
		c, s, 0,
		-s, c, 0,
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

func (m Mat3x3) Invert() Mat3x3 {
	var identity float32 = 1.0 / (m[0]*m[4]*m[8] + m[3]*m[7]*m[2] + m[6]*m[1]*m[5] - m[6]*m[4]*m[2] - m[3]*m[1]*m[8] - m[0]*m[7]*m[5])

	return Mat3x3{
		(m[4]*m[8] - m[5]*m[7]) * identity,
		(m[2]*m[7] - m[1]*m[8]) * identity,
		(m[1]*m[5] - m[2]*m[4]) * identity,
		(m[5]*m[6] - m[3]*m[8]) * identity,
		(m[0]*m[8] - m[2]*m[6]) * identity,
		(m[2]*m[3] - m[0]*m[5]) * identity,
		(m[3]*m[7] - m[4]*m[6]) * identity,
		(m[1]*m[6] - m[0]*m[7]) * identity,
		(m[0]*m[4] - m[1]*m[3]) * identity}
}
