#include <iostream>
#include <algorithm>
#include <string>
using namespace std;
string StringToUpper(string strToConvert)
{
	std::transform(strToConvert.begin(), strToConvert.end(), strToConvert.begin(), ::toupper);

	return strToConvert;
}
int main() {
	string date = "";
	int number = 0;
	cin >> date >> number;
	date = StringToUpper(date);
	if ((date == "OCT" && number == 31 )|| (date == "DEC" && number == 24))
	{
		cout << "yup";
	}
	else
	{
		cout << "nope";
	}
}