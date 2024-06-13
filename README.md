# `devenv` command

- [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Description

### Automates configuration

This command is designed to help you configure specific applications with ease. It automates the process of setting up custom configurations, so you don't have to do it manually. This can save you time and effort, especially if the application has a complex configuration process.

However, some applications may require manual installation. In such cases, this documentation will provide you with the necessary information to install and configure these applications. It will help you identify which applications require manual installation and provide step-by-step instructions on how to do it. This way, you can still benefit from the automated configuration process for applications that support it, while also being able to manually install and configure other applications that require it.

By providing both automated and manual configuration options, this documentation aims to be a comprehensive resource for configuring a wide range of applications. Whether you prefer a hands-off approach or want to have more control over the configuration process, this documentation has you covered.

### Productivity Configuration for Mac

This is the setup guide for configuring the development environment on a Mac running Ventura. It provides step-by-step instructions for initializing the development environment, including any necessary software installations and configuration settings.

1. [Install Raycast](https://github.com/cjairm/devenv/tree/main/raycast)
1. [Install Homebrew](https://github.com/cjairm/devenv/tree/main/homebrew)
1. [Install iTerm2](https://github.com/cjairm/devenv/tree/main/iterm2)
1. [Enhancing the ZSH Experience with Oh My ZSH](https://github.com/cjairm/devenv/tree/main/oh_my_zsh)
1. [Terminal themes](https://github.com/cjairm/devenv/tree/main/themes)
1. [Ripgrep](https://github.com/cjairm/devenv/tree/main/ripgrep)
1. [Nvim](https://github.com/cjairm/devenv/tree/main/nvim)
1. [Tmux](https://github.com/cjairm/devenv/tree/main/tmux)

## Features

- **Configures zsh**: This command configures the ZSH shell by adding custom configuration settings and customizations, such as aliases and exports. The command can be safely executed multiple times, as it includes checks to prevent duplication. This means that if you run the command multiple times with the same alias or export, it will not create multiple entries for the same setting. This makes it easy to add and modify your ZSH customizations without worrying about duplicates or conflicts.
  - zshconf
  - ohmyzshconf
- **Configures nvim**: ----
- **Configures yabai**: ----
- **Configures skhd**: ----

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

### Installation

By following these steps, you can ensure that the project is installed correctly and can be easily accessed from your command line.

Creating a new folder directory is optional, but it can be helpful if you want to install the project globally, rather than just in your current working directory. This way, you can easily access the project from any directory on your system.

```bash
git clone https://github.com/cjairm/devenv.git "${XDG_CONFIG_HOME:-$HOME/.config}"/devenv

################################################################
### If you want to make it a global command...
echo 'alias devenv="~/.config/devenv"' >> ~/.zshrc
source ~/.zshrc
################################################################
```

### Usage

From here as simple as

```bash
devenv configure --<flag-name>
```

| Flag name           | Default value | Description                                                |
| ------------------- | ------------- | ---------------------------------------------------------- |
| --all, -a           | `false`       | Configures all available services                          |
| --oh-my-zsh, -o     | `false`       | Add none commented lines in `./oh_my_zsh/.extension_zshrc` |
| --nvim, -n          | `false`       | Configures nvim plugins                                    |
| --tmux, -t          | `false`       | Configures tmux                                            |
| --window-manage, -w | `false`       | Cofigures yabai and skhd                                   |

---

Enjoy! :smiley:
