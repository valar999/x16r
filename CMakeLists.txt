cmake_minimum_required(VERSION 3.5)

if (NOT CMAKE_BUILD_TYPE)
    message(STATUS "No build type selected, defaulting to Release")
    set(CMAKE_BUILD_TYPE "Release")
endif ()

set(SOURCES
      x16r.c
      sha3/blake.c
      sha3/bmw.c
      sha3/groestl.c
      sha3/jh.c
      sha3/keccak.c
      sha3/skein.c
      sha3/cubehash.c
      sha3/echo.c
      sha3/luffa.c
      sha3/simd.c
      sha3/hamsi.c
      sha3/hamsi_helper.c
      sha3/fugue.c
      sha3/shavite.c
      sha3/shabal.c
      sha3/whirlpool.c
      sha3/sha2big.c
)

add_library(x16r_hash STATIC ${SOURCES})

