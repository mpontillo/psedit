# psedit

## Introduction

`psedit` was the first piece of software I released, back in the late 1990s.

After I found the long-lost source code on an old backup disk, I decided to
upload it to GitHub and make it open source.

## Organization

This repository contains a Go port of `psedit` (currently a work-in-progress),
plus the source, build assets, release, and binary files for the legacy DOS
version of `psedit`.

Files for the new Go port are organized as follows:

    *.go - Implementation of the `psedit` go package.
    cmd/psedit - Go implementation of `psedit`.

The files for the legacy DOS port are organized as follows:

    pascal/ - The original Pascal source files and other assets.
    pascal/bin/ - Original DOS executables for psedit, plus a utility that
                  was used to patch a subtle bug in Turbo Pascal so that
                  psedit would run on more recent machines.
    pascal/backup/ - Original released psedit .zip, plus tppatch distribution
                     and documentation.
    pascal/backup/tpu - Turbo Pascal libraries used to build psedit.
                        (The source code to the "myio" library is still lost.)

This repository also contains the following files, which can be used with
either version:

    data/ - Example save files (for testing).
