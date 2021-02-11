using System;

namespace Axe.Crypto.AnsiMacs
{
    public interface IAnsiMacAlgorithm : IDisposable
    {
        ReadOnlySpan<byte> Compute(ReadOnlySpan<byte> data);
	    string GetName();
    }
}