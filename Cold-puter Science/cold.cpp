#include <iostream>

using namespace std;

int main() {
	int ltz = 0;
	int num = 0;
	int temp = 0;
	cin >> num;
	for (size_t i = 0; i < num; i++)
	{
		cin >> temp;
		if (temp <0)
		{
			ltz++;
		}
	}
	cout << endl << ltz;
	return 0;
}
