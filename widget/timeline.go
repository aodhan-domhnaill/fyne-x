package widget

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Event struct {
	widget.BaseWidget

	Timestamp time.Time
	Summary   string
	Details   fyne.CanvasObject
}

func NewEvent(timestamp time.Time, summary string, details fyne.CanvasObject) *Event {
	e := &Event{
		Timestamp: timestamp,
		Summary:   summary,
		Details:   details,
	}
	e.ExtendBaseWidget(e)

	return e
}

func (e *Event) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		widget.NewLabel(e.Summary),
	)
}

type Timeline struct {
	widget.BaseWidget

	Events []Event
}

func NewTimeline() *Timeline {
	t := &Timeline{}
	t.ExtendBaseWidget(t)
	return t
}

func (t *Timeline) CreateRenderer() fyne.WidgetRenderer {
	box := container.NewVBox()

	for _, event := range t.Events {
		box.Add(&event)
	}

	return widget.NewSimpleRenderer(
		container.NewVScroll(box),
	)
}

func (t *Timeline) Add(e Event) {
	t.Events = append(t.Events, e)
	t.Refresh()
}
