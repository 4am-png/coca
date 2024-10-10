cmake_minimum_required(VERSION 3.10)
project(db_core)

set(CMAKE_CXX_STANDARD 17)

add_executable(db_core src/main.cpp)
target_include_directories(db_core PRIVATE include)
