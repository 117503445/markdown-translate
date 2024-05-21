#!/usr/bin/env bash

git checkout master

# merge the changes from the dev branch
git merge dev

# push the changes to the master branch
git push origin master

# go back to the dev branch
git checkout dev