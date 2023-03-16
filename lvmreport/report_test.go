package lvmreport

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGroupNameFields(t *testing.T) {
	got := LV.fields()
	want := []string{"lv_uuid", "lv_name", "lv_all"}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("fields difference (-got +want):\n%s", diff)
	}
}
