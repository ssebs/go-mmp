# Mini Macro Pad (go-mmp)
[![Go](https://github.com/ssebs/go-mmp/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ssebs/go-mmp/actions/workflows/go.yml)

Run Macros, shortcuts, and more at the press of a button. If you have a 3D Printer and some soldering knowledge, you can get yourself a Mini Macro pad!

> No device? No problem! You can click on the buttons to run the Macros too!

- [Mini Macro Pad (go-mmp)](#mini-macro-pad-go-mmp)
  - [Hardware](#hardware)
  - [3D Printed housing](#3d-printed-housing)
  - [The GUI](#the-gui)
  - [What kind of macros can you make?](#what-kind-of-macros-can-you-make)
    - [You can add multiple "actions" to a macro](#you-can-add-multiple-actions-to-a-macro)
  - [Getting started](#getting-started)
  - [Change your Macros on the fly](#change-your-macros-on-the-fly)
  - [Got an arduino all hooked up?](#got-an-arduino-all-hooked-up)
  - [Don't have an arduino but still want to run macros?](#dont-have-an-arduino-but-still-want-to-run-macros)
  - [The following Actions are available:](#the-following-actions-are-available)
      - [PressRelease](#pressrelease)
      - [Press](#press)
      - [Release](#release)
      - [SendText](#sendtext)
      - [Shortcut](#shortcut)
      - [Delay](#delay)
      - [Repeat](#repeat)
  - [Installing binary version](#installing-binary-version)
  - [Install dependencies \& get the code](#install-dependencies--get-the-code)
  - [Build and run the code](#build-and-run-the-code)
  - [LICENSE](#license)


## Hardware
You'll need a microcontroller, some key switches, and a 3D Printer. I'm using a Teensy LC, but you could use an Arduino Micro or ESP32.

Pic of mine below:

![Macro Pad](./res/mmpbuilt.png)

Wiring under the hood:
> Please forgive the newbie soldering!

![Wiring](./res/mmpwiring.png)

## 3D Printed housing
There are many available, but if you like the one I designed, check out my [thangs.com](https://than.gs/m/710028) profile.


## The GUI

Here's what the GUI looks like, you can click the buttons to run the macro, or use the arduino to press them.
![screenshot of gui](res/GUIScreenshot.png)

Most of my keybinds are for an FPS shooter, for example typing "gg" in the chat, but you can automate all sorts of things!


## What kind of macros can you make?
- Shortcuts:
  - CTRL + C, CTRL + V, etc.
- Press whatever key you want, as long as it's [in the list](https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys).
  - Skip song, type "enter", etc.
- Repeat keypresses (or mouse button presses)
  - Playing cookie clicker? Press your macro to repeatedly press your mouse button down until you click the macro again
- Whatever you can think of, feel free to submit PRs!

### You can add multiple "actions" to a macro
If you want a single button to type "gg" for you in VALORANT or CS, you can!
- You just need to add 3 actions:
  - `Shortcut: shift+enter`
  - `SendText: gg`
  - `PressRelease: enter`


## Getting started
If you have a arduino/serial macro pad ready, great! You get to use the full functionality of go-mmp.

If not, you can still run Macros at the press of a button.

- [Download the exe](https://github.com/ssebs/go-mmp/releases/)
- Double click the **go-mmp.exe** file
- It will generate a config for you in your home folder.
    - e.g. `C:\Users\ssebs\mmpConfig.yml` or `/home/ssebs/mmpConfig.yml`
- When you press a button on the Arduino based MacroPad, it will run a Macro that's set in your config.


## Change your Macros on the fly
New in `v2`, you can now update your Macros in the UI instead of from the config file.

![ConfigEditor](./res/ConfigEditor.png)

Just go to Edit > Edit Config and Drag and Drop your macros into the right positions, and click on the name to change what they do.

Here's the "gg" Macro for example.

![MacroEditor](./res/MacroEditor.png)


## Got an arduino all hooked up?
- You'll need an arduino/serial based device that sends [0-9] numbers over a serial connection.
  - See the [arduino-mmp.ino](./arduino-mmp.ino) source code.
  - > Connecting and understanding baudrate, etc. is out of the scope of this guide.

Just edit your config, edit metadata, and set the Serial Port Name, Baud rate, and change `GUIMode` to `NORMAL`.

> if your device sends 1 for the first button instead of 0, you can set the Indexing setting to 1

![MetadataEditor](./res/MetadataEditor.png)

## Don't have an arduino but still want to run macros?
You can still run this in GUI only mode, This is the default so you're all set! 

Just click on the buttons to run Macros.

![GUIScreenshot](./res/GUIScreenshot.png)

## The following Actions are available:
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

## Installing binary version
- Download latest release from https://github.com/ssebs/go-mmp/releases
  - Download either the `.exe.zip `for Windows or `.tar.xz` if you're on Linux.
- If you have fyne installed and setup, you can run
  - `go run github.com/ssebs/go-mmp@latest` to run
  - `go install github.com/ssebs/go-mmp@latest` to install 

## Install dependencies & get the code
- Install [Golang](https://go.dev/doc/install)
- Follow the install docs at https://docs.fyne.io/started/
  - This takes a while but is required.
- If running on linux
  - Add yourself to the `dialout` group
    - `sudo usermod -aG dialout <username>`
  - Install GTK3-dev
    - > This is for the file dialogs
    - `sudo apt install libgtk-3-dev`
- Source:
  - `git clone https://github.com/ssebs/go-mmp`
- Go pkg
  - `go get github.com/ssebs/go-mmp`

## Build and run the code
- Running the code:
  - `go run main.go` 
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
        - `fyne-cross windows`
  - Mac:
    - `$ fyne package -os darwin`
  - Linux:
    - `$ fyne package -os linux`
- Updating:
  - Make code changes 
  - Run upgrades: `go get -u && go mod tidy`
  - Update Version in `FyneApp.toml`
  - Create Pull Request
  - Once committed, git tag & push with same version from `FyneApp.toml`
  - Create release and upload exe and tar.xz 
  - `go get github.com/ssebs/go-mmp@<version>`
- Testing with a virtual serial device (Linux)
  - If you don't have a physical serial device, you can still simulate button presses.
  - Install `socat`, and open up two terminals
  - `$ socat -d -d PTY,raw,echo=0 PTY,raw,echo=0`
    - You'll get two `/dev/pts/<num>` devices listed, the first one is what you set in your config, and the second is used to send data.
  - `$ echo -n "<num>" > /dev/pts/num2`
    - e.g. use num 0 to press the first button

> If you're curious, check out the older python code at https://github.com/ssebs/MiniMacroPad/

## LICENSE
[Apache 2 License](./LICENSE)
