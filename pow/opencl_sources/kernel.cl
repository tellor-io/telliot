
typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned int uint32_t;
typedef unsigned long uint64_t;
typedef unsigned long long uint128_t;


// NOTE - if you are editing the opencl code, and change these includes *WITHOUT* changing this file itself
// cuda's terrible caching will give you the OLD version (lol)
// if you are developing, you can delete ~/.nv/ComputeCache to reset the cache
// see https://stackoverflow.com/questions/31338520/
#include "memstuff.c"
#include "ripemd160.c"
#include "sha2.c"
#include "mul32.c"
#include "keccak.cl"

constant char hex[] = "0123456789abcdef";


__kernel void tellor(
    //challenge + public address
   constant uint64_t *prefix,

   //pre-computed mulitiplier for remainder test
   // two 512 bit constants
   constant uint32_t *mulDivisor,

   //found nonce saved here
   __global uint32_t *output,

   //start index for nonces
   uint64_t base,

   //number of loops to run
   uint32_t count
)
{
    uint32_t gid = get_global_id(0);
    if(gid < 4) {
        output[gid] = 0;
    }

    uint8_t hashResult[32];
    uint8_t nonce[16];
    for(int i = 0; i < count; i++) {
        uint64_t index = base + gid*count + i;

        //create a 16 character hex string from an 64 bit nonce number
        uint8_t *dat8 = (uint8_t*)&index;
        for(int i = 0; i < 8; i++) {
            nonce[i*2 + 0] = hex[dat8[i] >> 4];
            nonce[i*2 + 1] = hex[dat8[i] & 0xf];
        }

        //run the tellor hash algo
        keccak(prefix, (uint32_t*)nonce, (uint64_t*)hashResult);
        ripemd160(hashResult, sizeof(hashResult), hashResult);
        sha256_Raw(hashResult, 20,hashResult);

        //the divisibility test was originally implemented using big.Int in the golang std lib.
        //big.Int.SetBytes() treats the input as big endian, and flips it before performing
        //internal calculations. Thus, we must also flip it here to match
        //this could be removed by re-writing the divisible() function and computing different
        //mulDivisor constants, but the perf cost is so low (<1%) that it's probably not worth the effort
        for(int i = 0; i < 16; i++) {
            uint8_t tmp = hashResult[i];
            hashResult[i] = hashResult[31-i];
            hashResult[31-i] = tmp;
        }

        //test if result is divisible by target difficulty
        //works by multiplying by a constant and then checking against another constant
        //see https://lemire.me/blog/2019/02/08/faster-remainders-when-the-divisor-is-a-constant-beating-compilers-and-libdivide/
        uint32_t *b = (uint32_t*)&hashResult;
        int result = divisible(b, &mulDivisor[0], &mulDivisor[16]);
        if (result != 0) {
            output[0] = ((uint32_t*)nonce)[0];
            output[1] = ((uint32_t*)nonce)[1];
            output[2] = ((uint32_t*)nonce)[2];
            output[3] = ((uint32_t*)nonce)[3];
            return;
        }
    }
}

