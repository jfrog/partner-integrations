#!/bin/sh -e

git checkout master
git pull
echo
echo "The latest tag:"
git describe --abbrev=0 --tags
echo
read -p "Please provide new tag: " tag
git tag -a -m "Release $tag" "$tag"
git push origin refs/tags/"$tag"
