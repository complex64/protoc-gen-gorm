package require

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/testing/protocmp"
)

func MessageOption(t require.TestingT, message, option proto.Message) {
	optKind := option.ProtoReflect().Descriptor()
	optName := optKind.FullName()

	if haveOpt := readOpt(optName, message); haveOpt != nil {
		EqualProtos(t, option, haveOpt)
	} else {
		msg := fmt.Sprintf("option %s is not present on %+v", optName, message)
		require.Fail(t, msg)
	}
}

func readOpt(name protoreflect.FullName, msg proto.Message) (opt proto.Message) {
	msg.ProtoReflect().
		Descriptor().
		Options().
		ProtoReflect().
		Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			if d.Message().FullName() != name {
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
	require.True(t, equal && diff == "" || diff == "", diff, diff)
}
