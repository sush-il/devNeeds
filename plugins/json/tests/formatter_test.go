package jsonUtils

import (
	"testing"

	jsonUtils "github.com/sush-il/devNeeds/plugins/json"
)

func TestJSONFormatting(t *testing.T) {
	t.Run("Test invalid json file", func(t *testing.T) {
		result := jsonUtils.FormatJSON("example.json", "space", 2)

	})
}
