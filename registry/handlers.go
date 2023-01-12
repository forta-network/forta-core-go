package registry

import (
	"reflect"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/utils/slicemap"
	log "github.com/sirupsen/logrus"
)

type MessageHandler[K registry.MessageInterface] func(logger *log.Entry, msg K) error

type MessageHandlers[K registry.MessageInterface] []MessageHandler[K]

type Handlers struct {
	AfterBlockHandler func(blk *domain.Block) error

	// implementation upgrades
	UpgradeHandlers MessageHandlers[*registry.UpgradeMessage]

	// registration
	SaveAgentHandlers         MessageHandlers[*registry.AgentSaveMessage]
	AgentActionHandlers       MessageHandlers[*registry.AgentMessage]
	SaveScannerHandlers       MessageHandlers[*registry.ScannerSaveMessage]
	ScannerActionHandlers     MessageHandlers[*registry.ScannerMessage]
	UpdateScannerPoolHandlers MessageHandlers[*registry.UpdateScannerPoolMessage]

	// scanner node versions
	ScannerNodeVersionHandlers MessageHandlers[*registry.ScannerNodeVersionMessage]

	// assignment
	DispatchHandlers MessageHandlers[*registry.DispatchMessage]

	// staking
	AgentStakeHandlers       MessageHandlers[*registry.AgentStakeMessage]
	ScannerStakeHandlers     MessageHandlers[*registry.ScannerStakeMessage]
	TransferSharesHandlers   MessageHandlers[*registry.TransferSharesMessage]
	ScannerPoolStakeHandlers MessageHandlers[*registry.ScannerPoolStakeMessage]

	AgentStakeThresholdHandlers   MessageHandlers[*registry.AgentStakeThresholdMessage]
	ScannerStakeThresholdHandlers MessageHandlers[*registry.ScannerStakeThresholdMessage]

	ScannerPoolAllocationHandlers MessageHandlers[*registry.ScannerPoolAllocationMessage]
}

type HandlerRegistry struct {
	afterBlockHandler func(blk *domain.Block) error
	all               slicemap.SliceMap[string, []reflect.Value]
}

// Handle finds the right handler for the message and calls it.
func (reg *HandlerRegistry) Handle(logger *log.Entry, msg registry.MessageInterface) error {
	h, ok := reg.all.Get(msgKey(msg))
	if !ok {
		logger.WithField("action", msg.ActionName()).Warn("no handlers found")
		return nil
	}
	for _, f := range h {
		ret := f.Call([]reflect.Value{reflect.ValueOf(logger), reflect.ValueOf(msg)})
		retErr := ret[0]
		// return when any of the handlers return an error
		if !retErr.IsNil() {
			return retErr.Interface().(error)
		}
	}
	return nil
}

// NewHandlerRegistry creates a new handler registry.
func NewHandlerRegistry(h Handlers) *HandlerRegistry {
	reg := &HandlerRegistry{}

	// set handlers
	reg.afterBlockHandler = h.AfterBlockHandler
	handlers := reflect.ValueOf(h)
	for i := 0; i < handlers.NumField(); i++ {
		field := handlers.Field(i)
		// must be an array
		if field.Kind() != reflect.Slice {
			continue
		}
		// array elements must be a func
		if field.Type().Elem().Kind() != reflect.Func {
			continue
		}
		// for each func in array, validate and collect
		var (
			methods []reflect.Value
			msgType registry.MessageInterface
		)
		for j := 0; j < field.Len(); j++ {
			f := field.Index(j)
			// for each arg type in func
			for k := 0; k < f.Type().NumIn(); k++ {
				in := f.Type().In(k)
				inVal := reflect.New(in.Elem()).Interface()
				// register handler if one of the args is a message
				v, ok := inVal.(registry.MessageInterface)
				if ok {
					methods = append(methods, f)
					msgType = v
				}
			}
		}
		// set handlers
		if methods != nil {
			reg.all.Set(msgKey(msgType), methods)
		}
	}

	return reg
}

func msgKey(msg registry.MessageInterface) string {
	return reflect.TypeOf(msg).String()
}
