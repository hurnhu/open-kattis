using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FizzBuzz
{
    class Program
    {
        static void Main(string[] args)
        {
            double x = 0;
            double y = 0;
            int end = 0;
            string temp;
            string[] temparr;

            temp = Console.ReadLine();
            temparr = temp.Split(' ');
            double.TryParse(temparr[0], out x);
            double.TryParse(temparr[1], out y);
            int.TryParse(temparr[2], out end);

            for (int i = 1; i <= end; i++)
            {
                if (i % y == 0 && i % x == 0)
                {
                    Console.WriteLine("FizzBuzz");
                }
                else if (i % x == 0)
                {
                    Console.WriteLine("Fizz");
                }
                else if (i % y == 0)
                {
                    Console.WriteLine("Buzz");
                }
                else
                {
                    Console.WriteLine(i);
                }
            }
            Console.ReadLine();
        }
    }
}
