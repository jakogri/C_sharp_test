using System;

namespace CalcArea
{
    class Circle : GeometricShapes //класс окружности
    {
        private double radius;

        public double Radius
        {
            get { return radius; }
            set
            {
                radius = value > 0 ? value : -value;
            }
        }

        protected Circle() : base("Circle")
        {

        }

        public Circle(double radius) : this()
        {
            Radius = radius;
        }

        public override double Area()
        {
            double Area = 0;
            Area = Math.PI * Radius * Radius;
            return (Area);
        }
    }
}
