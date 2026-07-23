#!/usr/bin/env bash

main() {
    local path="$1"

    if [ -z "$path" ]; then
        echo "Usage: $0 <path>"
        exit 1
    fi

    inotifywait --format '%w%f' -r -m "$path" -e create -e moved_to |
    while read -r file; do
        echo "File dropped: $file"
        if [[ "$file" == *.png ]]; then
            convert "$file" "${file%.png}.jpg"
            rm "$file"
        fi
    done
}

main "$@"
