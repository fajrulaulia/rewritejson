package rewritejson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	output := Response([]string{"code:number|500", "message:string|Sorry, Internal Server Error, Please contact Administrator"})
	assert.Equal(t, string(output), `{"code":500,"message":"Sorry, Internal Server Error, Please contact Administrator"}`, "Should be not error")
}
