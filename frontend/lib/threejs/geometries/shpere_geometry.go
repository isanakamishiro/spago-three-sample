package geometries

import "app/frontend/lib/threejs"

// SphereBufferGeometry is the BufferGeometry port of SphereGeometry.
type SphereBufferGeometry interface {
	threejs.BufferGeometry
}

type sphereBufferGeometryImp struct {
	threejs.BufferGeometry
}

// NewSphereBufferGeometry is constructor.
// radius — sphere radius. Default is 1.
// widthSegments — number of horizontal segments. Minimum value is 3, and the default is 8.
// heightSegments — number of vertical segments. Minimum value is 2, and the default is 6.
// phiStart — specify horizontal starting angle. Default is 0.
// phiLength — specify horizontal sweep angle size. Default is Math.PI * 2.
// thetaStart — specify vertical starting angle. Default is 0.
// thetaLength — specify vertical sweep angle size. Default is Math.PI.
//
// The geometry is created by sweeping and calculating vertexes around the Y axis (horizontal sweep) and the Z axis (vertical sweep).
// Thus, incomplete spheres (akin to 'sphere slices') can be created
// through the use of different values of phiStart, phiLength, thetaStart and thetaLength,
// in order to define the points in which we start (or end) calculating those vertices.
func NewSphereBufferGeometry(
	radius float64, widthSegments int, heightSegments int,
) SphereBufferGeometry {

	return &planeBufferGeometryImp{
		threejs.NewDefaultBufferGeometryFromJSValue(
			threejs.GetJsObject("SphereBufferGeometry").New(
				radius, widthSegments, heightSegments),
		),
	}
}
