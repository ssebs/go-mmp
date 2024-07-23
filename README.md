# Mini Macro Pad (go-mmp)
[![Go](https://github.com/ssebs/go-mmp/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ssebs/go-mmp/actions/workflows/go.yml)

Simplify task automation with an arduino and some 3D printing. This lets you create shortcuts and run them at the press of a button, customizable through a YAML config file.

> No device? No problem! You can still click on the buttons to run the macros.

Here's what the GUI looks like, you can click the buttons to run the macro, or use the arduino to press them.
![screenshot of gui](res/GUIScreenshot.png)

Most of my keybinds are for an FPS shooter, for example typing "gg" in the chat. 

## What kind of macros can you make?
- Shortcuts:
  - CTRL + C, CTRL + V, etc.
- Press whatever key you want, as long as it's [in the list](https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys).
  - Skip song, type "enter", etc.
- Repeat keypresses (or mouse button presses)
  - Playing cookie clicker? Press your macro to repeatedly press your mouse button down until you click the macro again
- Whatever you can think of, feel free to submit PRs!

### You can add multiple "actions" to a macro
If you want a single button to type "ggez" for you in VALORANT or CS, you can!
- You just need to add 3 actions:
  - Shortcut: shift+enter
  - SendText: ggez
  - PressRelease: enter

## Hardware
You'll need an arduino with some buttons. I'm using a Teensy LC, but you could use an Arduino Micro or ESP32.

Pic of mine below:
![Macro Pad](./res/mmpbuilt.png)

Wiring under the hood
![Wiring](./res/mmpwiring.png)

## 3D Printed housing
There are many available, but if you like the one I designed, check out my [thangs.com](https://than.gs/m/710028) profile.

<!-- TODO: Add video of it -->

## Getting started
- You'll need an arduino/serial based device that sends [0-9] numbers over a serial connection.
  - See [arduino-mmp.ino](./arduino-mmp.ino) source code to see how I did this.
  - Connecting and understanding baudrate, etc. is out of the scope of this guide.
- Download the app
  - [Download go-mmp.exe](https://github.com/ssebs/go-mmp/releases/)
  - [Build from source](#building)
- Double click the **go-mmp.exe** file
- It will generate a config for you at `$HOME/mmpConfig.yml`
    - e.g. `C:\Users\ssebs\mmpConfig.yml` or `/home/ssebs/mmpConfig.yml`
  - You can see what will be generated by looking at the [defaulyConfig.yml file](./res//defaultConfig.yml)
    - Take a look here to see what the file format should look like
  - You can edit this config from a text editor like Notepad, VSCode, or vim
    - All config changes must be made in this file, **this is how you create macros**
    - If someone would like to create a GUI based editor, please feel free 😁
- You can also run `go-mmp.exe` with a `mmpConfig.yml` file in the same directory to load from that config instead of the `$HOME` file
- In the config file, this is where you configure:
  - The MacroPad's layout (3x3 buttons)
  - The serial device info (portname, buad, etc)
    - You'll need to find out what the Serial port name is
  - The Macros themselves
    - The digit is used to position the macros, so keep the 1,2,3,4... for each macro you want.
    - Name:
      - This is what shows up on the buttons in the GUI. You can use emojis here.
    - Actions:
      - List of Actions, see [Actions](#actions) section below
- When you press a button on the Arduino based MacroPad, it should run the macro.
- You can also click the button in the UI to run the macro.

### Don't have an arduino but still want to run macros?
You can still run this in GUI only mode, but you'll need to open up a terminal

Open a terminal to where the go-mmp.exe file is
- Run `PS> go-mmp.exe --gui-only` and hit enter.

### CLI Usage:
> Not sure why, but the print statements stop working after I export to exe, so no help message.
```
Usage of go-mmp.exe:
  -gui-only
        Open Go-MMP in GUI Only Mode. Useful if you don't have a working arduino.
  -reset-config
        If you want to reset your mmpConfig.yml file.
```

## Actions:
All the available actions are listed below, the format is:
  ```yaml
  Actions:
    # FuncName: parameter
  ```

Example Actions:
  ```yaml
  Actions:
    - Shortcut: SHIFT+ENTER
    - SendText: ggez
    - PressRelease: enter
  ```

### The following Actions are available:
> The keyname must be found in https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys, or if it's a mouse button then it should be "LMB", "RMB", "MMB"

#### PressRelease
- e.g. `enter`, `a`, `alt`, `LMB`
- Press and release a key or mouse button.

#### Press
- e.g. `enter`, `a`, `alt`, `LMB`
- Press and hold a key, to release a key use "Release".

#### Release
- e.g. `enter`, `a`, `alt`, `LMB`
- Release a key from being held.

#### SendText
- e.g. `ggez`, `Thanks, ssebs`, `git-gud`
- Type out text from the keyboard.

#### Shortcut
- e.g. `CTRL+SHIFT+ESC`, `CMD+C`, `CMD+R`
  - Hotkey sequence, split up by "+" chars. All keys between the "+"'s will be pressed at the same time.
  - NOTE: The "Windows" button is "CMD", same for the "Command" button on a Mac.
- Run a shortcut by typing multiple keys at once.

#### Delay
- e.g. `10ms`
- You must enter a duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
  - This will add a delay of the time you select.
  - This is useful if your set of actions are too fast and you need to slow them down.
- 
#### Repeat
- e.g. `LMB+50ms`, `z+500ms`
  - KeyName + delay sequence, split up by "+" chars. Only 1 "+" is allowed, so you MUST only put 1 key/mouse button on the left, and a duration string on the right.
- Press and repeat a key or mouse button over and over with the delay between each press.
  - The delay must be formatted as a duration string like above.

## Building
#### Get the code
- Source:
  - Git clone https://github.com/ssebs/go-mmp
    - Install [Golang](https://go.dev/doc/install)
    - Follow the install docs at https://docs.fyne.io/started/
      - This takes a while but is required to get this working.
    - If running on linux, add yourself to the `dialout` group
      - `sudo usermod -aG dialout <username>`
- Go pkg
  - `go get github.com/ssebs/go-mmp`

#### Build the code
- `go build main.go` to build the go-mmp.exe file
- `go run main.go`  to run the go-mmp.exe file
- Unit tests:
  - `go test ./...`
- Coverage
  - `go test ./...  -coverpkg=./... -coverprofile ./coverage.out`
  - `go tool cover -func ./coverage.out`
- To package:
  - Make sure `fyne` CLI is installed
    - `go install fyne.io/fyne/v2/cmd/fyne@latest`
    - `go install github.com/fyne-io/fyne-cross@latest` for cross platform pkging
  - Windows:
    - `PS go-mmp> fyne package -os windows`
      - Linux pkg for Win:
        - `sudo fyne-cross windows`
  - Mac:
    - `$ fyne package -os darwin`
  - Linux:
    - `$ fyne package -os linux`
- Updates:
  - Make code changes 
  - Run upgrades: `go get -u && go mod tidy`
  - Update Version in `FyneApp.toml`
  - Create Pull Request
  - Once committed, git tag & push with same version from `FyneApp.toml`
  - `go get github.com/ssebs/go-mmp@<version>`

<hr/>

If you're curious, check out the older python code at https://github.com/ssebs/MiniMacroPad/

<details>
  <summary>Todo's</summary>

### Goals / To-do (general)
- [x] Get started
- [ ] Add support for running a script / program on keypress
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
  - [x] Add support for LMB/RMB/MMB clicks
    - [x] Support repeat press while holding btn down
- [WIP] Error handling
  - [x] GUI
  - [x] fix serial device errors (serial port busy, nonexistent, etc)
  - [ ] more!
  - [ ] Figure out the "Fyne error:  GLFW poll event error: InvalidValue: Invalid scancode 144"
- [ ] Spamkey should keep UI lit up until it's pressed again
- [ ] Add more tests!
- [x] CLI flag for GUI only, and non-GUI modes
- [x] GUI button should depress when matching Serial btn is pressed
- [x] GUI button should run the macro
- [ ] Make the config file easier to edit
  - [x] Save default config to $HOME/mmpConfig.yml
  - [ ] Support portable mode (load from ./mmpConfig.yml)
  - [ ] Support more than 1 device?
  - [ ] UI for CRUD'ing these macros
- [ ] Move main.go to a `cmd` pkg
- [ ] wiring diagram
- [ ] better instructions for hardware
- [x] Github actions
</details>

## Docs / References:
- GUI (fyne)
  - https://developer.fyne.io/
  - https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
- Serial
  - https://github.com/bugst/go-serial
- Keyboard & Mouse
  - https://github.com/go-vgo/robotgo
- Existing thing I want to improve
  - https://github.com/ssebs/MiniMacroPad/
- For testing macros, check out https://keyboard-test.space/

<details>
 <summary>Architecture / Flow Diagram</summary>

To update it, edit the [Architecture.drawio](./res/Architecture.drawio) file. I'm using [this](https://open-vsx.org/extension/hediet/vscode-drawio) VSCode extension.
 
![Diagram](./res/Architecture.png)
</details>

## LICENSE
[Apache 2 License](./LICENSE)
