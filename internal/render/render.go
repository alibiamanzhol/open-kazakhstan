package render

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"text/tabwriter"

	"github.com/alibiamanzhol/open-kazakhstan/internal/model"
)

func Table(w io.Writer, data []model.Dataset) {
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tTITLE\tCATEGORY\tFORMAT\tRECORDS\tUPDATED")
	for _, d := range data {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%d\t%s\n", d.ID, d.Title, d.Category, d.Format, d.Records, d.Updated)
	}
	tw.Flush()
}
func JSON(w io.Writer, data []model.Dataset) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}
func CSV(w io.Writer, data []model.Dataset) error {
	cw := csv.NewWriter(w)
	defer cw.Flush()
	cw.Write([]string{"id", "title", "organization", "category", "format", "updated", "records"})
	for _, d := range data {
		cw.Write([]string{d.ID, d.Title, d.Organization, d.Category, d.Format, d.Updated, strconv.Itoa(d.Records)})
	}
	return cw.Error()
}
