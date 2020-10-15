typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned int uint32_t;
typedef unsigned long uint64_t;

#define EndianSwap(n) (rotate(n & 0x00FF00FF, 24U)|(rotate(n, 8U) & 0x00FF00FF))


constant char hex[] = "0123456789abcdef";
