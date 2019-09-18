package hashcode

import (
	"github.com/hashicorp/terraform/helper/hashcode"
	"testing"
)

func TestHashCode(t *testing.T) {
	src := "joshua"
	res := hashcode.String(src)
	t.Log(res)
}
