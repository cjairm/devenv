# Install Ripgrep

## Installation

To install ripgrep, follow the official documentation from the [ripgrep GitHub page](https://github.com/BurntSushi/ripgrep).

## Useful Aliases

Here are some useful aliases that you can add to your .zshrc file:

- `alias nvim_sf="fd --type f --hidden --exclude .git | fzf-tmux -p --reverse | xargs nvim"`: This alias allows you to quickly open files in Neovim using fzf-tmux, a fuzzy finder for the terminal. It requires the fd and fzf tools to be installed.
  - To install the required tools for the aliases, run the following commands:
    - `brew install fd`
    - `brew install fzf`
- `alias tmclear="clear && tmux clear-history && clear"`: This alias allows you to clear the history of the current tmux window.

## Usage

To use ripgrep, simply run the rg command followed by the search term and the directory you want to search in. For example, to search for the term "example" in the current directory, run the following command:

```bash
rg cjairm .
```

Overall, ripgrep is a fast and powerful search tool that can help you find what you're looking for quickly and efficiently. By combining it with other tools, such as fd and fzf, you can create even more powerful aliases and customizations for your development environment.

### Useful commands
```bash
rg '// TODO' -C 2

rg calledWith -g "*Test.ts"

rg promisedRun -g '!*Test.ts'

rg -l '// TODO' --sort path

rg --files-without-match '\b(var|let|const)\b'

# Fixed strings
rg -F '?.'

# After context
rg '// TODO' -A 3

# Search specific files
rg -t js '@format'

# get only files
rg -l '// TODO'

# GET ONLY FOLDERS
rg profiles_sensitive -g '!acl' -g '!collections' | sed 's#/.*##' | sort -u

# COUNT CONCURRENCES
rg profiles_sensitive | grep -c ‘profiles_sensitive’
```
