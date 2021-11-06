using System;

namespace CalcArea
{
    public abstract class GeometricShapes
    {
        private string name;
        public GeometricShapes(string Name)
        {
            name = Name;
        }
        public abstract double Area(); //расчёт площади
    }
}
