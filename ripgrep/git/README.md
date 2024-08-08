### Useful commands
```bash
git clean -Xfd

git log master..branchName --oneline

git diff-tree -p <commit>
```

### Push code
```bash
git add -p
git commit
git push -u origin HEAD
```

### Put them on the newbranch
```bash
git checkout newbranch
git cherry-pick 612ecb3
git cherry-pick 453ac3d
git cherry-pick 9aa1233
```

### Merging a feature flag
```bash
git checkout master
git fetch
git remote prune origin
git pull origin master
git branch -d feature/my-branch-name
  git checkout main && git fetch && git remote prune origin && git pull origin main && git branch -D feature/tofu-1131-remove-role-check-in-posting-flyover-create
```

### Creating new branch
```bash
git checkout -b feature/my-branch-name
```

### Pull from existing brach
```bash
* git checkout -b branch-name origin/branch-to-clone
```

### Working on feature flag
```bash
# …work and commit some stuff
git merge master
git push origin feature/my-branch-name

# The git merge master assumes the master branch is up to date, so if it is not, you'll have to update master first.
git checkout master
git pull origin master
git checkout feature/my-branch-name
git merge master
```

### Change from master to main
```bash
# git branch -m master main && git fetch origin && git branch -u origin/main main && git remote set-head origin -a

# Switch to the "master" branch:
git checkout master

# Rename it to "main":
git branch -m master main

# Get the latest commits (and branches!) from the remote:
git fetch

# Remove the existing tracking connection with "origin/master":
git branch --unset-upstream

# Create a new tracking connection with the new "origin/main" branch:
git branch -u origin/main

# To unstage a commit
git reset HEAD

# Reset a file to previous commit
git checkout —

# Undo a commit: This moves all changes from last commit back to staging. The caret means move to one commit previous.
git reset —soft HEAD^

# Undo commit and discard all changes
git reset —hard HEAD^

# Add staged files to previous commit
git commit —amend -m “new message”

git log --branches --not --remotes
```
