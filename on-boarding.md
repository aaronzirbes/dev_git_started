Git Topics for Developers
=========================

Overview
--------

Why are we switching to Git?
Why are we forking the main repository?

 * Ensure all commits to bloom applications happen through pull requests.
 * Lets you squash your commit histories before pull requests get merged in

Squashing Commits

 * Doing something
 * Fixing Test
 * DE1234 - This is the acuall commit
 * 'rebase' + fix typo

Hg cannot handle more than ~1000 branches

Pull requests

Build-per-branch (Coming Soon)

Software
--------

* Git

    brew install git

or http://git-scm.com/download/mac

* Git Flow

    https://github.com/petervanderdoes/gitflow/wiki/Installing-on-Mac-OS-X

    brew install gnu-getopt

    echo 'alias getopt="$(brew --prefix gnu-getopt)/bin/getopt"' > ~/.gitflow_export

* GitHub Mac

    http://mac.github.com/

Setup
-----

https://help.github.com/articles/set-up-git

* Set global git name
* Set global git email
* Setup SSH keys 

To use credential-osxkeychain (HTTPS)

    git config --global credential.helper osxkeychain

Pretty

    git config --global color.ui true

Getting the Repositories
------------------------

https://github.com/bloomhealth

Quick and Easy

    git clone git@github.com:aaronzirbes/$repo.git
    git remote add upstream git@github.com:bloomhealth/$repo.git
    git fetch --all 
    git checkout -t origin/develop

Extended version

    # Create the folder to hold the repository
    mkdir $repo
    # Go into the folder
    cd $repo
    # Initialize the repo
    git init
    # Add the remote origin alias
    git remote add origin https://github.com/$username/$repo.git
    # Add the remote upstream alias
    git remote add upstream https://github.com/bloomhealth/$repo.git
    # Fetch all remote reference information (tags, branches, refs)
    git fetch --all
    # List all remote branches in origin
    all_branches=`git branch -r |grep -v ' -> ' |grep -E ' origin/' |sed -e 's#  origin/##'`
    # For each origin branch, create a local tracking branch
    for branch in ${all_branches}; do 
        git branch --set-upstream $branch origin/$branch
    done
    # checkout the develop repository
    git checkout develop


Keeping them up to date
-----------------------

    # What replaces 'hg incoming' and 'hg outgoing'?
    git fetch $remote
    git log ${somebranch}..${anotherbranch}
    git log develop..upstream/develop
    git log develop..origin/develop
    git log feature/my-feature..origin/develop

    # Update the develop branch (Get ready to start on a story)
    git update develop
    git pull upstream develop
        ~= `hg pull && hg merge`

    # Update a story branch (Pull in changes from develop into your story branch)
    git flow update feature/my-new-feature

    git pull -u origin feature/my-new-feature
    git pull upstream develop

    git fetch upstream develop
    git rebase -i upstream/develop

    git pull --rebase upstream develop

    # Update the master branch (Get ready to work on a hot-fix / patch)
    git update master
    git pull upstream master


    # Working on a branch with a team
    git remote add charliek https://github.com/bloomhealth/$repo.git
    git fetch --all
    git checkout -t charliek/feature/new-feature-charlie-started

    # Pull in changes from charlie
    git fetch charliek feature/new-feature-charlie-started
    git rebase -i charliek/feature/new-feature-charlie-started

    # If you want to push your local changes up to your GitHub fork:
    git push
    # ...or explicitly...
    git push origin $branch

Using Git
---------

The 'index'

    # add specific files
    git add modified-file.txt
    git add another-file.txt

    # add everything
    git add .

Committing 
----------

[ Git Style Commit messages ]( http://git-scm.com/book/en/Distributed-Git-Contributing-to-a-Project )

    git log --no-merges

* Example

    Short (50 chars or less) summary of changes

    More detailed explanatory text, if necessary.  Wrap it to about 72
    characters or so.  In some contexts, the first line is treated as the
    subject of an email and the rest of the text as the body.  The blank
    line separating the summary from the body is critical (unless you omit
    the body entirely); tools like rebase can get confused if you run the
    two together.

    Further paragraphs come after blank lines.

     - Bullet points are okay, too

     - Typically a hyphen or asterisk is used for the bullet, preceded by a
       single space, with blank lines in between, but conventions vary here

Commands

    # commit all pending files
    git commit
    git commit -m 'message' modified-file.txt another-file.txt
    git commit -a -m 'message'
    git commit --interactive -m 'message'


Using Git Flow
--------------

    # Initialize your repo
    git flow init -d

    git flow feature start s1234-my-story

When you are ready to create your pull request, or need to share with someone else

    git flow feature publish s1234-my-story

If you are working with someone else, and need to pull in their changes

    git flow feature pull s1234-my-story

After your pull request has been accepted

    git flow feature finish s1234-my-story

    # This should put you back on the 'develop' branch.  If not, run `git checkout develop`
    git pull upstream develop

The following will do nothing to your local branch

    git fetch upstream develop

And if you want to, you can update your GitHub fork of develop. (optional)

    git push

What are the changes between your current branch and the upstream develop branch?

    git diff --stat -r upstream/master

What is the diff between your current branch and the upstream develop branch?

    git diff -r upstream/master

Pull Request
------------

    https://github.com/$username/$repo/pull/new/feature/$feature_name

The Book
--------

http://git-scm.com/book

Other toys
----------

[ Hub ]( https://github.com/defunkt/hub )

[ SCM Breeze ]( https://github.com/ndbroadbent/scm_breeze )

