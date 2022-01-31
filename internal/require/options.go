package require

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/testing/protocmp"
)

// FileOptions asserts that message has options set for its protobuf message options.
func FileOptions(t require.TestingT, options, message proto.Message) {
	if haveOpts := fileOpts(options, message); haveOpts != nil {
		EqualProtos(t, options, haveOpts)
	} else {
		msg := fmt.Sprintf("missing file options: %s", protojson.Format(options))
		require.Fail(t, msg)
	}
}

// MessageOption asserts the message's parent file has options set on the file level.
func MessageOption(t require.TestingT, options, message proto.Message) {
	if haveOpts := msgOpts(options, message); haveOpts != nil {
		EqualProtos(t, options, haveOpts)
	} else {
		msg := fmt.Sprintf("missing message options: %s", protojson.Format(options))
		require.Fail(t, msg)
	}
}

func FieldOption(t require.TestingT, options, message proto.Message, field string) {
	if haveOpts := fieldOpts(options, message, field); haveOpts != nil {
		EqualProtos(t, options, haveOpts)
	} else {

		msg := fmt.Sprintf("field `%s` missing options: %s", field, protojson.Format(options))
		require.Fail(t, msg)
	}
}

func name(msg proto.Message) protoreflect.FullName {
	return msg.ProtoReflect().Descriptor().FullName()
}

func fileOpts(option, msg proto.Message) proto.Message {
	fd := msg.ProtoReflect().Descriptor().Parent().(protoreflect.FileDescriptor)
	optsName := name(option)
	optsMsg := fd.Options().ProtoReflect()
	return opts(optsName, optsMsg)
}

func msgOpts(options, msg proto.Message) proto.Message {
	optsName := name(options)
	optsMsg := msg.ProtoReflect().Descriptor().Options().ProtoReflect()
	return opts(optsName, optsMsg)
}

func fieldOpts(options, msg proto.Message, field string) proto.Message {
	fields := msg.ProtoReflect().Descriptor().Fields()
	if f := fields.ByName(protoreflect.Name(field)); f != nil {
		optName := name(options)
		fo := f.Options().ProtoReflect()
		return opts(optName, fo)
	}
	return nil
}

func opts(name protoreflect.FullName, msg protoreflect.Message) (opts proto.Message) {
	msg.Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if d.Message() == nil || d.Message().FullName() != name {
			return true // continue
		}
		opts = v.Message().Interface()
		return false // break
	})
	return opts
}

func EqualProtos(t require.TestingT, expected, actual proto.Message) {
	equal := cmp.Equal(expected.ProtoReflect().New(), actual, protocmp.Transform())
	diff := cmp.Diff(expected, actual, protocmp.Transform())
	require.True(t, equal && diff == "" || diff == "", diff)
}
