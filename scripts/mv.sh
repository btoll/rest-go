#!/bin/bash

# http://stackoverflow.com/a/14127035/125349
if [ "$PWD" != $(git rev-parse --show-toplevel) ]; then
    echo "$(tput setaf 3)[WARNING]$(tput sgr0) You need to run this command from the toplevel of the working tree."
    exit 1
fi

dest="$GOPATH/src/github.com/dgaedcke/nmg_admin_service"

#(
#    cd app;
#
#    for f in {contexts,controllers,media_types,user_types}.go; do
#        sed 's/\(package\) app/\1 admin/' $f > "$GOPATH/src/github.com/dgaedcke/nmg_shared/app/admin/$f"
#    done;
#)

# TODO: Brittle!
# Move controllers from top-level into subdirectory and change package name.
for f in {event,game,sport,team,team_opening_config}.go
do
    # Only change text within the first 10 lines.
    sed '1,10s/\(package\) main/\1 controllers/' $f > "$GOPATH/src/github.com/dgaedcke/nmg_admin_service/controllers/$f"
#    sed -i '1,10s/nmg_admin_service/nmg_shared/' "$GOPATH/src/github.com/dgaedcke/nmg_admin_service/controllers/$f"
#    sed -i '1,10s/\(nmg_shared\/app\)/\1\/admin/' "$GOPATH/src/github.com/dgaedcke/nmg_admin_service/controllers/$f"

    # Only change text starting at the 11th line.
#    sed -i '11,$s/\bapp\b/admin/' "$GOPATH/src/github.com/dgaedcke/nmg_admin_service/controllers/$f"
done

cp main.go "$dest/main/"
cp app/* "$dest/app/"
cp client/* "$dest/client/"
cp design/* "$dest/design/"
cp models/* "$dest/models/"
cp swagger/* "$dest/swagger/"
cp -r tool/* "$dest/tool/"

(
    cd $dest;

    for f in $(ag -l btoll)
    do
    sed -i 's/btoll\/rest-go/dgaedcke\/nmg_admin_service/' $f
    done;
)

