CLI app launcher for fuzzy finders

# Install
```sh
go install github.com/htsee/fzlaunch
```

If you use Nix, you can install the package using this repo's `flake.nix`

# Usage
## Subcommands
`list`: generate a list of applications

`run [app]`: run an application

`preview [app]`: provide information about an application

## With fzf

```sh
# bash, zsh
fzlaunch list | fzf --preview "fzlaunch preview {}" | xargs -d "\n" fzlaunch run
```

```nu
# nushell
fzlaunch list | fzf --preview "fzlaunch preview {}" | fzlaunch run $in
```

## With television
Put this in `$XDG_CONFIG_HOME/television/cable/fzlaunch.toml`

```toml
[metadata]
name = "fzlaunch"
description = "App launcher"
requirement = ["fzlaunch"]

[source]
command = "fzlaunch list"

[preview]
command = "fzlaunch preview '{}'"

[keybindings]
enter = "actions:open"

[actions.open]
description = "Launch application"
command = "fzlaunch run {}"
mode = "execute"
```

then run `tv fzlaunch`

# Build

```sh
git clone github.com/htsee/fzlaunch

# with go
go build

# with nix
nix build
```
