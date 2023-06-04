package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/x/fyne/layout"
)

type DetailsTable struct {
	widget.BaseWidget
	tableContainer *fyne.Container
}

type detailsRowElement struct {
	render fyne.CanvasObject
}

func (ro *detailsRowElement) CreateRenderer() fyne.WidgetRenderer {
	if ro.render == nil {
		return nil
	}
	if o, ok := ro.render.(fyne.Widget); ok {
		return o.CreateRenderer()
	}
	return widget.NewSimpleRenderer(ro.render)
}

func NewDetailsTable() *DetailsTable {
	rt := &DetailsTable{
		tableContainer: container.NewVBox(),
	}

	rt.ExtendBaseWidget(rt)

	return rt
}

func (rt *DetailsTable) AddRow(ele []fyne.CanvasObject, colConstants []float32) {
	l := layout.NewResponsiveLayout()
	scale := 1 / float32(len(ele))
	for i, e := range ele {
		row := layout.Responsive(
			e,
			colConstants[i]*scale,
			colConstants[i]*scale,
			colConstants[i]*scale,
			colConstants[i]*scale,
		)
		if i > 1 {
			row.Hidable(true)
		}
		l.Add(row)
	}
	rt.tableContainer.Add(l)
}

func (rt *DetailsTable) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewVScroll(rt.tableContainer))
}
