package registry

import (
	"testing"

	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

type testHandlerImpl1 struct {
	val string
}

func (impl *testHandlerImpl1) HandleMessage(logger *logrus.Entry, msg *registry.AgentSaveMessage) error {
	impl.val = msg.AgentID
	return nil
}

type testHandlerImpl2 struct {
	val string
}

func (impl *testHandlerImpl2) HandleMessage(logger *logrus.Entry, msg *registry.DispatchMessage) error {
	impl.val = msg.AgentID
	return nil
}

func TestHandlerRegistry(t *testing.T) {
	r := require.New(t)

	testID1 := "0001"
	testID2 := "0002"

	agentSave1 := &testHandlerImpl1{}
	agentSave2 := &testHandlerImpl1{}
	dispatch1 := &testHandlerImpl2{}

	handlerReg := NewHandlerRegistry(Handlers{
		SaveAgentHandlers: MessageHandlers[*registry.AgentSaveMessage]{agentSave1.HandleMessage, agentSave2.HandleMessage},
		DispatchHandlers:  MessageHandlers[*registry.DispatchMessage]{dispatch1.HandleMessage},
	})

	err := handlerReg.Handle(nil, &registry.AgentSaveMessage{
		AgentMessage: registry.AgentMessage{AgentID: testID1},
	})
	r.NoError(err)

	err = handlerReg.Handle(nil, &registry.DispatchMessage{
		AgentID: testID2,
	})
	r.NoError(err)

	r.Equal(testID1, agentSave1.val)
	r.Equal(testID1, agentSave2.val)
	r.Equal(testID2, dispatch1.val)
}
