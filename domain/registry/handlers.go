package registry

import (
	log "github.com/sirupsen/logrus"
)

type SaveAgentMessageHandler interface {
	HandleAgentSave(logger *log.Entry, agt *AgentSaveMessage) error
}

type AgentMessageHandler interface {
	HandleAgentMessage(logger *log.Entry, agt *AgentMessage) error
}

type SaveScannerMessageHandler interface {
	HandleScannerSave(logger *log.Entry, agt *ScannerSaveMessage) error
}

type ScannerMessageHandler interface {
	HandleScannerMessage(logger *log.Entry, agt *ScannerMessage) error
}

type DispatchMessageHandler interface {
	HandleDispatchMessage(logger *log.Entry, agt *DispatchMessage) error
}
