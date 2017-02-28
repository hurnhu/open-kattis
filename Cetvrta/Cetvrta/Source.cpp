#include <iostream>
#include <string>
struct point
{
	int x = 0;
	int y = 0;
};
using namespace std;
int main()
{
	point a, b, c, d;
	string input;
	cin >> a.x >> a.y;
	cin >> b.x >> b.y;
	cin >> c.x >> c.y;
	if (a.x == b.x)
	{
		d.x = c.x;
	}
	else if (b.x == c.x)
	{
		d.x = a.x;
	}
	else
	{
		d.x = b.x;
	}

	if (a.y == b.y)
	{
		d.y = c.y;
	}
	else if (b.y == c.y)
	{
		d.y = a.y;
	}
	else
	{
		d.y = b.y;
	}

	cout << d.x << " " << d.y << endl;
	return 0;
}