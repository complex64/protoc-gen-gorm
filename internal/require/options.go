package require

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/testing/protocmp"
)

func FileOptions(t require.TestingT, option, message proto.Message) {
	if haveOpt := fileOpt(option, message); haveOpt != nil {
		EqualProtos(t, option, haveOpt)
	} else {
		msg := fmt.Sprintf("missing file option: %+v", option)
		require.Fail(t, msg)
	}
}

func MessageOption(t require.TestingT, option, message proto.Message) {
	if haveOpt := msgOpt(option, message); haveOpt != nil {
		EqualProtos(t, option, haveOpt)
	} else {
		msg := fmt.Sprintf("missing message option: %+v", option)
		require.Fail(t, msg)
	}
}

func name(msg proto.Message) protoreflect.FullName {
	return msg.ProtoReflect().Descriptor().FullName()
}

func fileOpt(option, msg proto.Message) proto.Message {
	fd := msg.ProtoReflect().Descriptor().Parent().(protoreflect.FileDescriptor)
	optName := name(option)
	optMsg := fd.Options().ProtoReflect()
	return opt(optName, optMsg)
}

func msgOpt(option, msg proto.Message) proto.Message {
	optName := name(option)
	optMsg := msg.ProtoReflect().Descriptor().Options().ProtoReflect()
	return opt(optName, optMsg)
}

func opt(name protoreflect.FullName, msg protoreflect.Message) (opt proto.Message) {
	msg.Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if d.Message() == nil || d.Message().FullName() != name {
			return true // continue
		}
		opt = v.Message().Interface()
		return false // break
	})
	return opt
}

func EqualProtos(t require.TestingT, expected, actual proto.Message) {
	equal := cmp.Equal(expected.ProtoReflect().New(), actual, protocmp.Transform())
	diff := cmp.Diff(expected, actual, protocmp.Transform())
	require.True(t, equal && diff == "" || diff == "", diff)
}
