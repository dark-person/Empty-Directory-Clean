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
