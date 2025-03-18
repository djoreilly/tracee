//go:build (ppc64 || ppc64le)
package environment

// powerpc64 assumes that the kernel will be loaded at an address higher than
// 0x1000000000000000 but it defaults to 0xc000000000000000
//
// We can validate most kernel addresses but can't use any of the address space
// to hold the owner

const (
	minBits = 60
	kernelMask = ^uint64(1 << minBits)
)


type kernelSymbolInternal struct {
	name string
	address uint64
	ownerIndex uint16
}

func newKernelSymbolInternal(name string, address uint64, owner uint16) *kernelSymbolInternal {
	return &kernelSymbolInternal{
		name:           name,
		address:	address,
		ownerIndex:	owner,
	}
}

func (ks kernelSymbolInternal) Name() string {
	return ks.name
}

func (ks kernelSymbolInternal) Address() uint64 {
	return ks.address
}

func (ks kernelSymbolInternal) owner() uint16 {
	return ks.ownerIndex
}

func (ks kernelSymbolInternal) Contains(address uint64) bool {
	symbolAddr := ks.Address()
	return symbolAddr <= address && symbolAddr+maxSymbolSize > address
}

func (ks kernelSymbolInternal) Clone() kernelSymbolInternal {
	return kernelSymbolInternal{
		name:            ks.name,
		address:         ks.address,
		ownerIndex:	 ks.ownerIndex,
	}
}

// Percpus aren't technically valid but require lookups to validate
// If the kernel is built with a start address above 1 << 60, this could
// flag userspace addresses as valid but it should work generally
func validKernelAddr(addr uint64) bool {
	return (addr & kernelMask) != 0
}
