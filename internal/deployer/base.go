package deployer

import (
	"context"
	cm "sail/common"
)

type DeployersItf interface {
	SetInChan(inchan <-chan cm.InMessage)
	SetOutChan(outchan chan<- cm.InMessage)

	Install(context.Context, cm.InMessage) error
	Uninstall(cm.InMessage) error

	Run(context.Context)
}

type BaseDeployer struct {
	inChan  <-chan cm.InMessage
	outChan chan<- cm.InMessage
}

func (b *BaseDeployer) SetInChan(inchan <-chan cm.InMessage) {
	b.inChan = inchan
}

func (b *BaseDeployer) SetOutChan(outchan chan<- cm.InMessage) {
	b.outChan = outchan
}
