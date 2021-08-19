package orm

import (
	"log"
	"os"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestCreatingAnItem(t *testing.T) {
	setupBackend()
	defer os.RemoveAll(backendPath)
	e := &Entry{
		Key: proto.String("k1"),
		Val: proto.String("v1"),
	}
	i, err := AddItem(e)
	if err != nil {
		log.Fatal("Couldnt' create item")
	}
	if i != 0 {
		t.Fatal("Expected to create an item with index 0")
	}
}
func TestCreatingTwoItems(t *testing.T) {
	setupBackend()
	// defer os.RemoveAll(backendPath)
	e := &Entry{
		Key: proto.String("k1"),
		Val: proto.String("v1"),
	}
	i, err := AddItem(e)
	if err != nil {
		log.Fatal("Couldnt' create item")
	}
	if i != 0 {
		t.Fatal("Expected to create an item with index 0")
	}

	i, err = AddItem(e)
	if err != nil {
		log.Fatal("Couldnt' create item")
	}
	if i != 1 {
		t.Fatalf("Expected to create an item with index 1, got %d", i)
	}

}

// func TestCreatingAnItem(t *testing.T) {
// 	e := &Entry{
// 		Key: proto.String("k1"),
// 		Val: proto.String("v1"),
// 	}
// 	i, err := AddItem(e)
// 	if err != nil {
// 		log.Fatal("Couldnt' create item")
// 	}
// 	if i != 0 {
// 		t.Fatal("Expected to create an item with index 0")
// 	}
// }
