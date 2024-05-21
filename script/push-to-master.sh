#!/usr/bin/env bash

git checkout master

# merge the changes from the dev branch, no interactive mode
git merge dev --no-edit

# push the changes to the master branch
git push origin master

# go back to the dev branch
git checkout dev