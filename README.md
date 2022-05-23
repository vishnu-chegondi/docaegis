# DocAegis

DocAegis is cli tool to protect your directories from accidental termination.

## Planning

* database
* list
* gaurd

### start(later):

 Start sub-command runs a daemon process to check regularly all the directories.


### database

Sqlite database to store the directories information

### list

To list all the directories and files which are guarded.


### gaurd: 
- **Input:** Sub command which takes input as path to directory/file which should be protected from deleting.
- **Working:** Create a hidden directory *.aegis* in the base directory of given file/directory.
- **How:** When a user deletes the directory by mistake we need to run recover subcommand to recreate the directory.
