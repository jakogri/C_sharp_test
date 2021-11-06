using System;

namespace CalcArea
{
   public class CalcArea
    {
        public double GetFigureArea(double[] arr)
        {
            GeometricShapes gs = null;
            double area = 0;
            int arr_count = arr.Length;
            try
            {
                switch (arr_count)
                {
                    case 1:
                        gs = new Circle(arr[0]);
                        break;
                    case 3:
                        gs = new Triangle(arr[0], arr[1], arr[2]);
                        break;

                    default:
                        break;
                }

                if (gs != null) area = gs.Area();
            }
            catch (Exception)
            {
                if (gs != null) gs = null;
                return (0);
            }

            return (area);

        }


    }
}
