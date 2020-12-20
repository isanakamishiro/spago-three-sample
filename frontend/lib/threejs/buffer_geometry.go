package threejs

import (
	"syscall/js"
)

// BufferGeometry is an efficient representation of mesh, line, or point geometry.
// Includes vertex positions, face indices, normals, colors, UVs, and custom attributes
// within buffers, reducing the cost of passing all this data to the GPU.
//
// To read and edit data in BufferGeometry attributes, see BufferAttribute documentation.
//
// For a less efficient but easier-to-use representation of geometry, see Geometry.
type BufferGeometry interface {
	Geometry
}

// geometryImpl extend: [EventDispatcher]
type bufferGeometryImpl struct {
	Geometry
}

// NewBufferGeometry creates a new BufferGeometry.
func NewBufferGeometry() BufferGeometry {
	return NewDefaultBufferGeometryFromJSValue(
		GetJsObject("BufferGeometry").New(),
	)
}

// NewDefaultBufferGeometryFromJSValue creates a new BufferGeometry.
func NewDefaultBufferGeometryFromJSValue(value js.Value) BufferGeometry {
	return &bufferGeometryImpl{
		NewDefaultGeometryFromJSValue(value),
	}
}
