package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/x/fyne/layout"
)

type DetailsTable struct {
	widget.BaseWidget
	scroller       fyne.Widget
	tableContainer *fyne.Container
}

type detailsRowElement struct {
	widget.BaseWidget
	render fyne.CanvasObject
}

var _ fyne.Widget = (*detailsRowElement)(nil)

func newDetailsRowElement(o fyne.CanvasObject) layout.ResponsiveWidget {
	ro := &detailsRowElement{render: o}
	ro.ExtendBaseWidget(ro)
	return ro
}

func (ro *detailsRowElement) HandleResize(newPos fyne.Position, windowSize, containerSize fyne.Size) {
	if newPos.X+ro.MinSize().Width > containerSize.Width-theme.Padding() {
		ro.Hide()
	} else {
		ro.Show()
		ro.Resize(ro.MinSize())
		ro.Move(newPos)
	}
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
	box := container.NewVBox()
	scroller := container.NewVScroll(box)
	rt := &DetailsTable{
		tableContainer: box,
		scroller:       scroller,
	}
	rt.ExtendBaseWidget(rt)

	return rt
}

func (rt *DetailsTable) AddRow(ele []fyne.CanvasObject) {
	l := layout.NewResponsiveLayout()
	for _, e := range ele {
		l.Add(newDetailsRowElement(e))
	}
	rt.tableContainer.Add(l)
}

func (rt *DetailsTable) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(rt.scroller)
}
