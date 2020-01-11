

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
        uint8_t ripe160Result[20];
        ripemd160_transform_vector(hashResult, ripe160Result);
        sha2_fast(ripe160Result,hashResult);


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

