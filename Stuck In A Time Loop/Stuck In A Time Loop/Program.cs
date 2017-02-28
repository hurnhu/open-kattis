using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks; 

namespace Stuck_In_A_Time_Loop
{
    class Program
    {
        static void Main(string[] args)
        {
            int num = 0;
            int.TryParse(Console.ReadLine(), out num);
            for (int i = 1; i <= num; i++)
            {
                Console.WriteLine(i + " Abracadabra");
            }
        }
    }
}
