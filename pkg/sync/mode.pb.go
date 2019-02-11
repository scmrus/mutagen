// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync/mode.proto

package sync // import "github.com/havoc-io/mutagen/pkg/sync"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SynchronizationMode specifies the mode for synchronization, encoding both
// directionality and conflict resolution behavior.
type SynchronizationMode int32

const (
	// SynchronizationMode_SynchronizationModeDefault represents an unspecified
	// synchronization mode. It is not valid for use with Reconcile. It should
	// be converted to one of the following values based on the desired default
	// behavior.
	SynchronizationMode_SynchronizationModeDefault SynchronizationMode = 0
	// SynchronizationMode_SynchronizationModeTwoWaySafe represents a
	// bidirectional synchronization mode where automatic conflict resolution is
	// performed only in cases where no data would be lost. Specifically, this
	// means that modified contents are allowed to propagate to the opposite
	// endpoint if the corresponding contents on the opposite endpoint are
	// unmodified or deleted. All other conflicts are left unresolved.
	SynchronizationMode_SynchronizationModeTwoWaySafe SynchronizationMode = 1
	// SynchronizationMode_SynchronizationModeTwoWayResolved is the same as
	// SynchronizationMode_SynchronizationModeTwoWaySafe, but specifies that the
	// alpha endpoint should win automatically in any conflict between alpha and
	// beta, including cases where alpha has deleted contents that beta has
	// modified.
	SynchronizationMode_SynchronizationModeTwoWayResolved SynchronizationMode = 2
	// SynchronizationMode_SynchronizationModeOneWaySafe represents a
	// unidirectional synchronization mode where contents and changes propagate
	// from alpha to beta, but won't overwrite any creations or modifications on
	// beta.
	SynchronizationMode_SynchronizationModeOneWaySafe SynchronizationMode = 3
	// SynchronizationMode_SynchronizationModeOneWayReplica represents a
	// unidirectional synchronization mode where contents on alpha are mirrored
	// (verbatim) to beta, overwriting any conflicting contents on beta and
	// deleting any extraneous contents on beta.
	SynchronizationMode_SynchronizationModeOneWayReplica SynchronizationMode = 4
)

var SynchronizationMode_name = map[int32]string{
	0: "SynchronizationModeDefault",
	1: "SynchronizationModeTwoWaySafe",
	2: "SynchronizationModeTwoWayResolved",
	3: "SynchronizationModeOneWaySafe",
	4: "SynchronizationModeOneWayReplica",
}
var SynchronizationMode_value = map[string]int32{
	"SynchronizationModeDefault":        0,
	"SynchronizationModeTwoWaySafe":     1,
	"SynchronizationModeTwoWayResolved": 2,
	"SynchronizationModeOneWaySafe":     3,
	"SynchronizationModeOneWayReplica":  4,
}

func (x SynchronizationMode) String() string {
	return proto.EnumName(SynchronizationMode_name, int32(x))
}
func (SynchronizationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mode_a8cfd4437324afd0, []int{0}
}

func init() {
	proto.RegisterEnum("sync.SynchronizationMode", SynchronizationMode_name, SynchronizationMode_value)
}

func init() { proto.RegisterFile("sync/mode.proto", fileDescriptor_mode_a8cfd4437324afd0) }

var fileDescriptor_mode_a8cfd4437324afd0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xae, 0xcc, 0x4b,
	0xd6, 0xcf, 0xcd, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x09, 0x68,
	0x9d, 0x60, 0xe4, 0x12, 0x0e, 0xae, 0xcc, 0x4b, 0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2c,
	0xc9, 0xcc, 0xcf, 0xf3, 0xcd, 0x4f, 0x49, 0x15, 0x92, 0xe3, 0x92, 0xc2, 0x22, 0xec, 0x92, 0x9a,
	0x96, 0x58, 0x9a, 0x53, 0x22, 0xc0, 0x20, 0xa4, 0xc8, 0x25, 0x8b, 0x45, 0x3e, 0xa4, 0x3c, 0x3f,
	0x3c, 0xb1, 0x32, 0x38, 0x31, 0x2d, 0x55, 0x80, 0x51, 0x48, 0x95, 0x4b, 0x11, 0xa7, 0x92, 0xa0,
	0xd4, 0xe2, 0xfc, 0x9c, 0xb2, 0xd4, 0x14, 0x01, 0x26, 0x1c, 0x26, 0xf9, 0xe7, 0xa5, 0xc2, 0x4c,
	0x62, 0x16, 0x52, 0xe1, 0x52, 0xc0, 0xa9, 0x24, 0x28, 0xb5, 0x20, 0x27, 0x33, 0x39, 0x51, 0x80,
	0xc5, 0x49, 0x2d, 0x4a, 0x25, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f,
	0x23, 0xb1, 0x2c, 0x3f, 0x59, 0x37, 0x33, 0x5f, 0x3f, 0xb7, 0xb4, 0x24, 0x31, 0x3d, 0x35, 0x4f,
	0xbf, 0x20, 0x3b, 0x5d, 0x1f, 0xe4, 0xe5, 0x24, 0x36, 0xb0, 0xff, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x79, 0x67, 0x2b, 0x12, 0x01, 0x00, 0x00,
}