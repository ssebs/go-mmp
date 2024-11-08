# Model View Controller planning doc for data binding.

## Goals

### Config
- [ ] Load config data from disk
- [ ] Generate config if needed (default, test, --reset, etc.)
- [ ] Save config data to disk
- [ ] Update config data from GUI
- [ ] Update all GUI elements when config data changes
- [ ] Save config data from GUI to disk

### GUI
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

## The Plan
- Create custom struct types for each field in config
- Macro should have methods to update values
- Actions data structure should be changed from list of key/val pairs. it's hard to use

## Models
- Config is the "main" one that composes other types
- Structs with yaml support
- Have constructors

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
- Public functions to:
  - Manipulate data within Model from a View
  - Update Views when something in the Model is updated
  - e.g. for Actions, there should be AddAction() / GetActions() methods
- Should notify the model, and the view that data has been updated
  - Observer pattern?
- Can be in the same file as the model?

### Config Controller
- `notifyAll()`
- `SaveConfig(destinationFullPath string)`
- `LoadConfig(sourceFullPath) *Config`
- `AddMacro(newMacro Macro)`
- `DeleteMacro(idx int)`
- `UpdateMacro(idx int)`
- `Subscribe(id string?, callback func(c *config)?)`

### Metadata Controller

### Macro Controller

### Action Controller


## Views
- GUI widgets should be able to:
  - Send update data to the Controller
  - Recv updates when the Controller says so, and update the UI
    - (callback fn?)
- Drag/Drop type state should not update the data model until "done". 
  - e.g. DragEnd() should tell the controller what to update and how
    - e.g. SwapMacroPositions(idx1, idx2)
