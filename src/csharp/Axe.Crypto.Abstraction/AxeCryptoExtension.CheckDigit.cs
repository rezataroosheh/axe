using System;
using System.Buffers;
using System.Globalization;

namespace Axe.Crypto.Abstraction
{
    public static class AxeCryptoExtenssion
    {
        public static int ComputeCheckDigit(ICheckDigitAlgorithm algorithm, string input)
        {
            if (string.IsNullOrWhiteSpace(input))
                throw new ArgumentNullException(nameof(input));
            return ComputeCheckDigit(algorithm, input.AsSpan());
        }

        public static int ComputeCheckDigit(ICheckDigitAlgorithm algorithm, ReadOnlySpan<char> input)
        {
            using (var pool = MemoryPool<byte>.Shared.Rent(input.Length))
            {
                for (int i = 0; i < input.Length; i++)
                {
                    if (byte.TryParse(input.Slice(i, 1),
                                     NumberStyles.Integer,
                                     CultureInfo.InvariantCulture,
                                     out var digit))
                    {
                        pool.Memory.Span[i] = digit;
                    }
                    else
                    {
                        throw new ArgumentException("The value contains non-digit character.");
                    }
                }
                return algorithm.Compute(pool.Memory.Span);
            }
        }

        public static (bool, int) VerifyCheckDigit(ICheckDigitAlgorithm algorithm, string input)
        {
            if (string.IsNullOrWhiteSpace(input))
                throw new ArgumentNullException(nameof(input));
            return VerifyCheckDigit(algorithm, input.AsSpan());
        }

        public static (bool, int) VerifyCheckDigit(ICheckDigitAlgorithm algorithm, ReadOnlySpan<char> input)
        {
            if (input.Length < 2)
                throw new ArgumentException("Input is invalid");

            if (!byte.TryParse(input.Slice(input.Length - 1),
                                 NumberStyles.Integer,
                                 CultureInfo.InvariantCulture,
                                 out var checkDigit))
            {
                throw new ArgumentException("The value contains non-digit character.");
            }

            var computedCheckDigit = ComputeCheckDigit(algorithm, input.Slice(0, input.Length - 1));
            return (checkDigit == computedCheckDigit, computedCheckDigit);
        }
    }
}