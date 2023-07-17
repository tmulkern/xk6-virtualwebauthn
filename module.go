package virtualwebauthn

import (
	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
)

// Register the extensions on module initialization.
func init() {
	modules.Register("k6/x/virtualwebauthn", New())
}

type (
	RootModule struct{}

	ModuleInstance struct {
		vu      modules.VU
		exports map[string]interface{}
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{} //nolint:exhaustruct
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface and returns
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance { //nolint:varnamelen,ireturn
	instance := &ModuleInstance{
		vu:      vu,
		exports: make(map[string]interface{}),
	}

	instance.exports["VirtualWebAuthn"] = instance.newVirtualWebAuthn

	return instance
}

// Exports implements the modules.Instance interface and returns the exports
// of the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Named:   mi.exports,
		Default: newVirtualWebAuthn(mi.vu),
	}
}

func (mi *ModuleInstance) newVirtualWebAuthn(call goja.ConstructorCall) *goja.Object {
	rt := mi.vu.Runtime()

	return rt.ToValue(newVirtualWebAuthn(mi.vu)).ToObject(rt)
}
