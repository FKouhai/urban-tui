# URBAN-TUI
This project was done to toy around with [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss). <br>
It allows you to query the urban dictionary api and get a definition of the slang you want to know the meaning of.
## How to use it
You can download the tarball from the release according to your OS and CPU architecture <br>
```(bash)
tar xvf urban-tui_tag_os_arch.tar.gz
./urban-tui
```
### Keymaps
* Enter -> Searches the term
* Ctrl + n -> To search for a new term
* Ctrl + c/q -> Exit out of the program

## Install using nix
To install this package in nix, the provided flake can be used.
Firstly it is needed to add the input and output in the flake.nix of the destination system

```(nix)
inputs = {
  nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  urban-tui = {
    url  ="github:FKouhai/urban-tui/main";
  };
  --- snip ---
  outpus = { self,nixpkgs,urban-tui }@ inputs:
  {
  --- snip ---
  }
};
```
Then inside the configuration.nix file add the following
```(nix)
{ --- snip ---, inputs, ... }:
--- snip ---
environment.systemPackages = [
    inputs.urban-tui.packages."x86_64-linux".default
];
--- snip ---
```

![demo gif](./demo.gif)


