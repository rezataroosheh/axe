using System;
using System.Security.Cryptography;

namespace Axe.Crypto.AnsiMacs
{
    public class RetailMacAlgorithm : Ansi99MacAlgorithm
    {
        private readonly ICryptoTransform _cryptoTransform2;
        private readonly ICryptoTransform _cryptoTransform3;

        public RetailMacAlgorithm(byte[] key1, byte[] key2, byte[] key3)
            :base(key1)
        {
            _= key2 ?? throw new ArgumentNullException(nameof(key2));
            _= key3 ?? throw new ArgumentNullException(nameof(key3));
            if(key2.Length != 8 || key3.Length != 8)
                throw new ArgumentOutOfRangeException("Invalid Key size. Key size must be 8.");

            var provider = new TripleDESCryptoServiceProvider();
            _cryptoTransform2 = provider.CreateDecryptor(key2,new byte[8]);
            _cryptoTransform3 = provider.CreateEncryptor(key3,new byte[8]);
        }

        public override ReadOnlySpan<byte> Compute(ReadOnlySpan<byte> data)
        {
            var result = base.Compute(data).ToArray();
            result = _cryptoTransform2.TransformFinalBlock(result, 0, 8);
            result = _cryptoTransform3.TransformFinalBlock(result, 0, 8);
            return result;
        }

        public override string GetName() => "Retail MAC";
        
        protected override void Dispose(bool disposing)
        {
            if (disposing)
            {
                _cryptoTransform2.Dispose();
                _cryptoTransform3.Dispose();
            }
            base.Dispose(false);
        }
    }
}