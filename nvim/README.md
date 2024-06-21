# kickstart.nvim

- [Official documentation](https://github.com/nvim-lua/kickstart.nvim)

## Introduction

(This was copied an pasted from "Official documentation")

A starting point for Neovim that is:

- Small
- Single-file
- Completely Documented

**NOT** a Neovim distribution, but instead a starting point for your configuration.

## Installation

### Install Neovim

It's recommended to first clean up existent configuration

```bash
rm -rf ~/.local/share/nvim/lazy && rm -rf ~/.config/nvim && rm -rf ~/.local/share/nvim/
```

Kickstart.nvim targets _only_ the latest
['stable'](https://github.com/neovim/neovim/releases/tag/stable) and latest
['nightly'](https://github.com/neovim/neovim/releases/tag/nightly) of Neovim.
If you are experiencing issues, please make sure you have the latest versions.

### Install External Dependencies

> **NOTE**
> [Backup](#FAQ) your previous configuration (if any exists)

External Requirements:

- Basic utils: `git`, `make`, `unzip`, C Compiler (`gcc`)
- [ripgrep](file://../ripgrep)
- Language Setup:
  - If want to write Typescript, you need `npm`
  - If want to write Golang, you will need `go`
  - etc.

Neovim's configurations are located under the following paths, depending on your OS:

| OS           | PATH                                      |
| :----------- | :---------------------------------------- |
| Linux, MacOS | `$XDG_CONFIG_HOME/nvim`, `~/.config/nvim` |

Clone kickstart.nvim:

<details><summary> Mac Ventura </summary>

```bash
git clone https://github.com/cjairm/devenv.git

devenv configure --nvim
```

</details>

### Post Installation

Start Neovim

```sh
nvim
```

That's it! Lazy will install all the plugins you have. Use `:Lazy` to view
current plugin status.

Read through the `init.lua` file in your configuration folder for more
information about extending and exploring Neovim.

### Getting Started

See [Effective Neovim: Instant IDE](https://youtu.be/stqUbv-5u2s), covering the
previous version. Note: The install via init.lua is outdated, please follow the
install instructions in this file instead. An updated video is coming soon.

#### Examples of adding popularly requested plugins

<details>
  <summary>Adding autopairs</summary>

This will automatically install [windwp/nvim-autopairs](https://github.com/windwp/nvim-autopairs) and enable it on startup. For more information, see documentation for [lazy.nvim](https://github.com/folke/lazy.nvim).

In the file: `lua/custom/plugins/autopairs.lua`, add:

```lua
-- File: lua/custom/plugins/autopairs.lua

return {
  "windwp/nvim-autopairs",
  -- Optional dependency
  dependencies = { 'hrsh7th/nvim-cmp' },
  config = function()
    require("nvim-autopairs").setup {}
    -- If you want to automatically add `(` after selecting a function or method
    local cmp_autopairs = require('nvim-autopairs.completion.cmp')
    local cmp = require('cmp')
    cmp.event:on(
      'confirm_done',
      cmp_autopairs.on_confirm_done()
    )
  end,
}
```

</details>
<details>
  <summary>Adding a file tree plugin</summary>

This will install the tree plugin and add the command `:Neotree` for you. You can explore the documentation at [neo-tree.nvim](https://github.com/nvim-neo-tree/neo-tree.nvim) for more information.

In the file: `lua/custom/plugins/filetree.lua`, add:

```lua
-- Unless you are still migrating, remove the deprecated commands from v1.x
vim.cmd([[ let g:neo_tree_remove_legacy_commands = 1 ]])

return {
  "nvim-neo-tree/neo-tree.nvim",
  version = "*",
  dependencies = {
    "nvim-lua/plenary.nvim",
    "nvim-tree/nvim-web-devicons", -- not strictly required, but recommended
    "MunifTanjim/nui.nvim",
  },
  config = function ()
    require('neo-tree').setup {}
  end,
}
```

</details>

## Note
If you want to add copilot ([docs for neovim](https://docs.github.com/en/copilot/using-github-copilot/using-github-copilot-code-suggestions-in-your-editor?tool=vimneovim) and [tutorial here](https://www.youtube.com/watch?v=c9y7bKk-R7U))
```lua
  -- Copilot
  {
    'github/copilot.vim',
  },
```
- [Create PR summaries](https://www.youtube.com/watch?v=LlKjlaHGMr4)
- [Use natural language](https://www.youtube.com/watch?v=ZDbk5M4hbEI)
- [Using GitHub Copilot Chat in your IDE](https://docs.github.com/en/copilot/github-copilot-chat/copilot-chat-in-ides/using-github-copilot-chat-in-your-ide)
- [Creating a pull request summary with GitHub Copilot](https://docs.github.com/en/enterprise-cloud@latest/copilot/github-copilot-enterprise/copilot-pull-request-summaries/creating-a-pull-request-summary-with-github-copilot)
