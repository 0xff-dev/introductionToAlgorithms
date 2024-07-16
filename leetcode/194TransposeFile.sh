# Read from the file file.txt and print its transposed content to stdout.
awk '{ for (i=1; i<=NF; i++) { a[i,NR] = $i } } NF>p { p = NF } END { for(j=1; j<=p; j++) { str=a[j,1]; for(i=2; i<=NR; i++){ str=str" "a[j,i]; } print str } }' file.txt
