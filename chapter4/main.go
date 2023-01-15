package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	// checkbox()
	// radio()
	// slider()
	// selection()
	// progress_bar()
	// form()
	// group()
	// tab_container()
	// border_layout()
	// grid_layout()
	// scroll_container()
	// toolbar()
	// mainbar()
	// alert()
	// confirm()
	// custom_dialog()
	entry()

}

func checkbox() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello World")
	c := widget.NewCheck("Check!", func(f bool) {
		if f {
			l.SetText("Checked!")
		} else {
			l.SetText("not checked.")
		}
	})
	c.SetChecked(true)
	w.SetContent(
		widget.NewVBox(
			l, c,
		),
	)
	w.ShowAndRun()
}

func radio() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello World")
	r := widget.NewRadioGroup(
		[]string{"one", "two", "three"},
		func(s string) {
			if s == "" {
				l.SetText("not selected.")
			} else {
				l.SetText("selected: " + s)
			}
		},
	)
	r.SetSelected("one") // 初期値の設定
	w.SetContent(widget.NewVBox(l, r))
	w.ShowAndRun()
}

func slider() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne")
	s := widget.NewSlider(0.0, 10000.)
	b := widget.NewButton("Check", func() {
		l.SetText("value: " + strconv.Itoa(int(s.Value)))
	})
	w.SetContent(widget.NewVBox(l, s, b))
	w.ShowAndRun()
}

func selection() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	sl := widget.NewSelect([]string{
		"Eins", "Twei", "Drei"},
		func(s string) {
			l.SetText("selected: " + s)
		})
	w.SetContent(widget.NewVBox(l, sl))
	w.ShowAndRun()
}

func progress_bar() {
	v := 0.
	a := app.New()
	w := a.NewWindow("hello")
	l := widget.NewLabel("Hello Fyne!")
	p := widget.NewProgressBar()
	b := widget.NewButton("up", func() {
		v += 0.1
		if v > 1.0 {
			v = 0.
		}
		p.SetValue(v)
	})
	w.SetContent(widget.NewVBox(l, p, b))
	w.ShowAndRun()
}

func form() {
	a := app.New()
	w := a.NewWindow("hello")
	l := widget.NewLabel("Hello Fyne!")
	ne := widget.NewEntry()
	pe := widget.NewPasswordEntry()
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewForm(
				widget.NewFormItem("Name", ne),
				widget.NewFormItem("Pass", pe),
			),
			widget.NewButton("OK", func() {
				l.SetText(ne.Text + " & " + pe.Text)
			}),
		),
	)
	w.ShowAndRun()
}

func group() {
	a := app.New()
	w := a.NewWindow("hello")
	l := widget.NewLabel("Hello Fyne!")
	ck1 := widget.NewCheck("check 1", nil)
	ck2 := widget.NewCheck("check 2", nil)
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewGroup("Group", ck1, ck2),
			widget.NewButton("OK", func() {
				re := "Result: "
				if ck1.Checked {
					re += "Check-1"
				}
				if ck2.Checked {
					re += "Check-2"
				}
				l.SetText(re)
			}),
		),
	)
	w.ShowAndRun()
}

func tab_container() {
	a := app.New()
	w := a.NewWindow("hello")
	l := widget.NewLabel("Hello Fyne!")

	w.SetContent(

		widget.NewVBox(
			l,
			widget.NewTabContainer(
				widget.NewTabItem("First",
					widget.NewLabel("This is First tab item.")),
				widget.NewTabItem("Second",
					widget.NewLabel("This is Second tab item.")),

				widget.NewTabItem("Third",
					widget.NewLabel("This is Third tab item.")),
			),
		),
	)
	w.ShowAndRun()
}

func border_layout() {
	a := app.New()
	w := a.NewWindow("hello")
	bt := widget.NewButton("Top", nil)
	bb := widget.NewButton("Button", nil)
	bl := widget.NewButton("Left", nil)
	br := widget.NewButton("Right", nil)
	w.SetContent(
		fyne.NewContainerWithLayout(layout.NewBorderLayout(bt, bb, bl, br),
			bt, bb, bl, br,
			widget.NewLabel("Center."),
		),
	)
	w.ShowAndRun()
}

func grid_layout() {
	a := app.New()
	w := a.NewWindow("hello")
	w.SetContent(
		fyne.NewContainerWithLayout(
			// layout.NewGridLayout(3),
			// layout.NewGridLayoutWithRows(3),
			layout.NewGridWrapLayout(fyne.NewSize(100, 100)),
			widget.NewButton("One", nil),
			widget.NewButton("Two", nil),
			widget.NewButton("Three", nil),
			widget.NewButton("Four", nil),
			layout.NewSpacer(),
			widget.NewButton("Five", nil),
			widget.NewButton("Six", nil),
			layout.NewSpacer(),
			widget.NewButton("Seven", nil),
			widget.NewButton("Eight", nil),
			widget.NewButton("Nine", nil),
			widget.NewButton("Ten", nil),
		),
	)
	w.ShowAndRun()
}

func scroll_container() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewScrollContainer(
			widget.NewVBox(
				widget.NewButton("One", nil),
				widget.NewButton("Two", nil),
				widget.NewButton("Three", nil),
				widget.NewButton("Four", nil),
				widget.NewButton("Five", nil),
				widget.NewButton("Six", nil),
				widget.NewButton("Seven", nil),
				widget.NewButton("Eight", nil),
				widget.NewButton("Nine", nil),
				widget.NewButton("Ten", nil),
			),
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}

func toolbar() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("This is Sample widget.")
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			l.SetText("Select Home Icon!")
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			l.SetText("Select Information Icon!")
		}),
	)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(nil, tb, nil, nil), l, tb,
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

func mainbar() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("This is Sample widget.")

	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				l.SetText("select 'New' menu item.")
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)
	w.SetMainMenu(mm)
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewButton("ok", nil),
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	// ポップアップウィンドウではなくガチのツールバーに項目が増えている
	w.ShowAndRun()
}

func alert() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hell Fyne!")
	b := widget.NewButton("Click", func() {
		dialog.ShowInformation("Alert", "This is sample alert", w)
	})
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, b, nil, nil,
			),
			l, b,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(350, 250))
	w.ShowAndRun()
}

func confirm() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hell Fyne!")
	b := widget.NewButton("Click", func() {
		dialog.ShowConfirm("Alert",
			"please check yes!",
			func(f bool) {
				if f {
					l.SetText("OK, thank you.")
				} else {
					l.SetText("oh...")
				}
			},
			w,
		)
	})
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, b, nil, nil,
			),
			l, b,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(350, 250))
	w.ShowAndRun()
}

func custom_dialog() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hell Fyne!")
	e := widget.NewEntry()
	b := widget.NewButton("Click", func() {
		dialog.ShowCustomConfirm(
			"Enter message.",
			"OK",
			"Cancel",
			e,

			func(f bool) {
				if f {
					l.SetText("typed: '" + e.Text + "'.")
				} else {
					l.SetText("no message...")
				}
			},
			w,
		)
	})
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, b, nil, nil,
			),
			l, b,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(350, 250))
	w.ShowAndRun()
}

// ------------------------------------------------------------

// MyEntry is custom entry.
type MyEntry struct {
	widget.Entry
	entered func(e *MyEntry)
}

// NewMyEntry create MyEntry.
func NewMyEntry(f func(e *MyEntry)) *MyEntry {
	e := &MyEntry{}
	e.ExtendBaseWidget(e)
	e.entered = f
	return e
}

// KeyDown is Keydown Event.
func (e *MyEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		fmt.Println("return called.")
		e.entered(e)
	default:
		fmt.Println("keydown called.")
		e.Entry.KeyDown(key)
	}
}

func entry() {
	a := app.New()
	w := a.NewWindow("hello")
	l := widget.NewLabel("Hello Fyne")
	e := NewMyEntry(func(e *MyEntry) {
		s := e.Text
		e.SetText("")
		l.SetText("you type '" + s + "',")
	})
	w.SetContent(
		widget.NewVBox(l, e),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(300, 100))
	w.ShowAndRun()

}
