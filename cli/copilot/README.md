## How to install GitHub Copilot CLI?
Tutorial is here: https://www.youtube.com/watch?v=pw0SH7AHIFI

1. Install tool first ([docs here](https://www.npmjs.com/package/@githubnext/github-copilot-cli))
```
npm install -g @githubnext/github-copilot-cli
```
2. Authenticate
```
github-copilot-cli auth
```
3. Go to source file and register new commands
```
eval "$(github-copilot-cli alias -- "$0")"
# Don't forget to source file
```

There are 2 different CLIs (see this comment: https://github.com/orgs/community/discussions/75367#discussioncomment-8242181)

For organizations the one is: https://docs.github.com/en/copilot/github-copilot-in-the-cli/using-github-copilot-in-the-cli
- [Installation](https://docs.github.com/en/copilot/github-copilot-in-the-cli/installing-github-copilot-in-the-cli)
