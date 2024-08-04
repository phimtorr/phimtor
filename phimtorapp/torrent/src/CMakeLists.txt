# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(torrent_library VERSION 0.0.1)

add_custom_command(
    OUTPUT ./include/libtorrent.so
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMAND go build -o ./include/libtorrent.so -buildmode=c-shared main.go
    COMMENT "Building Go shared library"
)

add_custom_target(
    go_build ALL
    DEPENDS ./include/libtorrent.so
)

add_library(torrent STATIC IMPORTED GLOBAL)
add_dependencies(torrent go_build)

set_target_properties(torrent PROPERTIES
    IMPORTED_LOCATION ${CMAKE_CURRENT_SOURCE_DIR}/include/libtorrent.so
    INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_SOURCE_DIR}/build
    OUTPUT_NAME "torrent"
)

if (WIN32)
set_target_properties(torrent PROPERTIES
  WINDOWS_EXPORT_ALL_SYMBOLS ON
)
endif (WIN32)

# Creating a dummy target to apply compile definitions
add_library(torrent_defs INTERFACE)
target_compile_definitions(torrent_defs INTERFACE DART_SHARED_LIB)

# Use the dummy target's compile definitions with the torrent target
add_dependencies(torrent torrent_defs)
