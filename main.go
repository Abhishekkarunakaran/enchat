package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func generateKeys(bitSize int) (*rsa.PublicKey, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, nil, err
	}

	return &privateKey.PublicKey, privateKey, nil
}

func test() {
	pub, pvt, err := generateKeys(4096)
	if err != nil {
		fmt.Println(err.Error())
	}

	privBytes := x509.MarshalPKCS1PrivateKey(pvt)
	privPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})

	// Marshal and encode public key to PEM
	pubBytes := x509.MarshalPKCS1PublicKey(pub)
	pubPem := pem.EncodeToMemory(&pem.Block{
		// Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	})

	fmt.Println("Private Key:")
	fmt.Println(string(privPem))

	fmt.Println("Public Key:")
	fmt.Println(string(pubPem))
}

func setKeybindings(
	page *tview.Flex,
	app *tview.Application,
	userList *tview.List,
	chatWindow *tview.TextView,
	chatInput *tview.InputField) {

		page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

			switch event.Key() {
			case tcell.KeyESC:
				app.Stop()
			case tcell.KeyEnter:
				switch {
				case chatInput.HasFocus():
					//set the message
				}
			default:
				switch event.Name(){
				case "Shift+Left":
					if chatInput.HasFocus(){
						app.SetFocus(userList)
					}
			case "Shift+Right":
				if userList.HasFocus(){
					app.SetFocus(chatInput)
				}
				}
			}
			

			return event
		})
}
func runApp() {
	app := tview.NewApplication()

	listUsers := tview.NewList()

	listUsers.
		ShowSecondaryText(false).
		SetBorder(true).SetBorderPadding(0,0,1,0).
		SetBorderStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	
	chatWindow := tview.NewTextView()

	chatWindow.
			SetBorder(true).
			SetBorderStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)).
			SetBackgroundColor(tcell.ColorBlack).
			SetBorderPadding(0,1,1,0)
	
	chatInput := tview.NewInputField()

	chatInput.
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetBorderStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)).
		SetBackgroundColor(tcell.ColorBlack).
		SetBorderPadding(0,0,1,0)

	chatSide := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chatWindow,0,5,false).
		AddItem(chatInput,0,1,false)

	chatApp := tview.NewFlex().SetDirection(tview.FlexColumn).
	AddItem(listUsers,0,2,true).
	AddItem(chatSide,0,5,false)

	keyList := tview.NewTextView().
	SetTextStyle(tcell.Style.Foreground(tcell.StyleDefault, tcell.ColorGray)).
	SetText("↑/↓: nav • enter: select • esc: switch focus • Ctrl+C: exit")

	keyList.
	SetBackgroundColor(tcell.ColorBlack)
	

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(chatApp,0,20,true).
	AddItem(keyList,0,1,false)

	setKeybindings(layout, app, listUsers, chatWindow,chatInput)
	if err := app.SetRoot(layout, true).Run(); err != nil {
		log.Fatal("Error running application: %v", err)
	}
}
func main() {
	// test()
	runApp()

}
