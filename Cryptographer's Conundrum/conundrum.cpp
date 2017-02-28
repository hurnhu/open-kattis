#include <iostream>
using namespace std;
int main() {
    string a;
    int b = 0, c = 0;
    cin >> a;

    for (int i = 0; i < a.length(); i++) {
        if(a[i] != 'P' && c == 0){
            b++;
        }else if (a[i] != 'E' && c == 1) {
            b++;
        }else if(a[i] != 'R' && c == 2){
            b++;
        }
        (c==2)? c=0:c++;
    }
    cout << b;
    return 0;
}
