package serialiser

import (
	"os"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestSerilisationWorksProperly(t *testing.T) {
	fn := "test_filename"
	e := &Entry{
		Key: proto.String("k1"),
		Val: proto.String("v1"),
	}
	Serialise(fn, e)
	os.Remove(fn)
}

func TestUnserialisationWorksPropertly(t *testing.T) {
	fn := "test_filename"
	e := &Entry{
		Key: proto.String("k1"),
		Val: proto.String("v1"),
	}
	Serialise(fn, e)
	nm := &Entry{}
	Deserialise(fn, nm)
	os.Remove(fn)
}
