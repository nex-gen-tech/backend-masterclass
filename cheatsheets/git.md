# Git Cheat Sheet

1. git init - initialize a new Git repository

- `git init my-project` - create a new Git repository in a directory called "my-project"
- `git init` . - create a new Git repository in the current directory
- `git init --bare` - create a new bare repository (used for sharing code with others)

2. git clone - create a local copy of a remote repository

- `git clone git-url` - create a local copy of the remote repository at **[https://github.com/user/repo.git](https://github.com/user/repo.git)**
- `git clone git-url my-project` - create a local copy of the remote repository in a directory called "my-project"
- `git clone --depth 1 git-url` - create a shallow clone of the remote repository, only pulling the latest commit

3. git add - stage changes for commit

- `git add file1 file2` - stage changes to file1 and file2 for commit
- `git add -A` - stage all changes, including deleted files, for commit
- `git add .` - stage all changes, excluding deleted files, for commit

4. git commit - commit staged changes

- `git commit -m "Commit message"` - commit staged changes with a commit message
- `git commit -am "Commit message"` - stage and commit all changes, including modified and deleted files, with a commit message
- `git commit --amend` - add changes to the previous commit, instead of creating a new commit

5. git push - push commits to a remote repository

- `git push origin master` - push commits to the "master" branch of the remote repository named "origin"
- `git push origin feature-branch` - push commits to a branch called "feature-branch" on the remote repository named "origin"
- `git push -f origin master` - force push commits to the "master" branch, overwriting any remote changes

6. git pull - fetch and merge changes from a remote repository

- `git pull origin master` - fetch and merge changes from the "master" branch of the remote repository named "origin"
- `git pull --rebase origin master` - fetch and rebase (apply local commits on top of) changes from the "master" branch of the remote repository named "origin"
- `git pull --all` - fetch and merge changes from all branches of the remote repository

7. git fetch - fetch changes from a remote repository

- `git fetch origin` - fetch changes from the remote repository named "origin"
- `git fetch --all` - fetch changes from all remote repositories
- `git fetch --prune` - fetch changes and delete any branches that have been deleted on the remote repository

8. git merge - merge changes from another branch

- `git merge feature-branch` - merge changes from the "feature-branch" branch into the current branch
- `git merge --no-ff feature-branch` - merge changes and create a new commit, even if it's a fast-forward merge
- `git merge --squash feature-branch` - merge changes, but squash them into a single commit

9. git branch - manage branches

- `git branch feature-branch` - create a new branch called "feature-branch"
- `git branch -d feature-branch` - delete the "feature-branch" branch
- `git branch -m new-branch` - rename the current branch to "new-branch"

10. git checkout - switch to a different branch or restore files

- `git checkout feature-branch` - switch to the "feature-branch" branch
- `git checkout HEAD -- file1 file2` - restore file1 and file2 to their latest commit
- `git checkout -b new-branch` - create and switch to a new branch called "new-branch"

11. git stash - temporarily save changes that are not ready to be committed

- `git stash` - save changes and revert to the latest commit
- `git stash save "Stash message"` - save changes with a stash message
- `git stash apply` - apply the latest stash, but keep the stash

12. git tag - create and manage tags

- `git tag v1.0` - create a new tag called "v1.0" at the current commit
- `git tag -a v1.0 -m "Tag message"` - create an annotated tag called "v1.0" with a tag message
- `git tag -d v1.0` - delete the tag "v1.0"

13. git reset - undo commits and move the branch pointer

- `git reset HEAD~~1` - undo the latest commit and move the branch pointer to the previous commit
- `git reset --hard HEAD~~2` - undo the last two commits, discarding any changes, and move the branch pointer to the previous commit
- `git reset --mixed HEAD~~3` - undo the last three commits, keeping changes staged, and move the branch pointer to the previous commit

14. git rebase - apply commits from one branch onto another

- `git rebase master` - apply commits from the current branch onto the "master" branch
- `git rebase -i HEAD~~3` - interactively rebase the last three commits
- `git rebase --onto new-base old-base branch-name` - rebase "branch-name" onto "new-base", using "old-base" as the starting point

15. git cherry-pick - apply a specific commit from one branch onto another

- `git cherry-pick commit-hash` - apply the commit with the specified commit hash onto the current branch
- `git cherry-pick -x commit-hash` - apply the commit with the specified commit hash onto the current branch and append a cherry-pick notation to the commit message
- `git cherry-pick --abort` - stop the cherry-pick process and restore the previous state

16. git blame - view the commit history of a file

- `git blame file` - display the commit history of file, showing which lines were added or modified by each commit
- `git blame -M file` - detect moved or copied lines in file and show their original commit history
- `git blame -C file` - detect moved or copied lines in file and show their original commit history, even if the lines have been modified

17. git log - view commit history

- `git log` - display the commit history for the current branch
- `git log --oneline` - display the commit history in a single line per commit
- `git log` --author="Author name" - display the commit history for commits by "Author name"

18. `git show` - display a specific commit

- `git show commit-hash` - display the commit with the specified commit hash
- `git show --stat commit-hash` - display the commit with the specified commit hash, including a summary of changes
- `git show --patch commit-hash` - display the commit with the specified commit hash, including the full diff of changes

19. git diff - display differences between commits, branches, or files

- `git diff HEAD~1` - display differences between the current commit and the previous commit
- `git diff feature-branch` - display differences between the current branch and the "feature-branch" branch
- `git diff file1 file2` - display differences between file1 and file2

20. git status - display the status of the repository

- git status` - display the status of the current branch, including staged and unstaged changes
- `git status -s` - display the status in a shorter, more concise format
- `git status --ignored` - display the status, including ignored files

21. git config - manage Git configuration settings

- `git config --global user.name "Your name"` - set the global Git configuration setting for the user's name to "Your name"
- `git config --global user.email "your@email.com"` - set the global Git configuration setting for the user's email to "your@email.com"
- `git config --global --list` - display the global Git configuration settings

22. git remote - manage remote repositories

- `git remote add origin https://github.com/user/repo.git` - add a remote repository named "origin" at https://github.com/user/repo.git
- `git remote remove origin` - remove the "origin" remote repository
- `git remote -v` - display the remote repositories and their URLs

23. git ls-remote - display references in a remote repository

- `git ls-remote origin` - display references in the remote repository named "origin"
- `git ls-remote --heads origin` - display references to branches in the remote repository named "origin"
- `git ls-remote --tags origin` - display references to tags in the remote repository named "origin"

24. git rev-parse - parse revision information

- `git rev-parse HEAD` - display the commit hash for the current commit
- `git rev-parse --short HEAD` - display the abbreviated commit hash for the current commit
- `git rev-parse` --symbolic-full

25. git ls-files - list files in the index and working tree

- `git ls-files` - list all files in the index and working tree
- `git ls-files --stage` - list all files in the index, including their staging information
- `git ls-files --modified` - list all modified files in the working tree

26. git update-index - manage the Git index

- `git update-index --add file1 file2` - add file1 and file2 to the index
- `git update-index --force-remove file` - remove file from the index, even if it has been modified in the working tree
- `git update-index --assume-unchanged file` - mark file as unchanged in the index, ignoring any changes in the working tree

27. git gc - garbage collect loose objects and optimize the repository

- `git gc` - perform garbage collection and optimization on the current repository
- `git gc --aggressive` - perform more aggressive garbage collection and optimization
- `git gc --auto` - automatically perform garbage collection and optimization when necessary

28. git prune - remove unreachable objects from the object database

- git prune` - remove unreachable objects from the current repository
- `git prune --expire 2.weeks.ago` - remove unreachable objects that have not been accessed in the past 2 weeks
- `git prune --dry-run` - show what objects would be pruned, but do not actually prune them

29. git fsck - check the integrity of the Git database

- `git fsck - check` the integrity of the current repository
- `git fsck --full` - perform a more thorough check of the repository
- `git fsck --strict` - check the repository and fail if any errors are found

30. git instaweb - start a web server to display the repository locally

- `git instaweb` - start a web server to display the current repository
- `git instaweb --httpd=webrick` - start a web server using the WEBrick HTTP server
- `git instaweb --stop` - stop the web server
