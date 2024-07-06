#!/bin/bash -e
cmd="$( basename "$0" ".sh" )"

exefile="$( readlink -f "$0" )"
dir="$( dirname "$exefile" )"
cd "$dir"
exename=./"$( basename "$exefile" ".sh" )"

case "$cmd" in
    version)
        subtitle="./$$.vtt"
        version="$( "$exename" --version )"
        printf "WEBVTT\n\n00:00:00.000 --> 99:59:59.999\n%s\n" "$version" > "$subtitle"

        pkill -f mpv || true
        mpv --image-display-duration=10 --really-quiet --video-zoom -1 --sub-file="$subtitle" --video-zoom=-1 "../Imgs/++ sort Favorites ++.png"
        pkill -f mpv || true
        rm "$subtitle"
        exit 0
        ;;
    install)
        "$exename" --install
        ;;
    repair)
        "$exename" --repair
        ;;
    uninstall)
        cleaned="$( grep -Fv "$exename.sh:" '/mnt/data/misc/.favorite' )"
        echo "$cleaned" > '/mnt/data/misc/.favorite'
        "$exename" --repair
        rm -rf "../Imgs/$exename.png" "../$exename.sh" "$dir"
        ;;
    *)
        "$exename"
        ;;
esac

true && exit 0 ; : <<'--'

List of Arcade directories. Extend as you wish

#ARCADE /mnt/vendor/bin/arcade-plus.csv
ATOMISWAVE
CPS1
CPS2
CPS3
FBNEO
HBMAME
MAME
NAOMI
NEOGEO
PGM2
VARCADE
#END

--
