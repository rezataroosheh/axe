using System;

namespace Axe.Crypto.Abstraction
{
    public interface ICheckDigitAlgorithm
    {
        byte Compute(ReadOnlySpan<byte> data);
	    string GetName();
    }
}