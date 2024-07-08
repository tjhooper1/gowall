on run argv
    set filePath to item 1 of argv
    tell application "Finder"
        set desktop picture to POSIX file filePath
    end tell
end run