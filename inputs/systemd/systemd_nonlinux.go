//go:build !linux
// +build !linux

package systemd

import (
	"github.com/F1997/categraf/types"
)

func (s *Systemd) Init() error {
	return nil
}

func (s *Systemd) Gather(slist *types.SampleList) {
}
