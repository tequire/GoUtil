package gin

import (
	"testing"

	"github.com/gin-gonic/gin"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	Convey("gets user from context", t, func() {
		context, _ := gin.CreateTestContext(nil)
		user := TestUser()
		context.Set(UserInContext, &user)
		_, err := GetUser(context)
		assert.Nil(t, err, "no error")
	})
}
