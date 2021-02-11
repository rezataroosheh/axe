using System;

namespace Axe.Crypto.AnsiMacs
{
    public static class CryptoUtilities
    {
        public static void Xor(Span<byte> source, ReadOnlySpan<byte> value)
        {
            if(value.Length>source.Length)
                throw new InvalidOperationException("The length of value is greater than source length.");

            for (int i = 0; i < value.Length; i++)
            {
                source[i] ^= value[i];
            }
        }
    }
}