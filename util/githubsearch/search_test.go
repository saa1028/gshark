package githubsearch

import "testing"
import "github.com/neal1991/x-patrol/models"
import "github.com/neal1991/x-patrol"

func TestFiltetString(t *testing.T) {
	codeResult := new(models.CodeResult)
	id := 1
	_, _ = models.Engine.Table("code_result").Where("id=?", id).Get(codeResult)
	if !githubsearch.PassFilters(codeResult) {
		t.Log("pass the PassFilters function")
	} else {
		t.Error("failed to pass the PassFilters function")
	}
}
