using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Stacking_Cups
{
    class Program
    {
        static void Main(string[] args)
        {
            int loopittr = 0;
            int.TryParse(Console.ReadLine(), out loopittr);
            int d = 0;
            List<Tuple<int, string>> t = new List<Tuple<int, string>>();
            for (int i = 0; i < loopittr; i++)
            {
                string temp;
                string[] temparr;
                temp = Console.ReadLine();
                temparr = temp.Split(' ');
                bool test = int.TryParse(temparr[0], out d);
                if (!test)
                {
                    int.TryParse(temparr[1], out d);
                    t.Add(Tuple.Create((d*2), temparr[0]));
                }
                else
                {
                    t.Add(Tuple.Create(d, temparr[1]));
                }
            }
            t.Sort();
            for (int i = 0; i < t.Count; i++)
            {
                Console.WriteLine(t[i].Item2);
            }
        }
    }
}
