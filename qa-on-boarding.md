Using GitHub to Manage the Functional Test Scenarios for QA
===========================================================

Meeting Questions
-----------------

Tooling?  Sublime Edit, IntelliJ?
Committing? GitHub Mac?  IntelliJ?
Push Access? Only to fork. 

We want to try to keep the local tooling to a minumum.

We will need to automate the upstream synchronization to keep the QA fork up to date.

Fast build-per-branch?? (build diff only)

Fork Bloomhealth + Bhbo to sub repo.

Signing Up
----------

If you don not have a GitHub account yet, you can sign up for one 
[here](https://github.com/signup/free)

After you have signed up, send an email to helpdesk@bloomhealthco.com with your 
GitHub username, asking them to add your account to the QA team under the 
bloomhealth organization.

The Software
------------

You will be working with the bloom health functional testing repositories two different ways

* The GitHub Web interface

https://github.com/bloomhealth

* The GitHub Mac client

http://mac.github.com/

The Repositories
----------------

The files you will be working with are called `repositories`.  

There are currently three repositories used to store the files used in 
functional testing

* https://github.com/bloomhealth/test_geb_page_objects
* https://github.com/bloomhealth/webapp_bloomhealth
* https://github.com/bloomhealth/webapp_bhbo

You will not have write access to these repositories.  Instead, you will create your 
own copies to work on.  These copies are called `forks`.

Forking the Repositories
------------------------

To create your own `fork` of a repository, simply click the `Fork` button on the 
top of the bloomhealth repository page that you want the copy of.

If you have already `forked` the repository, you will see a message stating something like:

    Already forked to your_github_username! Go to the fork

Once you have forked the repository, you can create a local copy of your fork on 
your laptop by clicking the `Clone in Mac` button.  MAKE SURE TO CLONE FROM YOUR OWN FORK.

* Note: the first time you click this, you will probably see a dialog stating *External Protocol Request*, and asking what you want to do with the link you just clicked.  Just check the `Remember my choice for all links of this type.` checkbox, then click `Launch Application`.

*GitHub Mac* will then Ask you two things

* Clone as : This is the name of the local folder you want to create on your laptop to hold the repository.
* Where : This is the folder you wish to put the repository in.

For `Clone as`, keep this the same as the name as the repository for simplicity.
For `Where`, you can just put it in your `Documents` folder.

* Note: the first time you clone a repo locally, you will be asked if *GitHub 
for Mac* can use your confidential information stored in your keychain.  
Choose `Always Allow` to encrypt this using your login password.

Synchronizing your Repositories
-------------------------------

The GitHub Mac client does not automatically synchronize your personal GitHub 
`fork` of the official matching `upstream` bloomhealth repository.  To remedy
this, you will be given a command line script that you will be able to run to
update synchronize all of your repositories (the upstream official bloomhealth 
repo, your own GitHub fork of the repo called 'origin', and your local clone
of your GitHub fork that resides on your MacBook.

 * Details to be determined... stay tuned.

Getting your tests into the official Bloom codebase
---------------------------------------------------

Because you cannot write directly to the official bloom repositories, you will
need to work off of your own fork.  When you are done working off of your fork
for a certain suite of tests, you will need to get these changes into the 
official bloom repositories.  This process is called a `pull request`.  Each
pull request you create is tied to a modified version of your local 
repository clone.  This is called a 'branch'.  Each time you work on a new
set of functional tests, you can create a new 'branch' that is based on the 
main 'develop' branch, give it a useful name, write your tests, `push`, or 
upload your branch to your own GitHub fork, and once you feel it is ready,
create a `pull request` from your 'branch' to the official bloom repository
'develop' branch.

Writing a new Set of Test Cases
-------------------------------

These are the steps you will follow to create or update a set of functional
tests.

* Make sure your local 'develop' branch is up to date.
  This process is still to be determined.
* Create a 'my-new-fun-feature-blah' branch based on 'develop'
  This can be done by clicking the '+' next to the 'develop' branch
* Add tests to an existing test suite, or create a new test suite based on 
  the template:
  https://github.com/aaronzirbes/webapp_bloomhealth/tree/gold-standard/bloomhealth/test/functional
* Push, or 'Publish' your 'my-new-fun-feature-blah' branch back up to your personal 
  GitHub fork, often referred to as your 'origin' repository.
* When you feel it is ready, you will create a pull request by clicking
  on the `Pull Request` button on the GitHub page for your personal fork.
  https://github.com/your_github_username/repository_name/
* If someone comments on your pull request, and you need to make changes,
  then make changes.  Re-sync your 'origin', and the pull request
  will be automatically updated.
* Once your pull request is accepted, you can close your local branch.

Editor
------

While Microsoft Word is sufficient for composing a letter to your pen-pal
in Eastern Europe.  Under no circumstances, will (or should you ever) use
it to write code.  Functional tests are code.  There is no way around it,
you will be writing code.

All the pretty indentation, and macical auto-completion of Word and the 
family of products if comes from is evil in the mind of a coder.

Because of this you will need a descent text editor.  We highly suggest a 
product named [ Sublime Text ](http://www.sublimetext.com/).  Af first 
glimpse it seems too simple.  There are no fonts, point sizes, itallics
or underlineing.  But once you use it you will learn of the magic it contains.

A good text editor will do the following things for you.

* Color code key words
* find matching brackets, quotes, parentheses, etc...
* help you indent things cleanly
* respect spaces and tabs
* help you auto-complete key words, variables, functions, etc...

If you associate .groovy files with this editor, it will help you quickly open them.  This can be done by:
* Right click a file ending in .groovy
* Choose 'Get Info'
* At the bottom, under 'Open with:', choose 'Sublime Text 2' (or whatever you like), then click the 'Change All...' button.


Writing the test case
---------------------

* If you are writing a new test case, please start from the following template:
  https://github.com/aaronzirbes/webapp_bloomhealth/tree/gold-standard/bloomhealth/test/functional
** We will try to update this as often as we can to ensure that you have to make the least amount of changes in your pull request.
* Copy the template test class to the folder under 'bloomhealth/test/functional/' 
  in the application you feel is most applicable to your new test classes.
* Rename your test class to something appropriate to what you are testing.  For example
  'SponsorQualifyingEventCreationFunctionalSpec.groovy'
* Edit the first line, starting with 'package', to match the folder name you put your test in.
* Edit the 'class' name of your test to match the file name.
* Start creating test cases following the instruction inside the test template file.
* Save your test class file.
* Go into GitHub Mac and review your uncommitted changes for commit.
* Fill in the commit summary (keep it less than 50 letters)
* Fill in the extended description for what you want to commit.
* Select the files you want to commit
* Click `commit`

When you are done with all 'commmits', you can publish these back up to GitHub by clicking the `Publish Branch` button in the upper right hand corner.

If this fails, you might have cloned directly from the bloomhealth repo.  To get this fixed, you may have to ask a dev to help.
Just tell them:

    I need the origin on my local git repo changed from the https bloomhealth repo to my own https repo fork.

Once all your commits are done, you can create a pull request as outlined above.

Questions?
----------

If you have questions, you can ask a dev as they are rather familiar with the process and then when you 
know the answer, you can add it to this Wiki page to help others out in the future.

Good luck!

