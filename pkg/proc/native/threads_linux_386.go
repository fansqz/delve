package native

import (
	"fmt"
	"github.com/fansqz/delve/pkg/proc"
)

func (t *nativeThread) restoreRegisters(savedRegs proc.Registers) error {
	return fmt.Errorf("restore regs not supported on i386")
}
