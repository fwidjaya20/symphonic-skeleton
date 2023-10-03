package event

import ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"

type Kernel struct {
}

func (k Kernel) Subscribers() ContractEvent.Collection {
	return ContractEvent.Collection{}
}
