# Empty Directory Clean

A tiny `golang` script to remove folder that is empty. Support detection of nested directories.

## Quick Start

Given a folder tree:
```
folder1
├─ folderA
│  ├─ {...some file}
├─ folderB
folder2
├─ folderC
│  ├─ {...some file}
├─ folderD
│  ├─ .DS_Store
folder3
{executable of this repo}
```

Running executable will delete
- `folder1/folderB`, as no file inside
- `folder2/folderD` as one `.DS_Store` inside (unnecessary file)
- `folder3` as no file inside

## Configuration

By putting a `config.yaml` file in same directory of executable, user can specific behavior:

| Config Key | Usage                                    | Default value (No config found) |
| ---------- | ---------------------------------------- | ------------------------------- |
| files      | Delete file when exact match to filename | `.DS_Store`                     |
| extensions | Delete file that has that extension      | None                            |

A example of `config.yaml` can be found in `.examples/` of this repo.

If no config file is found, then default value will be used.