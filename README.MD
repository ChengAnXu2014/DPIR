# DPIR：The Intermediate Representation for Data eXchange Languages

Traditional data propagation generally use **DXL**(Data eXchange Language e.g. **XML**, **YAML**, **JSON**) as the data carrier.

This approach have two penalties:
1. **DXL**s are very high-level, convenient for user to edit but Inconvenient for program to parse.
2. There are too many **DXL**s exist, programs use different **DXL**s can't exchange data directly.

Here I introduce a better way for data propagation: design a intermediate representation for data propagation.
I name it **DPIR**(Data Propagating Intermediate Representation).

**DPIR** should be low-level enough for programs to access efficiently and expressive enough for transforming to and from **DXL**s with no information lose.







# Preliminary Design of DPIR

String encoding is another problem for data propagation, to avoid this problem **DPIR** save strings in a individual file: *StrFile*.
Non-string data is saved in *DataFile*.



### StrFile

*StrFile*'s name form is: `name.dpir-str-encode`.
Where `name` is the source file's name and `encode` is the string encoding used.

For example: **DPIR** from `config.json`, it's *StrFile*'s name should be `config.dpir-str-encode`.

For the *StrFile* saved with **UTF-8** it's name should be `config.dpir-str-utf8`;

For the *StrFile* saved with **UTF-16 LE** it's name should be `config.dpir-str-utf16le`;

For the *StrFile* saved with **UTF-16 BE** it's name should be `config.dpir-str-utf16be`.

There can be multiple *StrFiles*'s saved with different encodings in a single **DPIR**.

The strings in *StrFile* are terminated with `\n` and accessed with index started from 1, the index 0 is used to represent empty string("").
Therefore the line number of a line is also the index of the string saved in this line.



### DataFile

*DataFile*'s name form is: `name.dpir-data` where `name` is source file's name.

##### The code in file "dataLayout.go" illustrates *DataFile*'s format.