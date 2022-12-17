package dialogs

import (
	"github.com/gdamore/tcell/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"
	"github.com/rs/zerolog"
)

var _ = Describe("command dialog", Ordered, func() {
	var app *tview.Application
	var screen tcell.SimulationScreen
	var cmdDialog *CommandDialog
	var cmdTitle [][]string
	var runApp func()

	BeforeAll(func() {
		cmdTitle = [][]string{
			{"cmd01", "cmd01 description"},
			{"cmd02", "cmd02 description"},
		}
		app = tview.NewApplication()
		cmdDialog = NewCommandDialog(cmdTitle)

		screen = tcell.NewSimulationScreen("UTF-8")
		err := screen.Init()
		if err != nil {
			panic(err)
		}
		runApp = func() {
			if err := app.SetScreen(screen).SetRoot(cmdDialog, false).Run(); err != nil {
				panic(err)
			}
		}
		zerolog.SetGlobalLevel(zerolog.Disabled)
		go runApp()
	})

	It("display", func() {
		cmdDialog.Display()
		Expect(cmdDialog.IsDisplay()).To(Equal(true))
	})

	It("set focus", func() {
		app.SetFocus(cmdDialog)
		hasFocus := cmdDialog.HasFocus()
		Expect(hasFocus).To(Equal(true))
	})

	It("set rect", func() {
		x := 0
		y := 0
		width := 50
		height := 20
		ws := (width - cmdDialog.width) / 2
		hs := ((height - cmdDialog.height) / 2)
		yWants := y + hs
		xWants := x + ws
		cmdDialog.SetRect(x, y, width, height)
		x1, y1, w1, h1 := cmdDialog.Box.GetRect()
		Expect(x1).To(Equal(xWants))
		Expect(y1).To(Equal(yWants))
		Expect(w1).To(Equal(cmdDialog.width))
		Expect(h1).To(Equal(cmdDialog.height))
	})

	It("get total command counts", func() {
		// header + items (name:desc)
		Expect(cmdDialog.GetCommandCount()).To(Equal(3))
	})

	It("get selected item", func() {
		app.QueueEvent(tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone))
		app.Draw()
		Expect(cmdDialog.GetSelectedItem()).To(Equal("cmd02"))
	})

	It("command selected", func() {
		enterButton := "initial"
		enterButtonWants := "enter selected"
		enterFunc := func() {
			enterButton = enterButtonWants
		}
		cmdDialog.SetSelectedFunc(enterFunc)
		app.QueueEvent(tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone))
		app.Draw()
		Expect(enterButton).To(Equal(enterButtonWants))
	})

	It("cancel button selected", func() {
		cancelButton := "initial"
		cancelButtonWants := "cancel selected"
		cancelFunc := func() {
			cancelButton = cancelButtonWants
		}
		cmdDialog.SetCancelFunc(cancelFunc)
		app.SetFocus(cmdDialog.form)
		app.QueueEvent(tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone))
		app.Draw()
		Expect(cancelButton).To(Equal(cancelButtonWants))
	})

	It("hide", func() {
		cmdDialog.Hide()
		Expect(cmdDialog.IsDisplay()).To(Equal(false))
	})

	AfterAll(func() {
		app.Stop()
	})

})
