package environment


// s390 uses separate address spaces for kernel and userspace.
// Address range checks don't apply and will result in missed symbols
// Current kernel support up to 53 bits of physical memory and it is
// assumed that the entire address space is available for use.
//
// hex "Kernel image base address"
// range 0x100000 0x1FFFFFE0000000 if !KASAN
// range 0x100000 0x1BFFFFE0000000 if KASAN
// default 0x3FFE0000000 if !KASAN
// default 0x7FFFE0000000 if KASAN

type kernelSymbolInternal struct {
	name string
	address uint64
	ownerIndex uint16
}

func newKernelSymbolInternal(name string, address uint64, owner uint16) *kernelSymbolInternal {
	return &kernelSymbolInternal{
		name:            name,
		address: address,
		ownerIndex: owner,
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
func validKernelAddr(addr uint64) bool {
	return true
}
