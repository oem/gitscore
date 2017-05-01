package ansi

import (
	ui "github.com/gizak/termui"
	"github.com/oem/gitscore/github"
)

// Draw deals with "drawing" the whole dashboard. Details about widgets used and layout are internal only.
func Draw(contributors github.Contributors) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	bc := createBarChart(contributors)
	ui.Body.AddRows(ui.NewRow(ui.NewCol(12, 0, bc)))
	// calculate layout
	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop()
	return nil
}

func createBarChart(contributors github.Contributors) *ui.BarChart {
	bc := ui.NewBarChart()
	labels, data := labelsAndData(contributors)
	bc.Data = data
	bc.DataLabels = labels
	bc.BarColor = ui.ColorRed
	bc.TextColor = ui.ColorWhite
	bc.BorderLabelFg = ui.ColorWhite
	bc.NumColor = ui.ColorWhite
	bc.BarWidth = 7
	bc.BarGap = 2
	bc.Height = ui.TermHeight()
	bc.BorderLabel = "ALL TIME contributions"
	return bc
}

func labelsAndData(contributors github.Contributors) ([]string, []int) {
	labels := []string{}
	contribs := []int{}
	for _, contributor := range contributors {
		labels = append(labels, contributor.Name)
		contribs = append(contribs, contributor.Commits)
	}
	return labels, contribs
}
