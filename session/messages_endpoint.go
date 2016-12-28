package session

import (
	"github.com/havoc-io/mutagen/rsync"
	"github.com/havoc-io/mutagen/sync"
)

type initializeRequest struct {
	Session string
	Version Version
	Root    string
	Ignores []string
	Alpha   bool
}

type initializeResponse struct {
	PreservesExecutability bool
	Error                  string
}

type scanRequest struct {
	BaseSnapshotSignature    []rsync.BlockHash
	ExpectedSnapshotChecksum []byte
}

type scanResponse struct {
	SnapshotDelta []rsync.Operation
	Error         string
}

type transmitRequest struct {
	Path          string
	BaseSignature []rsync.BlockHash
}

type transmitResponse struct {
	Operation rsync.Operation
	Error     string
}

type stageRequest struct {
	Transitions []sync.Change
}

type stageResponse struct {
	Status StagingStatus
	Done   bool
	Error  string
}

type applyRequest struct {
	Transitions []sync.Change
}

type applyResponse struct {
	Changes  []sync.Change
	Problems []sync.Problem
	Error    string
}