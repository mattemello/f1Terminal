package takedata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeCar(t *testing.T) {

	Now = "2025-04-05T07:00:05"
	Previus = "2025-04-05T07:00:01"

	err := GetDataCar()
	assert.NoError(t, err, "It have an error")
}
