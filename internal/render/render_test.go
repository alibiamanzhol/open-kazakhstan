package render

import (
	"bytes"
	"github.com/alibiamanzhol/open-kazakhstan/internal/model"
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	var b bytes.Buffer
	Table(&b, []model.Dataset{{ID: "x", Title: "Roads", Category: "Transport", Format: "CSV", Records: 3, Updated: "2026-01-01"}})
	if !strings.Contains(b.String(), "Roads") {
		t.Fatal("missing title")
	}
}
