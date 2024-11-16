# Model View Controller planning doc for data binding.

## Goals
- [ ] Load config data from disk
- [ ] Generate config if needed (default, test, --reset, etc.)
- [ ] Save config data to disk
- [ ] Update config data from GUI
- [ ] Update all GUI elements when config data changes
- [ ] Save config data from GUI to disk
- [ ] `ports, err := serial.GetPortsList()` metadata editor

## Config data has:
- Metadata
  - Layout/size
  - Serial connection info
  - Default delay
  - GUI Mode
- Macros
  - Ordered list of Macros
    - err key/val of Macros, using ints as the key
    - replace this?
- Macro
  - Name/title to display
  - Ordered ist of Actions
- Action
  - Key/val pair
    - Doesn't have to be
  - key: Function to run e.g. `PressRelease`
  - val: Parameter for func e.g. `DELETE`

## Models
- Config is the "main" one that composes other types
- Structs with yaml support
- Have constructors
- Getters and Setters that can be used with Controllers

### Config Model:
```json
{
    "Metadata": null,
    "Macros": [
        null,
    ],
}
```

### Metadata Model:
```json
{
    "Cols": 3,
    "SerialPortName": "COM1",
    "SerialBaudRate": 9600,
    "Delay": "125ms",
    "GUIMode": "NORMAL"
}
```

### Macro Model:
```json
{
    "Name": "Open Task Manager",
    "Actions": [
        null,
    ]
}
```

### Action Model:
```json
{
    "FuncName": "Shortcut",
    "FuncParam": "CTRL+SHIFT+ESC"
}
```

## Controllers 
- Runs functions / acts as binding between model and view

### Action Controller
- `GetFunctionNames()` will return list of all allowed functions that can be used
- `SetFuncName(fn string)`
- `SetFuncParam(fp string)`
- `CheckValidParam(fn, fp string) bool`

### Config Controller
- `SaveConfig(destinationFullPath string)`
- `LoadConfig(sourceFullPath) *Config`
- `AddMacro(newMacro Macro)`
- `DeleteMacro(idx int)`
- `UpdateMacro(idx int, updatedMacro Macro)`
- `GetMacro(idx int)`

### Metadata Controller
- `SetCols(colCount int)`
- `SetSerial(portName string, baud int)`
- `SetDefaultDelay(time.Duration)`
- `SetGUIMode(mode GUIMode)`
- `GetSerialPorts() []string`

### Macro Controller
- `SetName(n string)`
- `AddAction(newAction Action)`
- `DeleteAction(idx int)`
- `UpdateAction(idx int, updatedAction Action)`
- `GetAction(idx int)`

## Views
- GUI widgets should be able to:
  - Send update data to the Controller
  - Recv updates when the Controller says so, and update the UI
    - (callback fn?)
- Drag/Drop type state should not update the data model until "done". 
  - e.g. DragEnd() should tell the controller what to update and how
    - e.g. SwapMacroPositions(idx1, idx2)
- [ ] Display Macros as buttons in correct order in a grid
- [ ] Click on Macro btn will run a macro
- [ ] File menu to:
  - [ ] Open config
  - [ ] Quit
- [ ] Edit menu to:
  - [ ] *Edit Macros* 
  - [ ] *Edit Config*
- *Edit Macros*:
  - [ ] Convert buttons to *Custom Macro Edit* widgets in place of where the buttons were
  - [ ] Drag and Drop the widgets to change position in grid
  - [ ] Cancel Button
    - Undo button would be nice too, but out of scope
  - [ ] Save/Save As button
  - [ ] *Custom Macro Edit* widgets
    - [ ] Draggable using hamburger icon on top left
    - [ ] Delete Macro using `x` in top right
    - [ ] Rename Title/Name from entry
    - [ ] *Edit Actions* button on the bottom
  - [ ] *Edit Actions* widget
    - [ ] New window
    - [ ] Macro name (entry so it can be edited too)
    - [ ] Short tutorial / explanation
    - [ ] List of *ActionItem Editor* widgets
      - [ ] Drag and Drop the widgets to change run order
    - [ ] Save btn
  - [ ] *ActionItem Editor* widget
    - [ ] Draggable using hamburger icon on left
    - [ ] Select with Function to run
    - [ ] Entry with Function Parameter as string
      - [ ] Validation - depends on Function
- [ ] *Edit Config* widget
  - [ ] New window
  - [ ] Edit the Serial conn info
  - [ ] Edit the default delay
  - [ ] Edit the gui mode
  - [ ] Save btn