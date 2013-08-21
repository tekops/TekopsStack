TekopsStack
===========


inotifywait -m -r -e close_write --exclude 'Tek*' . | while read line; do go build . && ./TekopsStack -v=2 -logtostderr; echo -e "\n\n";done