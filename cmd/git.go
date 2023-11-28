package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	commandFlag  string
	advancedFlag string
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Explains something about Git",
	Long: `This command provides explanations and examples related to Git.
For example:

- Explaining basic Git commands
- Giving extensive information about advanced git features`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(`Git command has no subcommands. Please provide one of the following flags:
--command
--advanced`)
			return
		}

		command, _ := cmd.Flags().GetString("command")
		if command != "" {
			explainGitCommand(command)
			return
		}

		advanced, _ := cmd.Flags().GetString("advanced")
		if advanced != "" {
			explainAdvancedGitConcepts(advanced)
			return
		}

		fmt.Println("Git is a distributed version control system that tracks changes in any set of computer files, usually used for coordinating work among programmers who are collaboratively developing source code during software development.")
		fmt.Println("- git init: Initialize a new Git repository.")
		fmt.Println("- git add: Add changes to the staging area.")
		fmt.Println("- git commit: Record changes to the repository.")
		fmt.Println("- git status: Show the status of changes as untracked, modified, or staged.")
		fmt.Println("- git branch: List, create, or delete branches.")
		fmt.Println("- git merge: Merge changes from different branches.")
		fmt.Println("- git pull: Fetch from and integrate with another repository or a local branch.")
		fmt.Println("- git push: Update remote refs along with associated objects.")
		fmt.Println("- git log: Display commit history.")
		fmt.Println("- git clone: Clone a repository.")
		fmt.Println("- git remote: Manage remote repositories.")
		fmt.Println("- git fetch: Fetch changes from a remote repository without merging.")
		fmt.Println("- git reset: Unstage changes or reset the repository to a previous state.")
		fmt.Println("- git tag: Create and manage tags for releases.")
	},
}

func init() {
	gitCmd.Flags().StringVarP(&commandFlag, "command", "c", "", "Specify a Git command to explain")
	gitCmd.Flags().StringVarP(&advancedFlag, "advanced", "a", "", "Explain advanced Git concepts")

	gitCmd.SetUsageTemplate(usageTemplate())
}

func explainGitCommand(command string) {
	switch command {
	case "init":
		fmt.Println("The 'git init' command initializes a new Git repository.")
		fmt.Println("Example: git init")
		fmt.Println("This command initializes a new Git repository in the current directory.")
	case "add":
		fmt.Println("The 'git add' command adds changes to the staging area.")
		fmt.Println("Example: git add file.txt")
		fmt.Println("This command stages the changes in 'file.txt' for the next commit.")
	case "commit":
		fmt.Println("The 'git commit' command records changes to the repository.")
		fmt.Println("Example: git commit -m 'Add a new feature'")
		fmt.Println("This command creates a new commit with a message describing the changes.")
	case "status":
		fmt.Println("The 'git status' command shows the status of changes as untracked, modified, or staged.")
		fmt.Println("Example: git status")
		fmt.Println("This command displays the current state of the working directory and staging area.")
	case "branch":
		fmt.Println("The 'git branch' command lists, creates, or deletes branches.")
		fmt.Println("Example: git branch feature-branch")
		fmt.Println("This command creates a new branch named 'feature-branch'.")
	case "merge":
		fmt.Println("The 'git merge' command merges changes from different branches.")
		fmt.Println("Example: git merge feature-branch")
		fmt.Println("This command merges the changes from 'feature-branch' into the current branch.")
	case "pull":
		fmt.Println("The 'git pull' command fetches from and integrates with another repository or a local branch.")
		fmt.Println("Example: git pull origin main")
		fmt.Println("This command fetches changes from the 'main' branch on the remote repository and merges them into the current branch.")
	case "push":
		fmt.Println("The 'git push' command updates remote refs along with associated objects.")
		fmt.Println("Example: git push origin feature-branch")
		fmt.Println("This command pushes the changes in 'feature-branch' to the remote repository.")
	case "log":
		fmt.Println("The 'git log' command displays the commit history of the repository.")
		fmt.Println("Example: git log")
		fmt.Println("This command shows a log of commits, including commit messages and authors.")
	case "clone":
		fmt.Println("The 'git clone' command clones a repository into a new directory.")
		fmt.Println("Example: git clone https://github.com/example/repo.git")
		fmt.Println("This command creates a copy of the specified repository in a new directory.")
	case "remote":
		fmt.Println("The 'git remote' command manages remote repositories.")
		fmt.Println("Example: git remote add origin https://github.com/example/repo.git")
		fmt.Println("This command adds a remote named 'origin' for the repository.")
	case "fetch":
		fmt.Println("The 'git fetch' command fetches changes from a remote repository without merging.")
		fmt.Println("Example: git fetch origin")
		fmt.Println("This command retrieves changes from the 'origin' remote repository.")
	case "reset":
		fmt.Println("The 'git reset' command unstages changes or resets the repository to a previous state.")
		fmt.Println("Example: git reset HEAD file.txt")
		fmt.Println("This command unstages changes made to 'file.txt'.")
	case "tag":
		fmt.Println("The 'git tag' command creates and manages tags for releases in the repository.")
		fmt.Println("Example: git tag -a v1.0 -m 'Version 1.0'")
		fmt.Println("This command creates an annotated tag 'v1.0' with a message.")
	default:
		fmt.Printf("Explanation for '%s' is not available. Try another Git command.\n", command)
	}
}

func explainAdvancedGitConcepts(command string) {
	switch command {
	case "rebase":
		fmt.Println("- Rebase is one of two Git utilities designed to integrate changes from one branch onto another. Rebasing is the process of combining or moving a sequence of commits on top of a new base commit. Git rebase is the linear process of merging.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git rebase:
git rebase <base>                   Performs the standard rebase
git rebase <base>                   Performs the standard rebase
git rebase – interactive <base>     Performs the interactive rebase
git rebase -- d                     The commit gets discarded from the final combined commit block during playback.
git rebase -- p                     This leaves the commit alone, not modifying the content or message, and keeping it as an individual commit in the branches’ history.
git rebase -- x                     This executes a command line shell script for each marked commit during playback.
git status                          Checks the rebase status.
git rebase -- continue              Continue with the changes that you made.
git rebase --skip                   skips the changes`)
		fmt.Print("\n")
		fmt.Println(`Example:

To rebase development to master the command is like the following
$ git rebase master development`)
	case "cherry-pick":
		fmt.Println("- Cherry-pick is a Git feature that allows you to apply a single commit or a range of commits from one branch to another. It's useful when you want to pick specific changes without merging the entire branch.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git cherry-pick:
git cherry-pick <commit>    Apply the changes introduced by the specified commit
git cherry-pick -x          Create a new commit with the same authorship information as the original commit
git cherry-pick -e          Edit the commit message before applying
git cherry-pick -n          Apply changes but don't commit, allowing further modifications
git cherry-pick -m <parent> Specify the mainline parent for the cherry-pick operation`)
		fmt.Print("\n")
		fmt.Println(`Example:

To apply changes from a specific commit:
$ git cherry-pick abc123`)
	case "submodule":
		fmt.Println("- Git Submodules are a way to include external repositories within a Git repository. They allow you to keep a reference to an external repository at a specific snapshot, making it easy to update the submodule to a newer version later.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git submodules:
git submodule add <repository> [<path>]    Add a new submodule
git submodule init                        Initialize submodules for the first time after a clone
git submodule update                      Update the submodules to the latest commit`)
		fmt.Print("\n")
		fmt.Println(`Example:

To add a submodule:
$ git submodule add https://github.com/example/repo.git path/to/submodule`)
	case "stash":
		fmt.Println("- Git stash is a command used to save changes that haven't been committed to a temporary area so that you can switch branches or perform other operations without committing incomplete changes.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git stash:
git stash                Save your changes to a new stash
git stash list           List all stashes
git stash apply          Apply the changes from the latest stash
git stash pop            Apply and remove the latest stash
git stash drop <stash>   Discard a stash`)
		fmt.Print("\n")
		fmt.Println(`Example:

To save changes to a stash:
$ git stash`)
	case "reflog":
		fmt.Println("- Git reflog, short for reference logs, records when the tips of branches and other references were updated in the local repository. It provides a way to review and recover previous states of the repository.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git reflog:
git reflog               Show a log of changes, including those that may not be visible in regular history
git reflog show <branch> Display the reflog for a specific branch`)
		fmt.Print("\n")
		fmt.Println(`Example:

To view the reflog for the current branch:
$ git reflog`)
	case "hooks":
		fmt.Println("- Git hooks are scripts that run automatically before or after certain Git commands. They allow you to customize and automate processes in your Git workflow.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git hooks:
No specific Git commands, but hooks are executed automatically based on predefined events.`)
		fmt.Print("\n")
		fmt.Println(`Example:

Implementing a pre-commit hook to check code formatting:
1. Create a script named pre-commit in the .git/hooks directory.
2. Add code to check code formatting.
3. Make the script executable: chmod +x .git/hooks/pre-commit`)
	case "gitflow":
		fmt.Println("- Gitflow is a branching model for Git that defines a standard set of branches and a consistent workflow. It provides a higher-level abstraction of the Git commands to support a successful branching strategy.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Gitflow:
git flow init                Initialize a new repository for Gitflow
git flow feature start      Start a new feature branch
git flow feature finish     Finish a feature branch and merge it into the develop branch
git flow release start      Start a new release branch
git flow release finish     Finish a release branch, merge it into master, and tag the release`)
		fmt.Print("\n")
		fmt.Println(`Example:

Using git-flow to start a new feature:
$ git flow feature start new-feature`)
	case "revert":
		fmt.Println("- Git revert is used to create a new commit that undoes the changes made by a previous commit. It's a safer way to undo changes compared to Git reset, as it doesn't modify existing commits.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git revert:
git revert <commit>    Create a new commit that undoes changes introduced by the specified commit`)
		fmt.Print("\n")
		fmt.Println(`Example:

To revert the changes made by a specific commit:
$ git revert abc123`)
	case "filter-branch":
		fmt.Println("- Git filter-branch is a complex and powerful command used for rewriting branch history. It allows you to filter and modify the branch's commit history based on specific criteria.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git filter-branch:
git filter-branch <options>    Rewrite the branch history based on specified options`)
		fmt.Print("\n")
		fmt.Println(`Example:

To remove a file from the entire commit history:
$ git filter-branch --tree-filter 'rm -f file.txt' -- --all`)
	case "bisect":
		fmt.Println("- Git bisect is a binary search tool used to find a specific commit that introduced a bug or regression. It helps narrow down the range of commits where the issue was introduced.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Git bisect:
git bisect start                  Start the bisecting process
git bisect good <commit>          Mark a commit as good (bug-free)
git bisect bad <commit>           Mark a commit as bad (buggy)
git bisect reset                  Finish the bisecting process`)
		fmt.Print("\n")
		fmt.Println(`Example:

To start a bisect session:
$ git bisect start`)
	default:
		fmt.Printf("Explanation for '%s' is not available. Try another advanced Git concept.\n", command)
	}
}

func usageTemplate() string {
	return `Usage:
  explain git [flags]

Flags:
  -a, --advanced string   Explain advanced Git concepts
  -c, --command string    Specify a Git command to explain
  -h, --help              help for git

Examples:
	explain git --command branch,
	explain git --c reset
	explain git --advanced rebase
	explain git -a cherry-pick

Available Commands:
  init, add, commit, status, branch, merge, pull, push, log, clone, remote, fetch, reset, tag

Available Advanced Topics:
  rebase, cherry-pick, submodule, stash, reflog, hooks, gitflow, revert, filter-branch, bisect
`
}
