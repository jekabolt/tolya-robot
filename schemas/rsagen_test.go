package schemas

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	// create the schema
	bs, _ := json.Marshal(Post{})
	fmt.Printf("%s", bs)
}
