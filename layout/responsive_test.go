package layout

import (
	"math"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

// Check if a simple widget is responsive to fill 100% of the layout.
func TestResponsive_SimpleLayout(t *testing.T) {
	padding := theme.Padding()
	w, h := float32(SMALL), float32(300)

	// build
	label := widget.NewLabel("Hello World")
	layout := NewResponsiveLayout(label)

	win := test.NewWindow(layout)
	defer win.Close()
	win.Resize(fyne.NewSize(w, h))

	size := layout.Size()

	assert.Equal(t, w-padding*2, size.Width)

}

// Test is a basic responsive layout is correctly configured. This test 2 widgets
// with 100% for small size and 50% for medium size or taller.
func TestResponsive_Responsive(t *testing.T) {
	padding := theme.Padding()

	// build
	label1 := Responsive(widget.NewLabel("Hello World"), 1, .5)
	label2 := Responsive(widget.NewLabel("Hello World"), 1, .5)

	win := test.NewWindow(
		NewResponsiveLayout(label1, label2),
	)
	win.SetPadded(true)
	defer win.Close()

	// First, we are at w < SMALL so the labels should be sized to 100% of the layout
	w, h := float32(SMALL), float32(300)
	win.Resize(fyne.NewSize(w, h))
	size1 := label1.Size()
	size2 := label2.Size()
	assert.Equal(t, w-padding*2, size1.Width)
	assert.Equal(t, w-padding*2, size2.Width)

	// Then resize to w > SMALL so the labels should be sized to 50% of the layout
	w = float32(MEDIUM)
	win.Resize(fyne.NewSize(w, h))
	size1 = label1.Size()
	size2 = label2.Size()
	// remove 2 * padding as there is 2 objects in a line
	assert.Equal(t, w/2-padding*2, size1.Width)
	assert.Equal(t, w/2-padding*2, size2.Width)
}

// Test is a basic responsive layout is correctly configured. This test 2 widgets
// with 100% for small size and 50% for medium size or taller.
func TestResponsive_HidableResponsive(t *testing.T) {
	padding := theme.Padding()

	// build
	label1 := HidableResponsive(widget.NewLabel("Hello World"), 1, .5)
	label2 := HidableResponsive(widget.NewLabel("Hello World"), 1, .5)
	label2.Hidable(true, false)

	win := test.NewWindow(
		NewResponsiveLayout(label1, label2),
	)
	win.SetPadded(true)
	defer win.Close()

	// First, we are at w < SMALL so the labels should be sized to 100% of the layout
	w, h := float32(SMALL), float32(300)
	win.Resize(fyne.NewSize(w, h))
	size1 := label1.Size()
	size2 := label2.Size()
	assert.Equal(t, w-padding*2, size1.Width)
	assert.Equal(t, float32(0), size2.Width)

	// Then resize to w > SMALL so the labels should be sized to 50% of the layout
	w = float32(MEDIUM)
	win.Resize(fyne.NewSize(w, h))
	size1 = label1.Size()
	size2 = label2.Size()
	// remove 2 * padding as there is 2 objects in a line
	assert.Equal(t, w/2-padding*2, size1.Width)
	assert.Equal(t, w/2-padding*2, size2.Width)
}

// Check if a widget that overflows the container goes to the next line.
func TestResponsive_GoToNextLine(t *testing.T) {
	w, h := float32(200), float32(300)

	// build
	label1 := Responsive(widget.NewLabel("Hello World"), .5)
	label2 := Responsive(widget.NewLabel("Hello World"), .5)
	label3 := Responsive(widget.NewLabel("Hello World"), .5)

	layout := NewResponsiveLayout(label1, label2, label3)
	win := test.NewWindow(layout)
	defer win.Close()
	w = float32(MEDIUM)
	win.Resize(fyne.NewSize(w, h))

	// label 1 and 2 are on the same line
	assert.Equal(t, label1.Position().Y, label2.Position().Y)

	// but not the label 3
	assert.NotEqual(t, label1.Position().Y, label3.Position().Y)

	// just to be sure...
	// the label3 should be at label1.Position().Y + label1.Size().Height
	assert.Equal(t, label1.Position().Y+label1.Size().Height, label3.Position().Y)
}

// Check if sizes are correctly computed for responsive widgets when the window size
// changes. There are some corner cases with padding. So, it needs to be improved...
func TestResponsive_SwitchAllSizes(t *testing.T) {
	// build
	n := 4
	labels := make([]fyne.CanvasObject, n)
	for i := 0; i < n; i++ {
		labels[i] = Responsive(widget.NewLabel("Hello World"), 1, 1/float32(2), 1/float32(3), 1/float32(4))
	}
	layout := NewResponsiveLayout(labels...)
	win := test.NewWindow(layout)
	defer win.Close()

	h := float32(1200)
	p := theme.Padding()
	// First, we are at w < SMALL so the labels should be sized to 100% of the layout
	w := float32(SMALL)
	win.Resize(fyne.NewSize(w, h))
	win.Content().Refresh()
	w = w - 2*p
	for i := 0; i < n; i++ {
		size := labels[i].Size()
		assert.Equal(t, w, size.Width)
	}

	// Then resize to w > SMALL so the labels should be sized to 50% of the layout
	w = float32(MEDIUM)
	win.Resize(fyne.NewSize(w, h))
	w = w - 2*p
	for i := 0; i < n; i++ {
		size := labels[i].Size()
		assert.Equal(t, w/2-p, size.Width) // 1 padding between 2 widgets
	}

	// Then resize to w > MEDIUM so the labels should be sized to 33% of the layout
	w = float32(LARGE)
	win.Resize(fyne.NewSize(w, h))
	w = w - 2*p
	for i := 0; i < n-1; i++ { // note: n-1 because the last element seems to be resized
		size := labels[i].Size()
		floor := math.Floor(float64(w)/3 - float64(p)/3)
		assert.Equal(t, float32(floor), size.Width) // 2 paddings between 3 widgets
	}

	// Then resize to w > LARGE so the labels should be sized to 25% of the layout
	w = float32(XLARGE)
	win.Resize(fyne.NewSize(w, h))
	w = w - 2*p
	for i := 0; i < n; i++ {
		size := labels[i].Size()
		// note: 1px is added to the size to avoid rounding errors
		assert.Equal(t, w/4-p/4-1, float32(math.Floor(float64(size.Width))))
	}
}

// Test if a widget is responsive to fill 100% of the layout
// when we don't provides rsponsive ratios.
func TestResponsive_NoArgs(t *testing.T) {
	label := widget.NewLabel("Hello World")
	resp := NewResponsiveLayout(Responsive(label))
	for _, child := range resp.Objects {
		ro, ok := child.(*responsiveWidget)
		assert.Equal(t, true, ok)
		for _, s := range []responsiveBreakpoint{SMALL, MEDIUM, LARGE, XLARGE} {
			assert.Equal(t, float32(1), ro.responsiveConfig[s])
		}
	}
}
