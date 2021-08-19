package orm

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/lymlan/go-orm/serialiser"
	"google.golang.org/protobuf/proto"
)

type Backend int

const (
	backendPath = "./local_backend"
)

var (
	index_path = fmt.Sprintf("%s/index", backendPath)
	lock       sync.Mutex
)

func setupBackend() {

	p := "./local_backend"
	ip := fmt.Sprintf("%s/index", p)
	if _, err := os.Stat(ip); os.IsNotExist(err) {
		os.Mkdir(p, os.FileMode(0777))
		f, err := os.Create(ip)
		if err != nil {
			panic("Could not create index file")
		}
		defer f.Close()
		f.WriteString(fmt.Sprintf("%d\n", 0))
	} else {
		fmt.Println("backend already setup")
	}
}

func AddItem(message proto.Message) (id int, err error) {

	lock.Lock()
	defer lock.Unlock()

	n, err := ioutil.ReadFile(index_path)
	i, err := strconv.Atoi(string(n))
	fn := fmt.Sprintf("%s/%s", backendPath, n)
	if _, err := os.Stat(fn); err == nil {
		panic("Cannot add item to an existing index")
	}
	serialiser.Serialise(fn, message)

	f, err := os.OpenFile(index_path, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic("cannot open the index_path")
	}
	defer f.Close()
	f.Truncate(0)
	_, err = f.WriteString(strconv.Itoa(i + 1))

	return i, err
}

func GetItem(message proto.Message, id int) (err error) {
	return
}

func UpdateItem(message proto.Message, id int) (err error) {
	return
}
