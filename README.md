# Deprecation warning
This program is no longer maintained and it's also bugged

# go-i3altlayout
A port of [i3altlayout](https://github.com/deadc0de6/i3altlayout) in Go using [i3ipc-go](https://github.com/mdirkse/i3ipc-go).

It should be entirely compatible with the original, just use the install script (remember to have go installed).

---

# What does it do?
`i3altlayout` helps you handle more efficiently your screen real estate in [i3wm](https://i3wm.org/)
by auto-splitting windows on their longest side.

If you open 4 windows, your workspace would look

```
             like this                               instead of this

+-----------------------------------+        +---------------------------------+
| +--------------+ +--------------+ |        | +-----+ +-----+ +-----+ +-----+ |
| |              | |              | |        | |     | |     | |     | |     | |
| |              | |      2       | |        | |     | |     | |     | |     | |
| |              | |              | |        | |     | |     | |     | |     | |
| |      1       | +--------------+ |        | |  1  | |  2  | |  3  | |  4  | |
| |              | +--------------+ |        | |     | |     | |     | |     | |
| |              | |      ||      | |        | |     | |     | |     | |     | |
| |              | |   3  ||  4   | |        | |     | |     | |     | |     | |
| +--------------+ +--------------+ |        | +-----+ +-----+ +-----+ +-----+ |
+-----------------------------------+        +---------------------------------+
```

---

# Installation

User the install.sh script
```bash
$ ./install.sh
```
if you only wish to build it without installing it, use the `build.sh`

---

# Usage

Start i3altlayout directly with [i3](https://i3wm.org/) by adding it to your config file
(usually under `~/.config/i3/config`):
```
exec --no-startup-id "killall i3altlayout"
exec --no-startup-id "i3altlayout"
```
(don't forget to add the killall, if you don't do it you will keep spawning a new process on each restart of i3)
