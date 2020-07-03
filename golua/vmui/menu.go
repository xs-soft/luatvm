package vmui
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
func menu()[]MenuItem{
	return []MenuItem{
		Menu{
			Text: "&File",
			Items: []MenuItem{
				Action{
					//AssignTo:    &openAction,
					Text:        "&Open",
					//Image:       "../img/open.png",
					Enabled:     Bind("enabledCB.Checked"),
					Visible:     Bind("!openHiddenCB.Checked"),
					Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
					//OnTriggered: mw.openAction_Triggered,
				},
				Menu{
					//AssignTo: &recentMenu,
					Text:     "Recent",
				},
				Separator{},
				Action{
					Text:        "E&xit",
					//OnTriggered: func() { mw.Close() },
				},
			},
		},
		Menu{
			Text: "&串口绑定",
			Items: []MenuItem{
				Action{
					Text:    "UART_1",
					Checked: Bind("enabledCB.Visible"),
				},
				Action{
					Text:    "UART_2",
					Checked: Bind("openHiddenCB.Visible"),
				},
			},
		},
		//Menu{
		//	Text: "&Help",
		//	Items: []MenuItem{
		//		Action{
		//			AssignTo:    &showAboutBoxAction,
		//			Text:        "About",
		//			OnTriggered: mw.showAboutBoxAction_Triggered,
		//		},
		//	},
		//},
	}
}