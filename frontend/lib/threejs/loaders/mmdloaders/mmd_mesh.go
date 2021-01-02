package mmdloaders

import (
	"app/frontend/lib/threejs"
	"syscall/js"
)

// MMDMesh ...
type MMDMesh interface {
	threejs.Mesh
}

type mmdMeshImp struct {
	threejs.Mesh
}

func newMMDMeshFromJSValue(v js.Value) MMDMesh {
	return &mmdMeshImp{
		Mesh: threejs.NewMeshFromJSValue(v),
	}
}
