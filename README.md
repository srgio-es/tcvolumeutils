# **tcvolumeutils**

[![Go](https://github.com/srgio-es/tcvolumeutils/actions/workflows/go.yml/badge.svg)](https://github.com/srgio-es/tcvolumeutils/actions/workflows/go.yml)

**Teamcenter Volume Management utils** is a set of cmd tools under development to help Siemens Teamcenter Administrators with Volume Management tasks.

## **CURRENT FEATURES**

* Reads the content of all .txt files generated by review_volumes and extracts OS Missing files into a nice and readable XLSX file.
* Reads the content of one or all .txt files in a folder generated by review_volumens and creates an almost empty file for all the files not found that are referenced by a datased in the location specified so you can purge them.

## **HOW TO USE**

```txt
tcvolumeutils -h

Teamcenter Volume management utils is a set of tools that will help Teamcenter administrator to perform some TC volumes management tedious tasks in an easier way.

Usage:
  tcvolumeutils [command]

Available Commands:
  createmissing Creates missing files with empty content found using review_volumes command. Useful when trying to clean volumes.
  help          Help about any command
  reportmissing Extracts missing OS files from review_volumes logs

Flags:
  -h, --help      help for tcvolumeutils
  -v, --verbose   increase verbosity for commands
```

### **createmissing:**
```txt
tcvolumeutils createmissing -h

This command creates empty files with the same exact name and locations as the missing found files.
This is helpful when administering volumes with commands such as dataset_cleanup or purge_datasets as they
will fail in the event of a missing reference.

Usage:
  tcvolumeutils createmissing [flags]

Flags:
  -h, --help                 help for createmissing
  -l, --log-file string      Specifies the log file to be processed
  -f, --logs-folder string   Specifies the location of the logs to be processed

Global Flags:
  -v, --verbose   increase verbosity for commands
```

Examples:
```
tcvolumeutils createmissing -l C:\Temp\VOL1.txt ##For unique file processing
tcvolumeutils createmissing -f C:\Temp ##To process all files within a folder
```

### **reportmissing:**

```txt
tcvolumeutils reportmissing -h

This command extracts missing OS files info from the logs generated by the review_volumes command.

The information is processed and writed into a XLXS file to improve readiness.
If the XLSX file specified contains data, the file is updated appending the new values.

Usage:
  tcvolumeutils reportmissing [flags]

Flags:
  -h, --help                 help for reportmissing
  -f, --logs-folder string   Specifies the location of the logs to be processed
  -r, --report string        Specifies the path to the XLSX file to populate with the results. (default "volumes-report.xlsx")

Global Flags:
  -v, --verbose   increase verbosity for commands
```

Examples:

```txt
tcvolumeutils reportmissing -f C:\reports\reviewvolumes\202120120 -r C:\results\reviewvolumes_202120120.xlsx -v
```

## **DISCLAIMER**

This piece of software manteined and released by the community and is not supported nor has any relation with Siemens or any of its subsidiaries.
