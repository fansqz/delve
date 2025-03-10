package ebpf

import (
	"reflect"

	"github.com/fansqz/delve/pkg/dwarf/godwarf"
	"github.com/fansqz/delve/pkg/dwarf/op"
)

type UProbeArgMap struct {
	Offset int64        // Offset from the stackpointer.
	Size   int64        // Size in bytes.
	Kind   reflect.Kind // Kind of variable.
	Pieces []int        // Pieces of the variables as stored in registers.
	InReg  bool         // True if this param is contained in a register.
	Ret    bool         // True if this param is a return value.
}

type RawUProbeParam struct {
	Pieces   []op.Piece
	RealType godwarf.Type
	Kind     reflect.Kind
	Len      int64
	Base     uint64
	Addr     uint64
	Data     []byte
}

type RawUProbeParams struct {
	FnAddr       int
	GoroutineID  int
	IsRet        bool
	InputParams  []*RawUProbeParam
	ReturnParams []*RawUProbeParam
}
