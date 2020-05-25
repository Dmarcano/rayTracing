package ray

import (
	"github.com/Dmarcan/rayTracing/vector"
	utils "github.com/Dmarcan/rayTracing/vector/utils"
)

type Ray struct {
	Origin vector.Vector3
	Dir    vector.Vector3
}

func (r Ray) PointAt(t float64) *vector.Vector3 {
	vector_out := utils.Mult(r.Dir, t)
	return utils.Add(r.Origin, *vector_out)
}
