package ansi

import ui "github.com/gizak/termui"

func Draw() error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()
	bc := createBarChart()
	ui.Render(bc)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop()
	return nil
}

func createBarChart() *ui.BarChart {
	bc := ui.NewBarChart()
	bc.Width = 50
	bc.Height = 30
	bc.BorderLabel = "score(total)"
	return bc
}
