using System;

namespace Axe.Crypto.AnsiMacs
{
    public class Ansi919MacAlgorithm : RetailMacAlgorithm
    {
        public Ansi919MacAlgorithm(byte[] key1, byte[] key2)
            :base(key1, key2, key1)
        {
        }

        public override string GetName() => "Ansi 9.19 MAC";
    }
}