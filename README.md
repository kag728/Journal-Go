# Journal-Go

This is a program to write [aes encrypted](https://golang.org/pkg/crypto/aes/) daily journal entries. One entry can be made every day, and that entry can only be edited during the day it's created.

In order to secure your entries you must use a 32-bit passphrase that can either be entered upon starting the program, or saved to a file at `.internal/.passphrase`.

This passphrase will be used for all the files in the `/entries` directory, so without it none can be decrypted.

# Getting Started

Steps
  1. Download and extract release archive.
  2. Run with `go run .`
  3. Choose a passphrase.
      * Either type in a 32-character passphrase when prompted.
      * Or store a 32-character long passphrase in the project directory in a file called `.internal/.passphrase`.
          * If the stored passphrase is less than 32 characters, you will be prompted for the remaining bytes. This can be used to set up a "pin" by leaving the last few bytes out of the `.passphrase` file so you are still required to provide a short password whenever starting the program, but do not have to type the whole 32-character passphrase every time.
  4. Choose `w` to create your first entry.
  5. Open `entries/editor` to write your entry, save, then press enter on terminal.
      * This will delete the editor file and save it's encrypted contents into a timestamped entry file in `entries/`.
  6. To read your entries, type `r` instead of `w` after creating your first entry. 
