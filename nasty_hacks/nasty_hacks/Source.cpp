#include <iostream>
using namespace std;

int main() {
	int iter, r, e, c;

	cin >> iter;
	for (size_t i = 0; i < iter; i++)
	{
		cin >> r >> e >> c;
		if ((r + c) < e)
		{
			cout << "advertise" << endl;
		}
		else if ((r + c) > e)
		{
			cout << "do not advertise" << endl;
		}
		else
		{
			cout << "does not matter" << endl;
		}
	}
}