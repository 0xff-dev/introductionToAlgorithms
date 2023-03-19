package challengeprogrammingdatastructure

/*
#include <stdio.h>
#include <stdlib.h>

	int search(int *arr, int n, int key) {
	    arr[n] = key;
	    int i = 0;
	    while(arr[i] != key)i++;
	    return i != n;
	}

	int main() {
	    int n, q, sum = 0, i, key;
	    scanf("%d", &n);
	    int *arr = (int*)malloc(sizeof(int) * (n+1));
	    for (i = 0; i < n; i++) {
	        scanf("%d", &arr[i]);
	    }
	    scanf("%d", &q);
	    for(i = 0; i < q; i++) {
	        scanf("%d", &key);
	        if(search(arr, n, key)) {
	            sum++;
	        }
	    }
	    free(arr);
	    printf("%d\n", sum);
	    return 0;
	}
*/
func LinearSearch(s, t []int) []int {
	ans := make([]int, 0)
	for _, i := range s {
		for _, j := range t {
			if i == j {
				ans = append(ans, i)
			}
		}
	}
	return ans
}
