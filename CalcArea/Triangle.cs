using System;
using System.Linq;

namespace CalcArea
{
    class Triangle : GeometricShapes
    {
        private double a;
        private double b;
        private double c;

        public double A
        {
            get { return a; }
            set
            {
                a = value > 0 ? value : -value;
            }
        }

        public double B
        {
            get { return b; }
            set
            {
                b = value > 0 ? value : -value;
            }
        }

        public double C
        {
            get { return c; }
            set
            {
                c = value > 0 ? value : -value;
            }
        }

        protected Triangle() : base("Triangle")
        {
            A = 1.0;
            B = 1.0;
            C = 1.0;
        }

        public Triangle(double a, double b, double c) : this()
        {
            if (a + b > c && a + c > b && b + c > a)
            {
                A = a;
                B = b;
                C = c;
            }
            else
            {
                Console.WriteLine("It is not a triangle - created default!");
                throw new ArgumentOutOfRangeException("Not a triangle");
            }
        }

        public override double Area()
        {
            double Area = 0;

            double HalfPerimeter = (A + B + C) / 2;
            Area = Math.Sqrt(HalfPerimeter * (HalfPerimeter - A) *
                                                (HalfPerimeter - B) *
                                                (HalfPerimeter - C));
            return (Area);
        }


        public bool CheckRightTriangle()
        {
            bool Result = false;

            double[] Sides = { A, B, C };

            int LongestSideIndex = Array.IndexOf(Sides, Sides.Max());

            double LongSidePow2 = 0;
            double SmallSidesPow2 = 0;

            for (int i = 0; i < Sides.Length; i++)
            {
                if (i != LongestSideIndex)
                {
                    SmallSidesPow2 += Math.Pow(Sides[i], 2);
                }
                else
                {
                    LongSidePow2 = Math.Pow(Sides[i], 2);
                }
            }

            if (LongSidePow2 == SmallSidesPow2)
            {
                Result = true;
            }

            return (Result);
        }
    }
}
