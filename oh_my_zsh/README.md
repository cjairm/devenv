## Enhancing Your ZSH Experience with Oh My ZSH

With the release of macOS Catalina 10.15.2, ZSH became the default shell,
replacing Bash. To make the most of ZSH, you can use the powerful Oh My ZSH
framework, which offers various features to increase your productivity.

1. Follow the installation instructions [here](https://ohmyz.sh/).
    - [Themes](https://github.com/ohmyzsh/ohmyzsh/wiki/Themes)
2. Run `devenv configure --oh-my-zsh`
    - Confirm commands were added with `cat ~/.zshrc`
3. Run `source ~/.zshrc`

One of the most convenient features of Oh My ZSH is its autocompletion
functionality. By pressing the `Tab` key, you can quickly select from a list of
available directories, commands, and files. This feature not only saves time
but also helps reduce typing errors.

Another useful feature is the ability to create and use alias commands. You can
view a list of all available aliases by running the alias command in your terminal.
Oh My ZSH also provides several convenient shortcuts, such as omitting the `cd`
(change directory) command when navigating directories. For instance, you can
use `..` instead of `cd ..`, and `~` for the home directory.

Oh My ZSH also includes several time-saving commands, such as the `take` command,
which creates a new directory and changes the path to it in one step. For example,
running `take testFolder` is equivalent to running `mkdir testFolder && cd
testFolder`.

Another useful feature is the ability to quickly navigate between your last and
current path using the `-` command. Oh My ZSH also offers many visually appealing
themes and a list of amazing plugins that can further enhance your ZSH experience.

In addition, Oh My ZSH provides Git integration, allowing you to easily access
Git commands and information without leaving your terminal. With its many features
and customization options, Oh My ZSH is an excellent tool for anyone looking to
enhance their ZSH experience and boost their productivity.

### Install zsh-autosuggestions

``` bash
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
```

### Install zsh-syntax-highlighting

``` bash
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
```

#### Note

Open the ”~/.zshrc” file in your desired editor and modify the plugins line to what you see below.
``` bash
plugins=(git zsh-autosuggestions zsh-syntax-highlighting web-search)
source ~/.zshrc
```
