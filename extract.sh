#!/bin/bash
cd modules
for zip in *.zip; do
	dirname=$(echo $zip | cut -d'-' -f1)
	echo "--> Extracting $dirname..."
	mkdir $dirname
	7z x -o$dirname $zip > /dev/null
	rm $zip
done
cd ..

echo "==> Done."
