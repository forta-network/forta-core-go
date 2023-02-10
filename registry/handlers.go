package registry

import (
	"context"
	"reflect"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/forta-network/forta-core-go/utils/slicemap"
	log "github.com/sirupsen/logrus"
)

// Handlers helps declaring all handlers in the order they are intended to be executed.
type Handlers struct {
	AfterBlockHandler func(blk *domain.Block) error

	// implementation upgrades
	UpgradeHandlers regmsg.HandlerFuncs[*registry.UpgradeMessage]

	// registration
	SaveAgentHandlers         regmsg.HandlerFuncs[*registry.AgentSaveMessage]
	AgentActionHandlers       regmsg.HandlerFuncs[*registry.AgentMessage]
	SaveScannerHandlers       regmsg.HandlerFuncs[*registry.ScannerSaveMessage]
	ScannerActionHandlers     regmsg.HandlerFuncs[*registry.ScannerMessage]
	UpdateScannerPoolHandlers regmsg.HandlerFuncs[*registry.UpdateScannerPoolMessage]

	// scanner node versions
	ScannerNodeVersionHandlers regmsg.HandlerFuncs[*registry.ScannerNodeVersionMessage]

	// assignment
	DispatchHandlers regmsg.HandlerFuncs[*registry.DispatchMessage]

	// staking
	AgentStakeHandlers       regmsg.HandlerFuncs[*registry.AgentStakeMessage]
	ScannerStakeHandlers     regmsg.HandlerFuncs[*registry.ScannerStakeMessage]
	TransferSharesHandlers   regmsg.HandlerFuncs[*registry.TransferSharesMessage]
	ScannerPoolStakeHandlers regmsg.HandlerFuncs[*registry.ScannerPoolStakeMessage]

	AgentStakeThresholdHandlers   regmsg.HandlerFuncs[*registry.AgentStakeThresholdMessage]
	ScannerStakeThresholdHandlers regmsg.HandlerFuncs[*registry.ScannerStakeThresholdMessage]

	ScannerPoolAllocationHandlers regmsg.HandlerFuncs[*registry.ScannerPoolAllocationMessage]
}

type HandlerRegistry struct {
	afterBlockHandler func(blk *domain.Block) error
	all               slicemap.SliceMap[string, []reflect.Value]
}

// Handle finds the right handler for the message and calls it.
func (reg *HandlerRegistry) Handle(ctx context.Context, logger *log.Entry, msg regmsg.Interface) error {
	logger = logger.WithField("action", msg.ActionName())

	h, ok := reg.all.Get(msgKey(msg))
	if !ok {
		logger.Warn("no handlers found")
		return nil
	}

	logger.Info("handling message")
	for _, f := range h {
		ret := f.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(logger), reflect.ValueOf(msg)})
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
			msgType regmsg.Interface
		)
		for j := 0; j < field.Len(); j++ {
			f := field.Index(j)
			// for each arg type in func
			// starts from 1 because first arg is context.Context
			for k := 1; k < f.Type().NumIn(); k++ {
				in := f.Type().In(k)
				inVal := reflect.New(in.Elem()).Interface()
				// register handler if one of the args is a message
				v, ok := inVal.(regmsg.Interface)
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

func msgKey(msg regmsg.Interface) string {
	return reflect.TypeOf(msg).String()
}
