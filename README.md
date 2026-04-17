# gscene + Ebitengine game template

How to use:

```
git clone https://github.com/quasilyte/gscene-game-template.git mygame
cd mygame
rm -rf .git
```

After that, you may want to add your own remote and push there:

```
git init
git remote add origin https://github.com/username/ld59game.git

# Then add/push to that remote and you're ready to go
```

This template is intended to be hard-editted - not everything here will fit an arbitrary game.

See [go.mod](go.mod) to study the libraries this template is using by default. You may not need all of it for your game, but it is easier to run `go tidy` than looking for a suitable library during the game jam.

It is built with game jams and quick prototyping in mind - you should not use it to build a full release game.

What is provided:

* A project layout that should be easy and extensible enough for prototypes
* Some basic stub-like assets for UI
* Resource loading system
* A bunch of useful libraries that work well together

To help you with code navigation, here are a few things to search for (`ctrl+shift+f`):

* `<<edit>>` - the first-priority edits, for every game

There are two ways to include game contents (sounds, images): embedded and dynamic. For embedded storage, see [image_resources.go](src/assets/image_resources.go), etc.
For dynamically loaded resources, put the files under `game_data` folder. You may need to adjust the resource loading, but the most basic stuff is already there.
