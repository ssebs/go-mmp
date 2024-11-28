# Actions

An Action is what is actually running when you use a Macro.

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