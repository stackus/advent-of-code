# advent-of-code
https://adventofcode.com/ 

## Requirements
- Go 1.16+ (https://golang.org/doc/install)
- Taskfile (https://taskfile.dev/#/installation)

## Installation
Be sure Go has been installed and is available in your PATH.
Verify by running `go version` in your terminal.
Next, make sure Taskfile is installed and available in your PATH.
Verify by running `task --version` in your terminal.

### Setup
Create a `.env.local` file in the root of the project and add the following environment variables:
```
AOC_SESSION=<your session cookie from adventofcode.com>
```
You can find this cookie by logging into adventofcode.com and inspecting the request headers for any request made to the site.
The cookie will be named `session`.

## Usage
### Prepping for the current days puzzle
```bash
task
```
This will create a new directory for the current puzzle containing a go file ready for you to fill with your solution.
It will also create a `input.txt` file containing the puzzle input.
The puzzle description will also be downloaded into `puzzle.md` for reference.

You can also specify a day and even a year to initialize past puzzles.
```bash
# Initialize puzzle for day 1 of the current year
task DAY=1
# Initialize puzzle for day 25 of 2022
task DAY=25 YEAR=2022
```

### Individual commands
The default task will run the puzzle `init`, `input`, and `puzzle` commands. You can run these individually as well.
```bash
# Get the input for the current day
task input
# Initialize the puzzle for the next day (assuming today is 1)
task init DAY=2
# get the puzzle description for a specific day
task puzzle DAY=1 YEAR=2020
```

Use `task --list` to see all available tasks.

## Solve the puzzle
Edit the bodies of the `puzzle1`, `puzzle2`, and `parseInput` functions to solve the puzzle.

To then run your solution for the current puzzle, run the following command:
```bash
go run <YEAR>/day<DAY> -puzzle <PUZZLE_NUMBER>
```
This will create a `solution-<PUZZLE_NUMBER>.txt` file in the directory of the puzzle.
You can then submit this file to adventofcode.com to get your stars!
```bash
# Submit the solution for puzzle 1 for the current day
task submit PUZZLE=1
# Submit the solution for puzzle 2 for day 25 of 2022
task submit PUZZLE=2 DAY=25 YEAR=2022
```
The response from the server will be printed to the console and saved into a file for quick reference.
