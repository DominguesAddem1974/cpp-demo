## Conan + Cmake + Gtest

## Connan Install

```bash
brew install conan
```

## Install GoogleTest

```Bash
conan search gtest --remote=conancenter
conan inspect gtest/1.11.0
conan profile new default --detect
## clang
conan profile update settings.compiler.libcxx=libc++ default
mkdir build && cd build
conan install .. --build=gtest
cmake ..
```

## VSCode Conan Support

```bash
Add set(CMAKE_EXPORT_COMPILE_COMMANDS ON) to your CMakeLists.txt (or add to cmake: cmake -DCMAKE_EXPORT_COMPILE_COMMANDS=ON ..) so a build/compile_commands.json will be generated.

VS Code (clion, etc) can utilize this file to support auto complete:

$ cat .vscode/c_cpp_properties.json
{
    "configurations": [
    {
        "name": "Linux",
            "defines": [],
            "compilerPath": "/usr/bin/g++",
            "cStandard": "c11",
            "cppStandard": "c++14",
            "intelliSenseMode": "clang-x64",
            "compileCommands": "${workspaceFolder}/build/compile_commands.json"
    }
    ],
    "version": 4
}
```