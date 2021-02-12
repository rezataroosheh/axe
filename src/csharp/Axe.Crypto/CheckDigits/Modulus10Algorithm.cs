using System;
using Axe.Crypto.Abstraction;

namespace Axe.Crypto.CheckDigits
{
    public class Modulus10Algorithm : ICheckDigitAlgorithm
    {
        public Modulus10Algorithm()
        {
        }

        public virtual byte Compute(ReadOnlySpan<byte> data)
        {
            var sum = 0;
	        var parity = data.Length % 2;
            var i = 0;

            foreach (var number in data)
            {
                if (number > 9)
			        throw new ArgumentOutOfRangeException("The numbers must be less than 10");
		        var value = number * (1 + (i % 2) ^ parity);
                
                if (value > 9)
                    value -= 9;
                sum += (int)value;
                i++;
            }

	        sum %= 10;
            if (sum == 0) {
                return 0;
            }
	        return (byte)(10 - sum);
        }

        public virtual string GetName() => "modulus10";
    }
}