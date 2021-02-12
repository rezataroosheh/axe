using System;
using Axe.Crypto.Abstraction;

namespace Axe.Crypto.CheckDigits
{
    public class VerhoeffAlgorithm : ICheckDigitAlgorithm
    {
        private static readonly int[][] multiplicationTable = new int[][]{
            new []{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
            new []{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
            new []{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
            new []{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
            new []{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
            new []{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
            new []{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
            new []{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
            new []{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
            new []{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	    };

        private static readonly int[][] permutationTable = new int[][]{
            new []{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
            new []{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
            new []{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
            new []{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
            new []{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
            new []{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
            new []{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
            new []{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
        };
	    private static readonly byte[] inverseTable = new byte[]{0, 4, 3, 2, 1, 5, 6, 7, 8, 9};

        public VerhoeffAlgorithm()
        {
        }

        public virtual byte Compute(ReadOnlySpan<byte> data)
        {
            var index = 0;
            var len = data.Length;
            for (int i = 0; i < len; i++)
            {
                if (data[i] > 9) {
                    throw new ArgumentOutOfRangeException("The numbers must be less than 10");
                }
                var firstIndex = (i + 1) % 8;
                var secondIndex = data[len-i-1];
                var permutation = permutationTable[firstIndex][secondIndex];
                index = multiplicationTable[index][permutation];
            }
            return inverseTable[index];
        }

        public virtual string GetName() => "verhoeff";
    }
}