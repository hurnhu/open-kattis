#include <iostream>

using namespace std;

int main() {
    int miles=0,hrs=0,a=0,b=0, total=0;
    cin >>a;
    while (a!=-1){
        b = 0;
        total = 0;
        for (int i = 0; i < a; ++i) {
            cin >> miles >> hrs;
            total += miles*(hrs-b);
            b = hrs;
        }
        cout << total << " miles" << endl;
        cin >>a;
    }
    return 0;
}
