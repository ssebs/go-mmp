# Go-MMP
MacroPad driver software, written in Golang. 

If you have an arduino powered device, you could use this to run various keyboard shortcuts and other macros.

> FYI: This is a re-write of https://github.com/ssebs/MiniMacroPad/

<!-- TODO: Add video of it -->

## Usage
- Connect an arduino/serial based device that sends [0-9] numbers over serial
- Run this GUI app
- It will generate a config for you at `$HOME/mmp-config.yml`
  - You can edit this config from a text editor.
  - > GUI editor coming soon™
  - This is where you configure:
    - The MacroPad's layout (3x3 buttons)
    - The serial device info (portname, buad, etc)
    - The Macros themselves
      - The digit is used to position the macros, so keep the 1,2,3,4... for each macro you want.
      - Name:
        - This is what shows up on the buttons in the GUI. You can use emojis here.
      - Actions:
        - List of Actions, see [Actions](#actions)
- When you press a button on the MacroPad, it should run the macro.
- TBD: CLI flags

## Building
- git clone https://github.com/ssebs/go-mmp
  - You'll need to install a C compiler. See https://developer.fyne.io/started/
  - If you want to use the Makefile on Windows, [install make from choco](https://stackoverflow.com/a/57042516)
- `make build` to build the go-mmp.exe file
- `make run` to run the go-mmp.exe file
- `make test` to run unit tests
- `make pkg` to package for Windows
  - For other platforms:
    - Make sure `fyne` CLI is installed
      - `go install fyne.io/fyne/v2/cmd/fyne@latest`
    - Windows:
      - `PS go-mmp> fyne package -os windows`
    - Mac:
      - `$ fyne package -os darwin`
    - Linux:
      - `$ fyne package -os linux`

## Goals / To-do
- [x] Get started
- [x] Create basic UI
- [x] Run keyboard macros
- [x] Load buttons from config into GUI
- [x] Listen for Serial data
  - [x] Take action from this data
  - [x] Add support for shortcuts/hotkeys
  - [x] Add support for running single keys
  - [x] Add support for strings
  - [ ] Clear whatever is coming in when we first open the app
    - [ ] !if someone pressed a button before opening the app, it will run it asap!
  - [ ] Add support for LMB/RMB/MMB clicks
    - [ ] Support repeat press while holding btn down
  - [ ] Add support for mouse?
- [WIP] Error handling
  - [x] GUI
  - [x] fix serial device errors (serial port busy, nonexistent, etc)
  - [ ] more!
  - [ ] Figure out the "Fyne error:  GLFW poll event error: InvalidValue: Invalid scancode 144"
- [ ] Add unit tests!
  - Current Coverage: 
  - [ ] config tests
    - [x] Constructor makes the same config as a manual file open
    - [ ] LoadConfig actually loads values, not just 0's (check Marshalling)
    - [ ] Validation of values
  - [ ] keyboard tests
    - [ ] Mock keyboard?
    - [ ] what shortcuts turn into
      - [ ] what happens if you want to use "+"
  - [ ] macro tests
    - [ ] load config
    - [ ] confirm macro ids match
    - [ ] confirm function names exist, etc
  - [ ] serialdevice tests
    - [ ] mock?
  - [ ] utils tests
  - [ ] gui tests?
- [ ] CLI flag for GUI only, and non-GUI modes
- [ ] GUI button should depress when matching Serial btn is pressed
- [ ] GUI button should run the macro
- [ ] Make the config file easier to edit
  - [ ] Support more than 1 device?
  - [ ] UI for CRUD'ing these macros
  - [ ] Save default config to $HOME/mmp-config.yml
- [ ] Move main.go to a `cmd` pkg

## Actions:
- PressKey: (string)
  - e.g. `VK_ENTER`
  - Press & release a key.
  - The keyname must be found in `keyboard/keymap.go`
- SendString: (string)
  - e.g. `cool`
  - Type a string of keys
- Shortcut: (string)
  - e.g. `CTRL+SHIFT+ESC`
    - Hotkey sequence, split up by "+" chars.
  - The keys between the "+"'s must be found in `keyboard/keymap.go`
- Delay: (durationString)
  - e.g. `10ms`
  - A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
  - This will add a delay

## Architecture Diagram
> To update it, edit the [Architecture.drawio](./res/Architecture.drawio) file. I'm using [this](https://open-vsx.org/extension/hediet/vscode-drawio) VSCode extension.
> 
![Architecture diagram](./res/Architecture.png).


## Docs / References:
- GUI (fyne)
  - https://developer.fyne.io/
  - https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
- Serial
  - https://github.com/bugst/go-serial
- Keyboard
  - https://github.com/micmonay/keybd_event
- Existing thing I want to improve
  - https://github.com/ssebs/MiniMacroPad/
- For testing macros, check out https://keyboard-test.space/


## Install from exe
> This is TBD
<!-- - go install github.com/ssebs/go-mmp -->

## LICENSE
[Apache 2 License](./LICENSE)
