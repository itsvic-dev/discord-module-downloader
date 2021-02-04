# Discord Module Downloader
Download the modules for the official Discord client with ease.

## Usage
### Downloading
To download the modules, run:
```bash
./download.py
```

On Windows, you'll need to run:
```cmd
python .\download.py
```
### Extracting
To extract the modules, run:
```bash
./extract.sh
```

Some modules (like `discord-desktop-core`) will have `.asar` files. To unpack them, simply install asar with `sudo npm install -g asar` and run `asar extract path/to/file.asar path/to/output`.

**Note**: You do need `p7zip` (on Debian and derivatives: `p7zip-full`) for the script to work.

**Windows users**: I don't have a script to automatically extract them (I'm an Arch Linux user), but you can use 7-Zip's "Extract to *" context menu option.

## Technical stuff
Here, I'll delve into how the Discord updater checks for host updates and module updates.

**Note**: I did this a few days ago, so some specifics and variable names may not be correct.

### Constants
The updater defines a base path, `https://discord.com/api`. For some reason, API the version is ommitted. No idea why.

It will also define the release (`stable`, `ptb` or `canary`) as a variable.

It also defines the platform (`win`, `linux` or `osx`) as a global query string parameter, named `platform`.

### Host updates
The latest host version is derived from `${BASE_PATH}/updates/${release}`, with the default, global query string.

If the latest host version is different from the current one, it downloads the latest one, then restarts the updater.

The updater will update if one of the following conditions are true:
- The parameter `ALWAYS_ALLOW_UPDATES` is set to true in settings.
- The host version is `0.0.0` **OR** the updater is a debug build.

This also affects module updates, I believe.

### Module updates
The updater will append a new variable to the global query string, `_`, which serves... what purpose does it serve, actually? It's fine if it's set to `0`, Discord won't mind.

It creates a new variable, `MODULE_UPDATE_PATH`, with its value being `${BASE_PATH}/modules/${release}`.

The updater then polls `${MODULE_UPDATE_PATH}/versions.json` for the latest modules, with the query string parameter being:
```js
{"host_version": currentHostVersion} + globalQueryString
```

It will check if any module updates require an update, and if it is the case, it will download the respective modules from this path: `${MODULE_UPDATE_PATH}/${moduleName}/${moduleVersion}` into `~/.config/discord/modules/pending/${moduleName}-${moduleVersion}.zip` (Linux) with the same query string as before. The `_` query string variable can be ommited.

The updater will then unpack the ZIP files of the modules into directories conveniently named after the module names, then remove the ZIP files.

The updater then presumably initializes the `discord-desktop-core` module.
