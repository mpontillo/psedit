# psedit

## Introduction

`psedit` allows the editing of Phantasy Star saved games. The save files must
be dumps from a Sega Master System emulator, or a file-based dump of
battery-backed-up SRAM data from the original cartridge, such as the contents
of a `.ssm` file from an Everdrive.

`psedit` is the first piece of software I released, back in the late 1990s,
before the Internet became popular. It was distributed on FidoNet via dial-up
bulletin board systems.

After I found the long-lost source code on an old backup disk, I decided to
upload it to GitHub and make it open source.

My interest in this old project of mine was reinvigorated both because my
ten-year-old son started getting interested in the original Phantasy Star
game, and because I wanted to improve my Go language skills.

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
