#include <iostream>
using namespace std;

int main() {
        int h, m, t = 45;
        cin >> h >> m;

        if (m < t) {
            (h == 0) ? h = 24 : 0;
            h--;
            t -= m;
            m = 60 - t;
        }else {
            m -= t;
        }
        cout << h << " " << m << endl;
    return 0;
}