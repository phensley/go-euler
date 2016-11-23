#!/bin/bash

# profiling broken on mac so use linux profile
go tool pprof ./go-euler-linux euler.hprof

