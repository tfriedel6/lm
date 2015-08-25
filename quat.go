package lm

import (
	"fmt"
	"github.com/void6/math32"
)

type Quat [4]float32

func (q Quat) String() string {
	return fmt.Sprintf("[%f,%f,%f,%f]", q[0], q[1], q[2], q[3])
}

func (q Quat) XYZVec() Vec3 {
	return Vec3{q[0], q[1], q[2]}
}

func QuatAxisRotation(axis Vec3, angle float32) Quat {
	var d = 1 / axis.Len()
	halfAngle := angle / 2.0
	s := math32.Sin(halfAngle)
	c := math32.Cos(halfAngle)
	return Quat{s * axis[0] * d, s * axis[1] * d, s * axis[2] * d, c}
}

func QuatIdentity() Quat {
	return Quat{0, 0, 0, 1}
}

func (q1 Quat) MulQuat(q2 Quat) Quat {
	return Quat{
		q1[3]*q2[0] + q1[0]*q2[3] + q1[1]*q2[2] - q1[2]*q2[1],
		q1[3]*q2[1] + q1[1]*q2[3] + q1[2]*q2[0] - q1[0]*q2[2],
		q1[3]*q2[2] + q1[2]*q2[3] + q1[0]*q2[1] - q1[1]*q2[0],
		q1[3]*q2[3] - q1[0]*q2[0] - q1[1]*q2[1] - q1[2]*q2[2]}
}

func (q Quat) Conjugate() Quat {
	return Quat{-q[0], -q[1], -q[2], q[3]}
}

/*
func (q Quaternion) YawPitchRoll() (yaw, pitch, roll float32) {
	yaw = float32(math.Atan2(float64(2*(q.W*q.X+q.Y*q.Z)), float64(1-2*(q.X*q.X+q.Y*q.Y))))
	pitch = float32(math.Asin(float64(2 * (q.W*q.Y - q.Z*q.X))))
	roll = float32(math.Atan2(float64(2*(q.W*q.Z+q.X*q.Y)), float64(1-2*(q.Y*q.Y+q.Z*q.Z))))
	return
}
*/

func (q Quat) Mat4x4() Mat4x4 {
	x, y, z, w := q[0], q[1], q[2], q[3]
	return Mat4x4{
		1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0,
		2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z + 2*w*x, 0,
		2*x*z + 2*w*y, 2*y*z - 2*w*x, 1 - 2*x*x - 2*y*y, 0,
		0, 0, 0, 1,
	}
}