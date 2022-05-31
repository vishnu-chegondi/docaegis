# DocAegis

DocAegis is cli tool to protect your files/directories from accidental termination in linux and MacOs. It recreates the files/directories which are gaurded by docaegis when deleted.

### Usage

Run below command to get all the usage details. All the other command requires **sudo** permissions to work.

```
docaegis -h
```

### Working

* When you gaurd any file, a hard link will be created in **.aegis** directory of the file location and all the file info will be stored inside **SQLLite** database at /var/lib/docaegis.db location in linux.
* When running restore command, all the information stored in SQlite database along with hardlink will be used to restore the file data and permissions.
* When running list it will list all the files which are gaured using the SQLite database.

### TODO:

* Gaurd all the files in the directory.
* Start sub-command runs a daemon process to check regularly all the directories and run restore to recreate all the files deleted.
