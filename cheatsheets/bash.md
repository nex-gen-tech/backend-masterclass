# Bash Cheat Sheet

1. ls - list directory contents

- `ls -l` - display files in long format
- `ls -a` - show hidden files
- `ls -al` - show hidden files in long format

2. cd - change directory

- `cd /home/user` - change to home directory of user
- `cd ..` - move up one directory level
- `cd -` - go back to previous directory

3. pwd - print working directory

- `pwd -L` - show logical current directory
- `pwd -P` - show physical current directory
- `pwd -LP` - show both logical and physical current directory

4. cat - concatenate files and print on standard output

- `cat file1 file2` - display contents of file1 and file2
- `cat file1 > file2` - write contents of file1 to file2
- `cat file1 >> file2` - append contents of file1 to file2

5. grep - search for a pattern in a file or stream

- `grep "pattern" file` - search for "pattern" in file
- `grep -i "pattern" file` - search for "pattern" in file, ignoring case
- `grep -v "pattern" file` - search for lines that do not match "pattern" in file

6. sed - stream editor for filtering and transforming text

- `sed 's/old/new/' file` - replace "old" with "new" in file
- `sed '1,10d' file` - delete lines 1-10 in file
- `sed -n '/pattern/p' file` - print only lines that contain "pattern" in file

7. awk - pattern scanning and processing language

- `awk '{print $1, $2}' file` - print first and second field of each line in file
- `awk '/pattern/ {print $0}' file` - print lines that contain "pattern" in file
- `awk '$1 > 5' file` - print lines where first field is greater than 5 in file

8. head - display first lines of a file

- `head -n 5 file` - display first 5 lines of file
- `head -c 10 file` - display first 10 characters of file
- `head -q file1 file2` - suppress printing of filenames when displaying contents of file1 and file2

9. tail - display last lines of a file

- `tail -n 5 file` - display last 5 lines of file
- `tail -c 10 file` - display last 10 characters of file
- `tail -f file` - display last lines of file and follow updates to file

10. find - search for files in a directory hierarchy

- `find . -name "pattern"` - search for files with names matching "pattern" in current directory and subdirectories
- `find / -user "username"` - search for files owned by "username" in root directory and subdirectories
- `find . -type f -exec grep "pattern" {} ;` - search for "pattern" in all regular files in current directory and subdirectories

11. wc - display line, word, and byte counts for files

- `wc -l file` - display number of lines in file
- `wc -w file` - display number of words in file
- `wc -c file` - display number of bytes in file

12. cut - remove sections from each line of files

- `cut -d ',' -f 1 file` - display first field (delimited by ',') of each line in file
- `cut -d ',' -f 1-3 file` - display first, second, and third field (delimited by ',') of each line in file
- `cut -d ',' -f 1,3 file` - display first and third field (delimited by ',') of each line in file

13. sort - sort lines of text files

- `sort file` - sort lines in file in alphabetical order
- `sort -r file` - sort lines in file in reverse alphabetical order
- `sort -n file` - sort lines in file in numerical order

14. uniq - remove duplicated lines from a sorted file

- `uniq file` - remove duplicated lines in file
- `uniq -c file` - display number of occurrences of each line in file
- `uniq -d file` - display only duplicated lines in file

15. tr - translate or delete characters

- `tr 'a-z' 'A-Z' < file` - translate lowercase characters to uppercase in file
- `tr -d 'a-z' < file` - delete all lowercase characters from file
- `tr ' ' '\n' < file` - replace spaces with newlines in file

16. tar - create and manipulate tar archives

- `tar -cf archive.tar file1 file2` - create tar archive with file1 and file2
- `tar -xvf archive.tar` - extract files from tar archive
- `tar -zcvf archive.tar.gz file` - create gzipped tar archive with file

17. gzip - compress or decompress files

- `gzip file` - compress file
- `gzip -d file.gz` - decompress file.gz
- `gzip -l file.gz` - display information about compressed file.gz

18. gunzip - uncompress gzip files

- `gunzip file.gz` - uncompress file.gz
- `gunzip -l file.gz` - display information about compressed file.gz
- `gunzip -c file.gz` > file - uncompress file.gz and write to file

19. bzip2 - compress or decompress files

- `bzip2 file` - compress file
- `bzip2 -d file.bz2` - decompress file.bz2
- `bzip2 -k file` - compress file, but keep original file

20. bunzip2 - uncompress bzip2 files

- `bunzip2 file.bz2` - uncompress file.bz2
- `bunzip2 -c file.bz2 > file` - uncompress file.bz2 and write to file
- `bunzip2 -k file.bz2` - uncompress file.bz2, but keep original file

21. zip - package and compress files

- `zip archive.zip file1 file2` - create zip archive with file1 and file2
- `zip -d archive.zip file1` - delete file1 from zip archive
- `zip -u archive.zip file1` - update file1 in zip archive

22. unzip - extract files from a zip archive

- `unzip archive.zip` - extract all files from zip archive
- `unzip archive.zip file1` - extract file1 from zip archive
- `unzip -l archive.zip` - list contents of zip archive

23. rsync - synchronize files between two directories

- `rsync -av source/ destination/` - copy all files from source to destination, preserving file attributes and displaying progress
- `rsync -az source/ destination/`- copy all files from source to destination, preserving file attributes and compressing data transfer
- `rsync -av --delete source/ destination/` - copy all files from source to destination, deleting any files in destination that do not exist in source

24. cp - copy files and directories

- `cp file1 file2` - copy file1 to file2
- `cp -R dir1 dir2` - copy dir1 and its contents to dir2
- `cp -p file1 file2` - preserve file attributes when copying file1 to file2

25. mv - move or rename files and directories

- `mv file1 file2` - rename file1 to file2
- `mv file1 dir1` - move file1 to dir1
- `mv -i file1 file2` - prompt before overwriting file2 when moving file1 to file2

26. rm - remove files and directories

- `rm file` - delete file
- `rm -r dir` - delete dir and its contents
- `rm -i file` - prompt before deleting file

27. mkdir - create directories

- `mkdir dir1` - create directory dir1
- `mkdir -p dir1/dir2/dir3` - create directories dir1, dir2, and dir3, including any parent directories that don't exist
- `mkdir -m 755 dir1` - create directory dir1 with permissions 755

28. chmod - change file permissions

- `chmod 755 file` - change permissions of file to 755
- `chmod -R 755 dir` - change permissions of dir and its contents to 755
- `chmod u+x file` - add execute permission for owner of file

29. chown - change file ownership

- `chown user1 file` - change ownership of file to user1
- `chown -R user1 dir` - change ownership of dir and its contents to user1
- `chown user1:group1 file` - change ownership of file to user1 and group1

30. touch - change file timestamps or create empty files

- `touch file` - update file's modification and access timestamps to current time
- `touch -a file` - update file's access timestamp to current time
- `touch -c file` - do not create file if it does not exist, only update timestamps if it does exist
