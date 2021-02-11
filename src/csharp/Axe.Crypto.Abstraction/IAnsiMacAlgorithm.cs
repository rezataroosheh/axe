using System;

namespace Axe.Crypto.Abstraction
{
    public interface IAnsiMacAlgorithm : IDisposable
    {
        ReadOnlySpan<byte> Compute(ReadOnlySpan<byte> data);
	    string GetName();
    }
}