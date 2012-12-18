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

If you don not have a GitHub account yet, you can sign up for one [here](https://github.com/signup/free)

After you have signed up, send an email to helpdesk@bloomhealthco.com with your GitHub username, asking them to
add your account to the QA team under the bloomhealth organization.

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

There are currently three repositories used to store the files used in functional testing

* https://github.com/bloomhealth/test_geb_page_objects
* https://github.com/bloomhealth/webapp_bloomhealth
* https://github.com/bloomhealth/webapp_bhbo

You will not have write access to these repositories.  Instead, you will create your own copies to work on.  These copies are called `forks`.

Forking the Repositories
------------------------

To create your own `fork` of a repository, simply click the `Fork` button on the top of the bloomhealth repository page that you want the copy of.

If you have already `forked` the repository, you will see a message stating something like:

    Already forked to your_github_username! Go to the fork

Once you have forked the repository, you can create a local copy of your fork on your laptop by clicking the `Clone in Mac` button.

* Note: the first time you click this, you will probably see a dialog stating *External Protocol Request*, and asking what you want to do with the link you just clicked.  Just check the `Remember my choice for all links of this type.` checkbox, then click `Launch Application`.

*GitHub Mac* will then Ask you two things

* Clone as : This is the name of the local folder you want to create on your laptop to hold the repository.
* Where : This is the folder you wish to put the repository in.

For `Clone as`, keep this the same as the name as the repository for simplicity.
For `Where`, you can just put it in your `Documents` folder.

* Note: the first time you clone a repo locally, you will be asked if *GitHub for Mac* can use your confidential information stored in your keychain.  Choose `Always Allow` to encrypt this using your login password.




