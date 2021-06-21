# Journal-Go

This is a program to write [aes encrypted](https://golang.org/pkg/crypto/aes/) daily journal entries. One entry can be made every day, and that entry can only be edited during the day it's created.

In order to secure your entries you must use a 32-bit passphrase that can either be entered upon starting the program, or saved to a file at `.internal/.passphrase`.

This passphrase will be used for all the files in the `/entries` directory, so without it none can be decrypted.
