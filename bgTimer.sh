#!/bin/bash
# get flags
while getopts t:u:r: flag
do
  case "${flag}" in
    t) time_value=$OPTARG;;
    u) time_unit=$OPTARG;;
    r) reason=$OPTARG;;
    *) echo "Valid flags:
       -t   set time amount
       -u   set time unit (sec, min or h)
       -r   set timer reason"; exit 1;;
  esac
done

# if ./BackgroundTimer does not exist, build it.
[ -f ./BackgroundTimer ] || go mod download && go build .

# initialize argument array with time duration amount
arg_array=(-t "$time_value")

# add time unit if it has been provided by the user
if [ -n "$time_unit" ]; then
  arg_array+=(-u "$time_unit")
fi

# add timer reason if it has been provided by the user
if [ -n "$reason" ]; then
  arg_array+=(-r "$reason")
fi

# run ./BackgroundTimer with provided arguments as job
"./BackgroundTimer" "${arg_array[@]}" &