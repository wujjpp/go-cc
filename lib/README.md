# 编译静态库

g++ -c MurmurHash3.cc
g++ -c bridge.c
ar -crs libhash64.a MurmurHash3.o bridge.o