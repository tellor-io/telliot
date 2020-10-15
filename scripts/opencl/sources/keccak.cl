
#ifndef KECCAK_ROUNDS
#define KECCAK_ROUNDS 24
#endif

constant uint64_t keccakf_rndc[24] =
{
  0x0000000000000001, 0x0000000000008082, 0x800000000000808a,
  0x8000000080008000, 0x000000000000808b, 0x0000000080000001,
  0x8000000080008081, 0x8000000000008009, 0x000000000000008a,
  0x0000000000000088, 0x0000000080008009, 0x000000008000000a,
  0x000000008000808b, 0x800000000000008b, 0x8000000000008089,
  0x8000000000008003, 0x8000000000008002, 0x8000000000000080,
  0x000000000000800a, 0x800000008000000a, 0x8000000080008081,
  0x8000000000008080, 0x0000000080000001, 0x8000000080008008
};


uint64_t rotl64(const uint64_t a, const int n) {
  return ((a << n) | ((a >> (64 - n))));
}


void keccak(constant uint64_t *prefix, uint32_t *nonce, uint64_t *output) {
    /**
     * Keccak
     */

    uint64_t a00 = prefix[0];
    uint64_t a01 = prefix[1];
    uint64_t a02 = prefix[2];
    uint64_t a03 = prefix[3];
    uint64_t a04 = prefix[4];
    uint64_t a10 = prefix[5];

    uint64_t a11, a12, a13, a14;

    if (prefix[7] != NULL) {
      //the prefix is 60 bytes long when enablePoolWorker
      a11 = prefix[6];

      uint32_t *a12s = (uint32_t*)&a12;
      a12s[0] = prefix[7];
      a12s[1] = nonce[0];

      uint32_t *a13s = (uint32_t*)&a13;
      a13s[0] = nonce[1];
      a13s[1] = nonce[2];

      uint32_t *a14s = (uint32_t*)&a14;
      a14s[0] = nonce[3];
      a14s[1] = 1; //bit padding for keccak (first bit after message ends = 1)
    } else {
      //the prefix is only 52 bytes long - the last 4 bytes in a11 are the first 4 bytes of the nonce
      uint32_t *a11s = (uint32_t*)&a11;
      a11s[0] = prefix[6];
      a11s[1] = nonce[0];

      uint32_t *a12s = (uint32_t*)&a12;
      a12s[0] = nonce[1];
      a12s[1] = nonce[2];

      uint32_t *a13s = (uint32_t*)&a13;
      a13s[0] = nonce[3];
      a13s[1] = 1; //bit padding for keccak (first bit after message ends = 1)

      a14 = 0;
    }
    uint64_t a20 = 0;
    uint64_t a21 = 0;
    uint64_t a22 = 0;
    uint64_t a23 = 0;
    uint64_t a24 = 0;
    uint64_t a30 = 0;
    uint64_t a31 = 0x8000000000000000; //end of block bit padding (last bit in block == 1)
    uint64_t a32 = 0;
    uint64_t a33 = 0;
    uint64_t a34 = 0;
    uint64_t a40 = 0;
    uint64_t a41 = 0;
    uint64_t a42 = 0;
    uint64_t a43 = 0;
    uint64_t a44 = 0;

    #define Rho_Pi(ad,r)     \
      bc0 = ad;              \
      ad = rotl64 (t, r);    \
      t = bc0;               \

    #ifdef _unroll
    #pragma unroll
    #endif
    for (int round = 0; round < KECCAK_ROUNDS; round++)
    {
      // Theta

      uint64_t bc0 = a00 ^ a10 ^ a20 ^ a30 ^ a40;
      uint64_t bc1 = a01 ^ a11 ^ a21 ^ a31 ^ a41;
      uint64_t bc2 = a02 ^ a12 ^ a22 ^ a32 ^ a42;
      uint64_t bc3 = a03 ^ a13 ^ a23 ^ a33 ^ a43;
      uint64_t bc4 = a04 ^ a14 ^ a24 ^ a34 ^ a44;

      uint64_t t;

      t = bc4 ^ rotl64 (bc1, 1); a00 ^= t; a10 ^= t; a20 ^= t; a30 ^= t; a40 ^= t;
      t = bc0 ^ rotl64 (bc2, 1); a01 ^= t; a11 ^= t; a21 ^= t; a31 ^= t; a41 ^= t;
      t = bc1 ^ rotl64 (bc3, 1); a02 ^= t; a12 ^= t; a22 ^= t; a32 ^= t; a42 ^= t;
      t = bc2 ^ rotl64 (bc4, 1); a03 ^= t; a13 ^= t; a23 ^= t; a33 ^= t; a43 ^= t;
      t = bc3 ^ rotl64 (bc0, 1); a04 ^= t; a14 ^= t; a24 ^= t; a34 ^= t; a44 ^= t;

      // Rho Pi

      t = a01;

      Rho_Pi (a20,  1);
      Rho_Pi (a12,  3);
      Rho_Pi (a21,  6);
      Rho_Pi (a32, 10);
      Rho_Pi (a33, 15);
      Rho_Pi (a03, 21);
      Rho_Pi (a10, 28);
      Rho_Pi (a31, 36);
      Rho_Pi (a13, 45);
      Rho_Pi (a41, 55);
      Rho_Pi (a44,  2);
      Rho_Pi (a04, 14);
      Rho_Pi (a30, 27);
      Rho_Pi (a43, 41);
      Rho_Pi (a34, 56);
      Rho_Pi (a23,  8);
      Rho_Pi (a22, 25);
      Rho_Pi (a02, 43);
      Rho_Pi (a40, 62);
      Rho_Pi (a24, 18);
      Rho_Pi (a42, 39);
      Rho_Pi (a14, 61);
      Rho_Pi (a11, 20);
      Rho_Pi (a01, 44);

      //  Chi

      bc0 = a00; bc1 = a01; bc2 = a02; bc3 = a03; bc4 = a04;
      a00 ^= ~bc1 & bc2; a01 ^= ~bc2 & bc3; a02 ^= ~bc3 & bc4; a03 ^= ~bc4 & bc0; a04 ^= ~bc0 & bc1;

      bc0 = a10; bc1 = a11; bc2 = a12; bc3 = a13; bc4 = a14;
      a10 ^= ~bc1 & bc2; a11 ^= ~bc2 & bc3; a12 ^= ~bc3 & bc4; a13 ^= ~bc4 & bc0; a14 ^= ~bc0 & bc1;

      bc0 = a20; bc1 = a21; bc2 = a22; bc3 = a23; bc4 = a24;
      a20 ^= ~bc1 & bc2; a21 ^= ~bc2 & bc3; a22 ^= ~bc3 & bc4; a23 ^= ~bc4 & bc0; a24 ^= ~bc0 & bc1;

      bc0 = a30; bc1 = a31; bc2 = a32; bc3 = a33; bc4 = a34;
      a30 ^= ~bc1 & bc2; a31 ^= ~bc2 & bc3; a32 ^= ~bc3 & bc4; a33 ^= ~bc4 & bc0; a34 ^= ~bc0 & bc1;

      bc0 = a40; bc1 = a41; bc2 = a42; bc3 = a43; bc4 = a44;
      a40 ^= ~bc1 & bc2; a41 ^= ~bc2 & bc3; a42 ^= ~bc3 & bc4; a43 ^= ~bc4 & bc0; a44 ^= ~bc0 & bc1;

      //  Iota

      a00 ^= keccakf_rndc[round];
    }

    #undef Rho_Pi

    output[0] = a00;
    output[1] = a01;
    output[2] = a02;
    output[3] = a03;


}
