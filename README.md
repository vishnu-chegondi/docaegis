# DocAegis

DocAegis is cli tool to protect your files/directories from accidental termination in linux and MacOs. It recreates the files/directories along with permissions which are guarded by docaegis when deleted.


Install the docaegis with the command ```go install github.com/vishnu-chegondi/docaegis@latest```. Go will install the package in your ```$GOPATH/bin``` directory which should be in ```$PATH```.

Once Installed you should have the ```docaegis``` command available.

## Usage

Run below command to get all the usage details. All the other command requires **sudo** permissions to work.

```
docaegis -h
```

## Working

### docaegis guard

**Flags**: -s --source string source_path which should be guarded

The ```docaegis guard -s source_path``` command will guard your files/directories in source_path. Under the hood when you guard any source path, hard links will be created in **.aegis** directory for all the files in source_path and all the files/directories info will be stored inside **SQLLite** database at **/var/lib/docaegis.db**.

e.g.

``` sh
docaegis guard -s /source/directory

ls -lart /source/*
drw-r--r--  40 root            staff  1280 Jun  7 17:34 .aegis
drwxr-xr-x   6 vishnuchegondi  staff   192 Jun  7 17:34 directory
```

### docaegis restore

**Flags**: -f --file string source_path which should be restored

The ```docaegis restore -f /source/directory``` command will restore your files/directories which are guarded with docaegis. Under the hood when running restore command, all the information stored in SQlite database along with hardlinks will be used to restore the the file data and permissions in the source path.

``` sh
docaegis restore -f /source/directory
```

### docaegis list

The ```docaegis list``` command will list all the source_paths which are guarded using docaegis.


Developed by @vishnu-chegondi