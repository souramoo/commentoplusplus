#!/bin/bash

echo Setting config
git config --global user.email "actions@circleci.com"
git config --global user.name "GitHub"
git config pull.rebase true

echo Replacing origin
git remote rm origin
git remote add origin https://gitlab.com/commento/commento.git

echo Checking out master
git checkout master

echo Branching
git branch -m master-holder

echo Fetching
git fetch

echo Checking out new master
git checkout master

echo Pulling new master from new origin
git pull origin master

echo Merging.....
git merge -s recursive master-holder --allow-unrelated-histories --no-commit
git commit -m "[skip ci] Updates from GitLab"

echo Removing new origin back to github
git remote rm origin
git remote add origin https://github.com/souramoo/commento-heroku.git

echo Pushing back changes...
git push --set-upstream origin master
