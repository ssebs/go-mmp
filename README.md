# Go-MMP
MiniMacroPad driver software, written in Golang. 

This is a re-write of https://github.com/ssebs/MiniMacroPad/

## Goals
- [ ] Listen for Serial data
  - [ ] Take action from this data
- [ ] Run keyboard macros
- [ ] CRUD a config file
  - [ ] Support many macros/devices
- [ ] UI for CRUD'ing these macros
- [ ] Support clicking on the button
- [ ] Support keyboard shortcuts?
- [ ] CLI for GUI only, and no-GUI modes

## LICENSE
[Apache 2 License](./LICENSE)

## Docs:
- GUI (fyne)
  - https://developer.fyne.io/
  - https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
- Serial
  - https://github.com/bugst/go-serial
- Keyboard
  - https://github.com/micmonay/keybd_event
