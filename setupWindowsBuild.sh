#!/bin/bash

CC=x86_64-w64-mingw32-gcc-win32
if ! type $CC &> /dev/null 
then
	echo "Missing mingw compiler cmd: $CC"
	echo "On ubuntu 19.04+:"
	echo "sudo apt-get install mingw-w64"
	exit 1
fi

#install headers
wget https://github.com/KhronosGroup/OpenCL-Headers/archive/master.zip
unzip master.zip
cp -R OpenCL-Headers-master/CL/ /usr/x86_64-w64-mingw32/include/
rm -rf OpenCL-Headers-master/
rm master.zip

#build openCL library for mingw
wget https://github.com/KhronosGroup/OpenCL-ICD-Loader/raw/master/loader/windows/OpenCL.def
sed -i '1 i\LIBRARY OpenCL.dll' OpenCL.def
x86_64-w64-mingw32-dlltool -d OpenCL.def -l libOpenCL.a
mv libOpenCL.a /usr/x86_64-w64-mingw32/lib/
rm OpenCL.def





