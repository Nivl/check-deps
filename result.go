package main

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/Nivl/check-deps/modutil"
	"github.com/olekukonko/tablewriter"
)

// Results contains all the modules that need to be reported
type Results struct {
	Updated  []*modutil.Module
	Replaced []*modutil.Module
	Old      []*modutil.Module
}

// HasModules checks if the results contains any modules
func (r *Results) HasModules() bool {
	return len(r.Updated) > 0 ||
		len(r.Replaced) > 0 ||
		len(r.Old) > 0
}

func (r *Results) print(w io.Writer) {
	if len(r.Updated) > 0 {
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{"Module", "Current Version", "New Version", "Indirect"})
		for _, m := range r.Updated {
			table.Append([]string{
				m.Path,
				m.Version,
				m.Update.Version,
				strconv.FormatBool(m.Indirect),
			})
		}
		table.Render()
	}

	if len(r.Replaced) > 0 {
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{"Module", "Replaced By", "Indirect"})
		for _, m := range r.Replaced {
			table.Append([]string{
				m.Path,
				m.Replace.Path,
				strconv.FormatBool(m.Indirect),
			})
		}
		table.Render()
	}

	if len(r.Old) > 0 {
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{"Module", "Last update", "Indirect"})
		for _, m := range r.Old {
			monthsPassed := time.Since(*m.Time) / (24 * time.Hour) / 30
			table.Append([]string{
				m.Path,
				fmt.Sprintf("%d months ago (%s)", monthsPassed, m.Time.Format("2006/01/02")),
				strconv.FormatBool(m.Indirect),
			})
		}
		table.Render()
	}
}