package environment

const (
	ownerShift          = 48                           // Number of bits to shift the owner into the upper 16 bits
	addressMask         = (1 << ownerShift) - 1        // Mask to extract the address from the addressAndOwner field
	kernelAddressPrefix = uint64(0xffff) << ownerShift // Precomputed prefix for kernel addresses
)
// kernelSymbolInternal is a memory efficient representation of
// a kernel symbol, used internally for storing all symbols.
type kernelSymbolInternal struct {
	name string
	// We save only the low 48 bits of the address, as all (non-percpu) symbols are at 0xffffXXXXXXXXXXXX
	// Owner is a 16-bit index into a slice of seen owners for the symbol table this symbol belongs to.
	// It can only be translated to the owner name if we have the symbol table.
	// To conserve memory, we encode both of them as a single 64-bit integer where the lower 48-bits
	// are the address and the hight 16-bits are the owner index.
	addressAndOwner uint64
}

func newKernelSymbolInternal(name string, address uint64, owner uint16) *kernelSymbolInternal {
	return &kernelSymbolInternal{
		name:            name,
		addressAndOwner: (uint64(owner) << ownerShift) | (address & addressMask),
	}
}

func (ks kernelSymbolInternal) Name() string {
	return ks.name
}

func (ks kernelSymbolInternal) Address() uint64 {
	// Convert truncated address to the real kernel address
	return kernelAddressPrefix | (ks.addressAndOwner & addressMask)
}

func (ks kernelSymbolInternal) owner() uint16 {
	return uint16(ks.addressAndOwner >> ownerShift)
}

func (ks kernelSymbolInternal) Contains(address uint64) bool {
	symbolAddr := ks.Address()
	return symbolAddr <= address && symbolAddr+maxSymbolSize > address
}

func (ks kernelSymbolInternal) Clone() kernelSymbolInternal {
	return kernelSymbolInternal{
		name:            ks.name,
		addressAndOwner: ks.addressAndOwner,
	}
}

// All kernel symbols are at 0xffffXXXXXXXXXXXX, except percpu symbols which we ignore
func validKernelAddr(addr uint64) bool {
	return addr&kernelAddressPrefix == kernelAddressPrefix
}
