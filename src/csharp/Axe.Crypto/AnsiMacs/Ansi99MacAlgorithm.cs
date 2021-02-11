using System;
using System.Security.Cryptography;

namespace Axe.Crypto.AnsiMacs
{
    public class Ansi99MacAlgorithm : IAnsiMacAlgorithm
    {
        private readonly ICryptoTransform _cryptoTransform;
        private bool disposedValue;

        public Ansi99MacAlgorithm(byte[] key)
        {
            _= key ?? throw new ArgumentNullException(nameof(key));
            if(key.Length != 8)
                throw new ArgumentOutOfRangeException("Invalid Key size. Key size must be 8.");

            var provider = new TripleDESCryptoServiceProvider();
            _cryptoTransform = provider.CreateEncryptor(key,new byte[8]);
        }

        public virtual ReadOnlySpan<byte> Compute(ReadOnlySpan<byte> data)
        {
            var result = new byte[8];
            var blockSize = 8;
            var iteration = Math.Ceiling((double)data.Length/blockSize);

            for (int i = 0; i < iteration; i++)
            {
                var start = i * blockSize;
                var end = start + blockSize;
                var chunk = data.Slice(start, end);
                CryptoUtilities.Xor(result, chunk);
                result = _cryptoTransform.TransformFinalBlock(result, 0, 8);
            }

            return new ReadOnlySpan<byte>(result);
        }

        public virtual string GetName() => "Ansi 9.9 MAC";
        
        protected virtual void Dispose(bool disposing)
        {
            if (!disposedValue)
            {
                if (disposing)
                {
                    _cryptoTransform.Dispose();
                }
                disposedValue = true;
            }
        }

        public void Dispose()
        {
            Dispose(disposing: true);
            GC.SuppressFinalize(this);
        }
    }
}