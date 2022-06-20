package deployers

import (
	"context"
	cm "sail/common"
)

type DeployersItf interface {
	SetInChan(inchan <-chan cm.Message)
	SetOutChan(outchan chan<- cm.Message)

	Install(context.Context, cm.Message) error
	Uninstall(cm.Message) error

	Run(context.Context)
}

type BaseDeployer struct {
	inChan  <-chan cm.Message
	outChan chan<- cm.Message
}

func (b *BaseDeployer) SetInChan(inchan <-chan cm.Message) {
	b.inChan = inchan
}

func (b *BaseDeployer) SetOutChan(outchan chan<- cm.Message) {
	b.outChan = outchan
}
