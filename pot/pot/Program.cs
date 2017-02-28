using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace pot
{
    class Program
    {
        static void Main(string[] args)
        {
            double total = 0;
            double final = 0;
            List<string> sarr = new List<string>();
            double.TryParse(Console.ReadLine(), out total);
            for (int i = 0; i < total; i++)
            { 
                sarr.Add(Console.ReadLine());
            }
            for (int i = 0; i < sarr.Count; i++)
            {
                double num = 0;
                num = double.Parse(sarr[i].Substring(0, sarr[i].Count() -1));

                double pow = 0;
                pow = double.Parse(sarr[i].Substring(sarr[i].Count() - 1));

                final += Math.Pow(num,pow);
            }
            Console.WriteLine(final);
        }
    }
}
