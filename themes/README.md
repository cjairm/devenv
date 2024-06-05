# Terminal themes (+fonts)

## Fonts

Copy the content of `./fonts/` and move it to `~/Library/Fonts`. Then use it from iterm

## Color themes

Download the dracula.itermcolors file

## Spaceship ZSH

[Spaceship ZSH](https://github.com/spaceship-prompt/spaceship-prompt) is a beautiful and customizable ZSH prompt theme that provides a clean and modern look for your terminal. It's highly configurable, allowing you to customize the appearance and behavior of your prompt to suit your needs.

Spaceship ZSH includes a wide range of features, such as:

- Git integration, showing the current branch and status
- Symbolic link indication
- Current directory and username display
- Customizable prompt symbols and colors
- Battery and hostname display

To install Spaceship ZSH, you can simply run the following command in your terminal:

```bash
curl -L https://raw.githubusercontent.com/spaceship-prompt/spaceship-prompt/master/install.sh | sh
```

This will install the theme and set it as the default prompt for your ZSH shell. You can then customize the theme by editing the ~/.zshrc file and adding or modifying the various configuration options.

Overall, Spaceship ZSH is a great choice for anyone looking to customize their ZSH prompt and make their terminal look more modern and visually appealing.

## Powerlevel10k

[Powerlevel10k](https://github.com/romkatv/powerlevel10k#oh-my-zsh) is a fast and highly customizable ZSH prompt theme that provides a clean and modern look for your terminal. It's designed to be highly configurable, allowing you to customize the appearance and behavior of your prompt to suit your needs.

Powerlevel10k includes a wide range of features, such as:

- Git integration, showing the current branch and status
- Symbolic link indication
- Current directory and username display
- Customizable prompt symbols and colors
- Battery and hostname display
- Customizable prompt segments, allowing you to add or remove any information you want to display

To install Powerlevel10k, you need to have the Oh My ZSH framework installed. If you don't have it yet, you can install it by running the following command:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

Once you have Oh My ZSH installed, you can install Powerlevel10k by running the following command:

```bash
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k
```

This will install the theme and set it as the default prompt for your ZSH shell. You can then customize the theme by editing the ~/.zshrc file and modifying the ZSH_THEME variable.

Overall, Powerlevel10k is a great choice for anyone looking to customize their ZSH prompt and make their terminal look more modern and visually appealing. With its wide range of features and high configurability, you can create a prompt that's tailored to your needs and preferences.

### For more suggestions
https://www.youtube.com/watch?v=CF1tMjvHDRA
