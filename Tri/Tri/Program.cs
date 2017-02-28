using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Tri
{
    class Program
    {
        static void Main(string[] args)
        {
            int a, b, c;
            string temp;
            string[] temparr;

            temp = Console.ReadLine();
            temparr = temp.Split(' ');

            int.TryParse(temparr[0], out a);
            int.TryParse(temparr[1], out b);
            int.TryParse(temparr[2], out c);

            if ((a+b)==c || (b+c)==a)
            {
                if ((a + b) == c)
                {
                    Console.WriteLine(a + "+" + b + "=" + c);
                }
                else
                {
                    Console.WriteLine(a + "=" + b + "+" + c);
                }
            }
        }
    }
}
