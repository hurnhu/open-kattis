#include <iostream>

using namespace std;


int main() {
	int num;
	cin >> num;
	((num % 2 == 0) ? cout << "Bob" : cout << "Alice");
	return 0;
}
