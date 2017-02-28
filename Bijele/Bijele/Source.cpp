#include <iostream>

using namespace std;

int main() {
	int types[6] = { 1, 1, 2, 2, 2, 8 };
	int input[6];
	cin >> input[0] >> input[1] >> input[2] >> input[3] >> input[4] >> input[5];
	for (size_t i = 0; i < 6; i++)
	{
		if ((types[i] - input[i]) != 0)
		{
			cout << (types[i] - input[i]) << " ";
		}
		else
		{
			cout << 0 << " ";
		}
	}
	return 0;
}