# git-hook-commit-msg

This is a repo with git-hooks AtlantBH challenge written in Go.

# Installation

To clone this repository to your local machine, use the following command in your terminal:

```bash
git clone https://github.com/tepavcevic/git-hook-commit-msg.git
```

After installation run following command in your terminal to set git conf for hooks path:

```bash
./set-git-hooks-path.sh
```

# Using git-hooks

Now you can try and make a new commit to try and test the hook.

Allowed messages will be:
- Less than 50 characters long
- Will include one of the following
    - feature (for new features)
    - fix (for bug fixes)
    - refactor (for refactoring code)
    - wip (for work in progress, commits to be pushed later)


I've included binaries for linux, windows and macOS (I've only been able to test on linux though)

Implementation in Go is pretty straight forward, I even ommited using regex for matching commit types in favor of a predefined string array that can easily be expanded with new keywords.
Go was chosen because I wanted to try something different, and while being new to it, it was a lot of fun trying to do it.
The biggest challenge will be for the binaries to be properly configured so it can be run on different OS's

Even though this would be most suited for bash, I can with utmost confidence say that it works with Go (only on my machine :))