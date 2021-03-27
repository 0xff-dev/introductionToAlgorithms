cat file.txt| awk 'NR==10 {print $0}';
head -m filename | tail -1;
nl filename| sed -n 'mp' # nl filename |sed -n '100p'