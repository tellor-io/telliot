typedef enum sha2_32_constants {
  // SHA-256 Initial Hash Values
  SHA256M_A=0x6a09e667,
  SHA256M_B=0xbb67ae85,
  SHA256M_C=0x3c6ef372,
  SHA256M_D=0xa54ff53a,
  SHA256M_E=0x510e527f,
  SHA256M_F=0x9b05688c,
  SHA256M_G=0x1f83d9ab,
  SHA256M_H=0x5be0cd19,

  // SHA-256 Constants
  SHA256C00=0x428a2f98,
  SHA256C01=0x71374491,
  SHA256C02=0xb5c0fbcf,
  SHA256C03=0xe9b5dba5,
  SHA256C04=0x3956c25b,
  SHA256C05=0x59f111f1,
  SHA256C06=0x923f82a4,
  SHA256C07=0xab1c5ed5,
  SHA256C08=0xd807aa98,
  SHA256C09=0x12835b01,
  SHA256C0a=0x243185be,
  SHA256C0b=0x550c7dc3,
  SHA256C0c=0x72be5d74,
  SHA256C0d=0x80deb1fe,
  SHA256C0e=0x9bdc06a7,
  SHA256C0f=0xc19bf174,
  SHA256C10=0xe49b69c1,
  SHA256C11=0xefbe4786,
  SHA256C12=0x0fc19dc6,
  SHA256C13=0x240ca1cc,
  SHA256C14=0x2de92c6f,
  SHA256C15=0x4a7484aa,
  SHA256C16=0x5cb0a9dc,
  SHA256C17=0x76f988da,
  SHA256C18=0x983e5152,
  SHA256C19=0xa831c66d,
  SHA256C1a=0xb00327c8,
  SHA256C1b=0xbf597fc7,
  SHA256C1c=0xc6e00bf3,
  SHA256C1d=0xd5a79147,
  SHA256C1e=0x06ca6351,
  SHA256C1f=0x14292967,
  SHA256C20=0x27b70a85,
  SHA256C21=0x2e1b2138,
  SHA256C22=0x4d2c6dfc,
  SHA256C23=0x53380d13,
  SHA256C24=0x650a7354,
  SHA256C25=0x766a0abb,
  SHA256C26=0x81c2c92e,
  SHA256C27=0x92722c85,
  SHA256C28=0xa2bfe8a1,
  SHA256C29=0xa81a664b,
  SHA256C2a=0xc24b8b70,
  SHA256C2b=0xc76c51a3,
  SHA256C2c=0xd192e819,
  SHA256C2d=0xd6990624,
  SHA256C2e=0xf40e3585,
  SHA256C2f=0x106aa070,
  SHA256C30=0x19a4c116,
  SHA256C31=0x1e376c08,
  SHA256C32=0x2748774c,
  SHA256C33=0x34b0bcb5,
  SHA256C34=0x391c0cb3,
  SHA256C35=0x4ed8aa4a,
  SHA256C36=0x5b9cca4f,
  SHA256C37=0x682e6ff3,
  SHA256C38=0x748f82ee,
  SHA256C39=0x78a5636f,
  SHA256C3a=0x84c87814,
  SHA256C3b=0x8cc70208,
  SHA256C3c=0x90befffa,
  SHA256C3d=0xa4506ceb,
  SHA256C3e=0xbef9a3f7,
  SHA256C3f=0xc67178f2,

} sha2_32_constants_t;

constant uint32_t k_sha256[64] =
{
  SHA256C00, SHA256C01, SHA256C02, SHA256C03,
  SHA256C04, SHA256C05, SHA256C06, SHA256C07,
  SHA256C08, SHA256C09, SHA256C0a, SHA256C0b,
  SHA256C0c, SHA256C0d, SHA256C0e, SHA256C0f,
  SHA256C10, SHA256C11, SHA256C12, SHA256C13,
  SHA256C14, SHA256C15, SHA256C16, SHA256C17,
  SHA256C18, SHA256C19, SHA256C1a, SHA256C1b,
  SHA256C1c, SHA256C1d, SHA256C1e, SHA256C1f,
  SHA256C20, SHA256C21, SHA256C22, SHA256C23,
  SHA256C24, SHA256C25, SHA256C26, SHA256C27,
  SHA256C28, SHA256C29, SHA256C2a, SHA256C2b,
  SHA256C2c, SHA256C2d, SHA256C2e, SHA256C2f,
  SHA256C30, SHA256C31, SHA256C32, SHA256C33,
  SHA256C34, SHA256C35, SHA256C36, SHA256C37,
  SHA256C38, SHA256C39, SHA256C3a, SHA256C3b,
  SHA256C3c, SHA256C3d, SHA256C3e, SHA256C3f,
};


#define SHIFT_RIGHT_32(x,n) ((x) >> (n))

#define SHA256_S0_S(x) (rotate ((x), 25u) ^ rotate ((x), 14u) ^ SHIFT_RIGHT_32 ((x),  3u))
#define SHA256_S1_S(x) (rotate ((x), 15u) ^ rotate ((x), 13u) ^ SHIFT_RIGHT_32 ((x), 10u))
#define SHA256_S2_S(x) (rotate ((x), 30u) ^ rotate ((x), 19u) ^ rotate ((x), 10u))
#define SHA256_S3_S(x) (rotate ((x), 26u) ^ rotate ((x), 21u) ^ rotate ((x),  7u))

#define SHA256_S0(x) (rotate ((x), 25u) ^ rotate ((x), 14u) ^ SHIFT_RIGHT_32 ((x),  3u))
#define SHA256_S1(x) (rotate ((x), 15u) ^ rotate ((x), 13u) ^ SHIFT_RIGHT_32 ((x), 10u))
#define SHA256_S2(x) (rotate ((x), 30u) ^ rotate ((x), 19u) ^ rotate ((x), 10u))
#define SHA256_S3(x) (rotate ((x), 26u) ^ rotate ((x), 21u) ^ rotate ((x),  7u))

#define SHA256_F0(x,y,z)  (((x) & (y)) | ((z) & ((x) ^ (y))))
#define SHA256_F1(x,y,z)  ((z) ^ ((x) & ((y) ^ (z))))
#define SHA256_F0o(x,y,z) (bitselect ((x), (y), ((x) ^ (z))))
#define SHA256_F1o(x,y,z) (bitselect ((z), (y), (x)))

#define SHA256_STEP_S(F0,F1,a,b,c,d,e,f,g,h,x,K)  \
{                                                 \
  h = hc_add3_S (h, K, x);                        \
  h = hc_add3_S (h, SHA256_S3_S (e), F1 (e,f,g)); \
  d += h;                                         \
  h = hc_add3_S (h, SHA256_S2_S (a), F0 (a,b,c)); \
}

#define SHA256_EXPAND_S(x,y,z,w) (SHA256_S1_S (x) + y + SHA256_S0_S (z) + w)

#define SHA256_STEP(F0,F1,a,b,c,d,e,f,g,h,x,K)    \
{                                                 \
  h = hc_add3 (h, K, x);                          \
  h = hc_add3 (h, SHA256_S3 (e), F1 (e,f,g));     \
  d += h;                                         \
  h = hc_add3 (h, SHA256_S2 (a), F0 (a,b,c));     \
}

#define SHA256_EXPAND(x,y,z,w) (SHA256_S1 (x) + y + SHA256_S0 (z) + w)

#define EndianSwap(n) (rotate(n & 0x00FF00FF, 24U)|(rotate(n, 8U) & 0x00FF00FF))

uint32_t hc_add3 (const uint32_t a, const uint32_t b, const uint32_t c)
{
  return a + b + c;
}

//DECLSPEC void sha256_transform_vector (const uint32_t *w0, const uint32_t *w1, const uint32_t *w2, const uint32_t *w3, uint32_t *digest)

void sha2_fast(uint32_t *w0, uint32_t *digest) {

    digest[0] = SHA256M_A;
    digest[1] = SHA256M_B;
    digest[2] = SHA256M_C;
    digest[3] = SHA256M_D;
    digest[4] = SHA256M_E;
    digest[5] = SHA256M_F;
    digest[6] = SHA256M_G;
    digest[7] = SHA256M_H;

    uint32_t a = digest[0];
    uint32_t b = digest[1];
    uint32_t c = digest[2];
    uint32_t d = digest[3];
    uint32_t e = digest[4];
    uint32_t f = digest[5];
    uint32_t g = digest[6];
    uint32_t h = digest[7];

    uint32_t w0_t = EndianSwap(w0[0]);
    uint32_t w1_t = EndianSwap(w0[1]);
    uint32_t w2_t = EndianSwap(w0[2]);
    uint32_t w3_t = EndianSwap(w0[3]);
    uint32_t w4_t = EndianSwap(w0[4]);
    uint32_t w5_t = EndianSwap((uint32_t)(0x80));
    uint32_t w6_t = 0;
    uint32_t w7_t = 0;
    uint32_t w8_t = 0;
    uint32_t w9_t = 0;
    uint32_t wa_t = 0;
    uint32_t wb_t = 0;
    uint32_t wc_t = 0;
    uint32_t wd_t = 0;
    uint32_t we_t = 0;
    uint32_t wf_t = 20*8;


    #define ROUND_EXPAND()                            \
    {                                                 \
    w0_t = SHA256_EXPAND (we_t, w9_t, w1_t, w0_t);  \
    w1_t = SHA256_EXPAND (wf_t, wa_t, w2_t, w1_t);  \
    w2_t = SHA256_EXPAND (w0_t, wb_t, w3_t, w2_t);  \
    w3_t = SHA256_EXPAND (w1_t, wc_t, w4_t, w3_t);  \
    w4_t = SHA256_EXPAND (w2_t, wd_t, w5_t, w4_t);  \
    w5_t = SHA256_EXPAND (w3_t, we_t, w6_t, w5_t);  \
    w6_t = SHA256_EXPAND (w4_t, wf_t, w7_t, w6_t);  \
    w7_t = SHA256_EXPAND (w5_t, w0_t, w8_t, w7_t);  \
    w8_t = SHA256_EXPAND (w6_t, w1_t, w9_t, w8_t);  \
    w9_t = SHA256_EXPAND (w7_t, w2_t, wa_t, w9_t);  \
    wa_t = SHA256_EXPAND (w8_t, w3_t, wb_t, wa_t);  \
    wb_t = SHA256_EXPAND (w9_t, w4_t, wc_t, wb_t);  \
    wc_t = SHA256_EXPAND (wa_t, w5_t, wd_t, wc_t);  \
    wd_t = SHA256_EXPAND (wb_t, w6_t, we_t, wd_t);  \
    we_t = SHA256_EXPAND (wc_t, w7_t, wf_t, we_t);  \
    wf_t = SHA256_EXPAND (wd_t, w8_t, w0_t, wf_t);  \
    }

    #define ROUND_STEP(i)                                                                   \
    {                                                                                       \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, a, b, c, d, e, f, g, h, w0_t, k_sha256[i +  0]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, h, a, b, c, d, e, f, g, w1_t, k_sha256[i +  1]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, g, h, a, b, c, d, e, f, w2_t, k_sha256[i +  2]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, f, g, h, a, b, c, d, e, w3_t, k_sha256[i +  3]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, e, f, g, h, a, b, c, d, w4_t, k_sha256[i +  4]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, d, e, f, g, h, a, b, c, w5_t, k_sha256[i +  5]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, c, d, e, f, g, h, a, b, w6_t, k_sha256[i +  6]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, b, c, d, e, f, g, h, a, w7_t, k_sha256[i +  7]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, a, b, c, d, e, f, g, h, w8_t, k_sha256[i +  8]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, h, a, b, c, d, e, f, g, w9_t, k_sha256[i +  9]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, g, h, a, b, c, d, e, f, wa_t, k_sha256[i + 10]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, f, g, h, a, b, c, d, e, wb_t, k_sha256[i + 11]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, e, f, g, h, a, b, c, d, wc_t, k_sha256[i + 12]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, d, e, f, g, h, a, b, c, wd_t, k_sha256[i + 13]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, c, d, e, f, g, h, a, b, we_t, k_sha256[i + 14]); \
    SHA256_STEP (SHA256_F0o, SHA256_F1o, b, c, d, e, f, g, h, a, wf_t, k_sha256[i + 15]); \
    }

    ROUND_STEP (0);

    #ifdef _unroll
    #pragma unroll
    #endif
    for (int i = 16; i < 64; i += 16) {
        ROUND_EXPAND (); ROUND_STEP (i);
    }

    #undef ROUND_EXPAND
    #undef ROUND_STEP

    digest[0] += a;
    digest[1] += b;
    digest[2] += c;
    digest[3] += d;
    digest[4] += e;
    digest[5] += f;
    digest[6] += g;
    digest[7] += h;
}
