# goFMDataMigration

goFMDataMigration is a command line tool for easy data migration using the FMDataMigration tool. FileMaker is a trademark of FileMaker, Inc., registered in the U.S. and other countries.

## Intro

* Copy the source file of the production server to the `resources/prod` folder.
* Copy the development server's clone file to the `resources/clone` folder.
* Execution of the command will be saved in the `resources/migrated` folder.
* For example, `gofmdatamigration -a account -p password -k key -iv -ia -if`
* Note: [FMDataMigration](https://fmhelp.filemaker.com/docs/edition/en/migration/index.html) is required.

## Getting Started

1. You must have a [Go](http://golang.org) compiler installed.
2. Download [FMDataMigration](https://fmhelp.filemaker.com/docs/edition/en/migration/index.html) and move to the project folder.
3. Download and build goFMDataMigration: `go get github.com/frudens/gofmdatamigration`
4. Either copy the `gofmdatamigration` executable in `$GOPATH/bin` to a directory in
   your `PATH`, or add `$GOPATH/bin` to your `PATH`.
5. Download the folder or file to be migrated from the production server and copy it to the `resources/prod` folder.
6. Prepare the clone file from the backup of the development server and copy it to the `resources/clone` folder.
7. Run `gofmdatamigration -a account -p password`
8. The same file as the prod folder is created in the `resources/migrated` folder.


## Download

You can download it from the [release](https://github.com/frudens/gofmdatamigration/releases) page.

## Usage

Create a folder for data migration.

```
$ cd ~/Desktop
$ mkdir goFMDataMigration
$ cd goFMDataMigration
$ pwd
/Users/user/Desktop/goFMDataMigration
```

Move or copy source and clone files.
The following is a sample.

```
$ tree
.
├── resources
│   ├── clone
│   │   ├── dir1
│   │   │   └── Inventory Clone.fmp12
│   │   ├── dir2
│   │   │   └── Meetings Clone.fmp12
│   │   ├── Contacts Clone.fmp12
│   │   └── Tasks Clone.fmp12
│   └── prod
│       ├── dir1
│       │   └── Inventory.fmp12
│       ├── dir2
│       │   └── Meetings.fmp12
│       ├── Contacts.fmp12
│       └── Tasks.fmp12
├── FMDataMigration
└── gofmdatamigration
```

Execute the command.

```
$ gofmdatamigration -a admin -p ""

resources/prod/Contacts.fmp12
----------
Start: Fri Aug 24 08:25:57 2018
 == Mapping source privileges to target privileges ==
 == Copying changed custom value lists ==
 == Analyzing font lists ==
 == Mapping source tables to target tables ==
 == Mapping fields in source table "Phone Numbers" to target table "Phone Numbers" ==
  -- Block mode migration for source table "Phone Numbers" --
 == Mapping fields in source table "Email Addresses" to target table "Email Addresses" ==
  -- Block mode migration for source table "Email Addresses" --
 == Mapping fields in source table "Addresses" to target table "Addresses" ==
  -- Block mode migration for source table "Addresses" --
 == Mapping fields in source table "Contacts" to target table "Contacts" ==
  -- Block mode migration for source table "Contacts" --
 == Summary ==
  Accounts migrated: 2
  Accounts changed: 0
  Custom value lists migrated: 0
  Font entries added: 0
  Tables migrated: 4
  Tables not migrated: 0
  Fields migrated: 42
  Fields not migrated: 0
  Fields triggering recalculations: 0
  Fields with evaluation errors: 0
  Fields with fewer repetitions: 0
  Serial numbers updated: 0
End: Fri Aug 24 08:25:57 2018


resources/prod/Tasks.fmp12
----------
Start: Fri Aug 24 08:26:00 2018

Continue.

```

A migrated folder is created with the same configuration as the prod folder.

```
$ tree
.
├── resources
│   ├── clone
│   │   ├── dir1
│   │   │   └── Inventory Clone.fmp12
│   │   ├── dir2
│   │   │   └── Meetings Clone.fmp12
│   │   ├── Contacts Clone.fmp12
│   │   └── Tasks Clone.fmp12
│   ├── prod
│   │   ├── dir1
│   │   │   └── Inventory.fmp12
│   │   ├── dir2
│   │   │   └── Meetings.fmp12
│   │   ├── Contacts.fmp12
│   │   └── Tasks.fmp12
│   └── migrated
│       ├── dir1
│       │   └── Inventory.fmp12
│       ├── dir2
│       │   └── Meetings.fmp12
│       ├── Contacts.fmp12
│       └── Tasks.fmp12
├── FMDataMigration
└── log.txt

10 directories, 14 files

```

The log is saved in log.txt.

```
resources/prod/Contacts.fmp12
----------
Start: Fri Aug 24 08:25:57 2018
 == Mapping source privileges to target privileges ==
 == Copying changed custom value lists ==
 == Analyzing font lists ==
 == Mapping source tables to target tables ==
 == Mapping fields in source table "Phone Numbers" to target table "Phone Numbers" ==
  -- Block mode migration for source table "Phone Numbers" --
 == Mapping fields in source table "Email Addresses" to target table "Email Addresses" ==
  -- Block mode migration for source table "Email Addresses" --
 == Mapping fields in source table "Addresses" to target table "Addresses" ==
  -- Block mode migration for source table "Addresses" --
 == Mapping fields in source table "Contacts" to target table "Contacts" ==
  -- Block mode migration for source table "Contacts" --
 == Summary ==
  Accounts migrated: 2
  Accounts changed: 0
  Custom value lists migrated: 0
  Font entries added: 0
  Tables migrated: 4
  Tables not migrated: 0
  Fields migrated: 42
  Fields not migrated: 0
  Fields triggering recalculations: 0
  Fields with evaluation errors: 0
  Fields with fewer repetitions: 0
  Serial numbers updated: 0
End: Fri Aug 24 08:25:57 2018

resources/prod/Tasks.fmp12
----------
Start: Fri Aug 24 08:26:00 2018
 == Mapping source privileges to target privileges ==
 == Copying changed custom value lists ==
 == Analyzing font lists ==
 == Mapping source tables to target tables ==
 == Mapping fields in source table "Attachments" to target table "Attachments" ==
  -- Block mode migration for source table "Attachments" --
 == Mapping fields in source table "Assignments" to target table "Assignments" ==
  -- Block mode migration for source table "Assignments" --
 == Mapping fields in source table "Tasks" to target table "Tasks" ==
  -- Block mode migration for source table "Tasks" --
 == Mapping fields in source table "Assignees" to target table "Assignees" ==
  -- Block mode migration for source table "Assignees" --
 == Summary ==
  Accounts migrated: 2
  Accounts changed: 0
  Custom value lists migrated: 0
  Font entries added: 0
  Tables migrated: 4
  Tables not migrated: 0
  Fields migrated: 42
  Fields not migrated: 0
  Fields triggering recalculations: 0
  Fields with evaluation errors: 0
  Fields with fewer repetitions: 0
  Serial numbers updated: 0
End: Fri Aug 24 08:26:00 2018

resources/prod/dir1/Inventory.fmp12
----------
Start: Fri Aug 24 08:26:02 2018
 == Mapping source privileges to target privileges ==
 == Copying changed custom value lists ==
 == Analyzing font lists ==
 == Mapping source tables to target tables ==
 == Mapping fields in source table "Products" to target table "Products" ==
  -- Block mode migration for source table "Products" --
 == Mapping fields in source table "Inventory Transactions" to target table "Inventory Transactions" ==
  -- Block mode migration for source table "Inventory Transactions" --
 == Summary ==
  Accounts migrated: 2
  Accounts changed: 0
  Custom value lists migrated: 0
  Font entries added: 0
  Tables migrated: 2
  Tables not migrated: 0
  Fields migrated: 32
  Fields not migrated: 0
  Fields triggering recalculations: 0
  Fields with evaluation errors: 0
  Fields with fewer repetitions: 0
  Serial numbers updated: 0
End: Fri Aug 24 08:26:02 2018

resources/prod/dir2/Meetings.fmp12
----------
Start: Fri Aug 24 08:26:04 2018
 == Mapping source privileges to target privileges ==
 == Copying changed custom value lists ==
 == Analyzing font lists ==
 == Mapping source tables to target tables ==
 == Mapping fields in source table "Topics" to target table "Topics" ==
  -- Block mode migration for source table "Topics" --
 == Mapping fields in source table "Action Items" to target table "Action Items" ==
  -- Block mode migration for source table "Action Items" --
 == Mapping fields in source table "Meetings" to target table "Meetings" ==
  -- Block mode migration for source table "Meetings" --
 == Summary ==
  Accounts migrated: 2
  Accounts changed: 0
  Custom value lists migrated: 0
  Font entries added: 0
  Tables migrated: 3
  Tables not migrated: 0
  Fields migrated: 33
  Fields not migrated: 0
  Fields triggering recalculations: 0
  Fields with evaluation errors: 0
  Fields with fewer repetitions: 0
  Serial numbers updated: 0
End: Fri Aug 24 08:26:04 2018
```

## Author

frudens Inc. <https://frudens.com>

## License

This software is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt for more information.
