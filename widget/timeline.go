package widget

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Event struct {
	widget.BaseWidget

	Timestamp time.Time
	Summary   string
	Details   fyne.CanvasObject

	Expanded bool
}

func NewEvent(timestamp time.Time, summary string, details fyne.CanvasObject) *Event {
	e := &Event{
		Timestamp: timestamp,
		Summary:   summary,
		Details:   details,
		Expanded:  false,
	}
	e.ExtendBaseWidget(e)

	return e
}

func (e *Event) CreateRenderer() fyne.WidgetRenderer {
	box := container.NewVBox()

	button := widget.NewButtonWithIcon(
		"", theme.MenuExpandIcon(), func() {},
	)

	box.Add(container.NewHBox(
		button,
		widget.NewRichTextFromMarkdown(
			"**"+e.Timestamp.Format(time.UnixDate)+"**",
		),
		widget.NewRichTextFromMarkdown(e.Summary),
	))

	button.OnTapped = func() {
		e.Expanded = !e.Expanded
		if e.Expanded {
			button.Icon = theme.MenuDropDownIcon()
			box.Add(e.Details)
		} else {
			button.Icon = theme.MenuExpandIcon()
			box.Remove(e.Details)
		}
		button.Refresh()
		box.Refresh()
	}

	return widget.NewSimpleRenderer(box)
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
