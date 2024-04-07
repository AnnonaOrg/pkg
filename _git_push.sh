#!/bin/bash

VERSION=0.0.11
APPNAME=pkg

git add .
git commit -m "debug"


#git remote add github git@github.com:AnnonaOrg/pkg.git
#git branch -M main
git push -u github main

git tag "v${VERSION}"
git push --tags -u github main

