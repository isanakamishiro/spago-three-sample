package threejs

import "syscall/js"

// Mesh is class representing triangular polygon mesh based objects.
// Also serves as a base for other classes such as SkinnedMesh.
type Mesh interface {
	Object3D
}

// MeshImpl extend: [Object3D]
type meshImpl struct {
	Object3D
}

// NewMesh is factory method for MeshImpl.
func NewMesh(geometry Geometry, material Material) Mesh {
	return &meshImpl{
		NewObject3DFromJSValue(GetJsObject("Mesh").New(geometry.JSValue(), material.JSValue())),
	}
}

// NewMeshWithMultiMaterial is factory method for MeshImpl.
func NewMeshWithMultiMaterial(geometry Geometry, materialSlice []Material) Mesh {

	var a []interface{} = make([]interface{}, len(materialSlice))
	for i, v := range materialSlice {
		a[i] = v.JSValue()
	}

	return &meshImpl{
		NewObject3DFromJSValue(GetJsObject("Mesh").New(geometry.JSValue(), a)),
	}
}

// NewMeshFromJSValue is factory method for MeshImpl.
func NewMeshFromJSValue(value js.Value) Mesh {
	return &meshImpl{
		NewObject3DFromJSValue(value),
	}
}
