# bgTimer
This is my first project using Go. This program is used as a timer which notifies the user via a desktop notification 
as soon as the provided time duration has passed. Tested on GNU/Linux with GNOME.
---
### How to run bgTimer
```shell
bash bgTimer.sh -t $time_amount -u $time_unit -r $timer_reason
```
time_unit and reason are not necessary. The standard unit is seconds and standard reason "Timer has finished".