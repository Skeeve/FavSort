#!/bin/bash -e
(
    cd "$( dirname "$0" )"
    me="$( basename "$0" )"
    dir="$PWD"
    # Get the owner of this directory as parameters to tar
    setOwner="$( stat -c "--no-same-owner --owner=%U --group=%G" "$dir" )"
    # Check where the Archive in this file starts
    archiveAt="$( grep -nE '^# ARCHIVE$' "$me" | head -n 1 | cut -d: -f1 )"
    # It's one line after
    ((++archiveAt))
    # Extract the archive
    tail +$archiveAt "$me" | base64 -d | tar $setOwner -xzf -

    # Initialize
    "$dir/FavSort/append"

    # Show Version information
    "$dir/FavSort/version"

    # Remove Installer
    rm "$dir/$me"

) > "$0.log" 2>&1
# Remove the log
[[ -s "$0.log" ]] || rm "$0.log"
exit 0
# ARCHIVE
