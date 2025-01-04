package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetEmployeeName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (vp *MyVicePresident) GetName() string {
	return vp.Name
}

func (vp *MyVicePresident) GetVicePresidentName() string {
	return vp.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "razi", GetEmployeeName[Manager](&MyManager{Name: "razi"}))
	assert.Equal(t, "aziz", GetEmployeeName[VicePresident](&MyVicePresident{Name: "aziz"}))
}
