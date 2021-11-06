using System;
using Xunit;
using CalcArea;

namespace CalcAreaTest
{
    public class CalcAreaTest
    {
        CalcArea.CalcArea ca;
        double[] arr;
        double arrea;
        public CalcAreaTest()
        {
            ca = new CalcArea.CalcArea();
        }
        [Fact]
        public void Test1()
        {
          try
            {
                arr = new double[] { 10, 2, 3 };
                try
                {
                    arrea = ca.GetFigureArea(arr);
                    throw new ArgumentOutOfRangeException("MyExeption");
                }
                catch (Exception ex)
                {
                    Assert.Equal("MyExeption", ex.Message);
                }
            }
          finally
            {
                arr = null;
            }
        }
        public void Test2()
        {
          try
            {
                arr = new double[] { 3, 4, 5 };
                arrea = ca.GetFigureArea(arr);
                Assert.True(arrea == 6);
            }
          finally
            {
                arr = null;
            }
        }
        public void Test3()
        {
            try
            {
                arr = new double[] { 1 };
                arrea = ca.GetFigureArea(arr);
                Assert.True(arrea == Math.PI);
            }
            finally
            {
                arr = null;
            }
        }
    }
}
