# Readme

`FavSort` allows you to easily sort your favorites on your Anbernic RG35XX H alphabetically.

I'm pretty confident it will run as well on the "Plus" and "SP", but couldn't test that.

FavSort will run at least with Stock OS from 1.1.0 to 1.1.4.
It's also compatible with cbepx-me mods for these OSs.

I'd be happy to hear from users of Tom's mods.

**NOTE** Some games, especially Arcade ones, on the Anbernic will display another name than the filename.
Sorting in this case is done by the english name of the game even though you might see the name displayed in another language.
The full (english) names are taken from `/mnt/vendor/bin/arcade-plus.csv`.
The location of that file is defined in [favsort.yaml](favsort.yaml).

## Recommendation

**Before** installing FavSort, it is recommended, that you check your favorites list.
This is done by simply scrolling through the list.

In case there is a wrong entry, the device will reboot.
Figure out, which entry is the faulty one.
After installation, the entry can be removed as described below.

**Precaution:** You may want to back up your favorites.
Copy `/mnt/data/misc/.favorites` using `File Manager`.

## Installation

You already unzipped `FavSort.zip`.
Now you have this `Readme.md` file, a file `build-id.txt`, and the file `Install FavSort.sh`.

The `build-id.txt` File contains the "build id" of the executable.
When installing on the Anbernic device, the Version and build id will be displayed for a few seconds.
You can compare the displayed build id with the one in the file.
If the time is too short, you can always recall this screen later.

Copy the `Install Favsort.sh` File to the `Roms/PORTS` directory of the sdcard (TF1 or TF2).

- If you have Roms on both cards, choose either of them, **not both**.
- If you have Roms on only one card, use that one for your convenience.

Start `Install FavSort` from `RA Game` >> `PORTS`.

It will install itself under the name `++ sort Favorites ++` into the same directory where the installer is located (i.e. the `PORTS` directory).
It then appends itself as a favorite.
As a last step, the Installer will get removed.

### Installed helpers

Under `Roms/PORTS/FavSort` you will find a few helpers which you can invoke from `File Manager`:

- version \
  This will display the version information for 10 seconds
- repair \
  This will fix a manually edited `.favorites` file.
  It is required after making changes to `.favorites` using, for example, `File Manager`'s text editor.
  See also [Repair broken Favorites / modify Favorites](#repair)
- append \
  This will just append `++ sort Favorites ++` to your favorites list without sorting it.
- remove \
  This will just remove `++ sort Favorites ++` from your favorites list without sorting it.
- uninstall/uninstall \
  **Note:** This utility is located in the `uninstall` directory to avoid running it by mistake.
  It will completely remove FavSort from your device.

## Usage

When you sort your favorites, the last entry in your favorites list will always be

`++ sort Favorites ++`

When you later add more favorites, they will appear **after** `++ sort Favorites ++`.

Simply call `++ sort Favorites ++` to sort the newly created entries and get `++ sort Favorites ++` back to the end again.

Favorites will always be sorted in ascending order unless they already were in ascending order.
In that case they will be ordered descending.


## Repair broken Favorites / modify Favorites <a id="repair"></a>

Some Favorites lists make the device reboot.

I had one case, where I had a favorites entry containing `…:APPS:…`.
This must have been a leftover from an earlier installation, but each time that entry would scroll into view, the device rebooted.

With FavSort there is a way to repair such lists:

1. Start `File Manager`
2. Move to `/mnt/data/misc`
3. Edit `.favorites`
4. Remove the line which caused the reboot
5. Leave the editor but **not** `File Manager`
6. Move to the `Rom/PORTS/FavSort` folder and start `repair`

This should fix your favorites list.

## View Installation Information (Version) again

1. Start `File Manager`
2. Move to the `Rom/PORTS/FavSort` folder and start `version`

## More Usage Tricks

It's possible to call `++ sort Favorites ++` with commandline parameters, if you are able to `ssh` to your device.

These are the flags you can use:

- --repair (-r) \
  Don't sort, but just repair.
  Fixes linenumbers and checksum.
- --append \
  Don't sort, but just repair and append the favsort entry.
- --remove \
  Don't sort, but just repair and remove the favsort entry.
- --tsv (-t) \
  Just print out the sorted Game names and all its information.
- --short (-s) \
  Just print out the sorted and shortened Game names, as used for sorting and all its information.
- --version \
  Show the version number and build id
- --asc (-a) \
  Always sort ascending.
- --desc (-d) \
  Always sort descending.
- --help (-h) \
  Show help

As you saw above, calling `++ sort Favorites ++` will toggle between ascening and descending sorts.
If you prefer to always have one of those, edit the shell script `++ sort Favorites ++.sh` to pass `--asc` or `--desc` to the call:

```shell
    *)
        "$exename" --asc
        ;;
```

The `--repair` option does not sort and also does **not** add `++ sort Favorites ++` to the end.
So this is a way how you can manually sort your favorites.
Sort with whatever editor you prefer and call the programm with `--repair` or
use the `repair` command from `PORTS/FavSort`.

## Uninstall

To uninstall simply call, using `File Manager` the `uninstall` program in
`PORTS/FavSort/uninstall`.
This will:

1. Remove `.. sort Favorites ++` from the favorites.
2. Delete all files which were installed.

## Credits

Many thanks to the reddit user [AnonymousTokenus](https://www.reddit.com/user/AnonymousTokenus/), for

- giving me the idea to this tools
- beta-testing
- helping me figuring out the meaning of the fields in the favorites file
- and for challenging my programming skills

Also many thanks to discord user @ rymsar for finding a critical bug just in time before I wanted to publish FavSort.
