#include <iostream>

using namespace std;

int main(){
	int num = 0;
	int maxTime = 0;
	int temp = 0;
	cin >> num >> maxTime;

	int* array1 = new int[num];

	for (int i = 0; i < num; i++)
	{
		cin >> array1[i];
	}
	int a = 0;
	bool run = true;
	while (run)
	{
		if ((temp + array1[a]) >= maxTime+1 || a == num)
		{

			run = false;
		}
		else
		{
			temp += array1[a];
			a++;
		}
	}
	cout << a;
	delete[] array1;
	return 0;
}