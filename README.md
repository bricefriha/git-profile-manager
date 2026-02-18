# GPM - Git Profile Manager
GPM is a CLI tool to manage git profiles.
The idea of buidling this tool came about after I needed to find a way to switch quickly between my github work account and my personal one.

## Create a profile
1. Create a file `.gitconfig-{name of the profile}`. example: if you want to create a profile called `personal` it would be `.gitconfig-personal`
2. Fill up this file with the gitconfig that is relevant to the profile

## Switch between profiles
Straight forwardly:
``` bash
gpm switchto {profile name}
```
