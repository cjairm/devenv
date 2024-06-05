# Install Tmux

The tmux configuration is used to customize the behavior and appearance of the terminal multiplexer tmux. The configuration file is written in the tmux configuration language and is loaded when tmux starts.

## Installation

```bash
brew install tmux
```

- [Official docs](https://github.com/tmux/tmux/wiki)

Tmux's configurations are located under the following paths, depending on your OS:

| OS           | PATH |
| :----------- | :--- |
| Linux, MacOS | `~/` |

<details><summary> Mac Ventura </summary>

```bash
git clone https://github.com/cjairm/devenv.git

devenv configure --tmux
```

## Configuration

The tmux configuration is divided into several sections, each of which customizes a specific aspect of tmux.

## General

The general section contains configuration options that apply to all tmux sessions.

### Window Title

The window title is set to the current session name and the current window index.

### Status Bar

The status bar is customized to show the current session name, the current window index, and the current pane index.

### Key Bindings

Key bindings are customized to provide a more intuitive and efficient user experience.

## Sessions

The sessions section contains configuration options that apply to specific tmux sessions.

### Default Session

The default session is named main and is created when tmux starts.

### Window Splitting

Window splitting is customized to use the v key for vertical splitting and the h key for horizontal splitting.

### Pane Switching

Pane switching is customized to use the prefix key (default is ctrl+b) followed by the arrow keys.

## Windows

The windows section contains configuration options that apply to specific tmux windows.

### Default Window

The default window is named home and is created when a new tmux session is started.

### Window Layout

The window layout is customized to use a vertical split with two panes, each running a different terminal emulator.

## Panes

The panes section contains configuration options that apply to specific tmux panes.

### Default Pane

The default pane runs the bash shell.

### Pane Title

The pane title is set to the current command or process name.

#### Create new pane

```
Cntrl-Space v -> split vertically
Cntrl-Space h -> split horizontally
```

#### Move from pane to pane

```
Cntrl-h -> Left
Cntrl-j -> Down
Cntrl-k -> Up
Cntrl-l -> Right
```

#### Windows

```
Cntrl-Space + n -> Rename window
Cntrl-Space + w -> Create new
```

#### Copy mode

```
Cntrl-Space + [
```

#### Move window to window

```
Cntrl-u -> Prev window
Cntrl-o -> Next window
```

#### Choose tree

```
Cntrl-t
```

#### Clock

```
Cntrl-Space + t
```
