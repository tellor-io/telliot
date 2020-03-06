

#define _W  (32U) // word size in bits
#define _M (0xffffffffU)        // digit mask

#define _W2 (16U)   // half word size in bits
#define _B2 (1U << _W2) // half digit base
#define _M2 (_B2 - 1U)  // half digit mask


// z1<<_W + z0 = x*y + c
uint32_t mulAddWWW(uint32_t x,uint32_t y, uint32_t c, uint32_t *z0) {
	uint32_t zz0 = x * y;
	*z0 = zz0 + c;
	uint32_t c1 = !!(*z0 < zz0);
	return mad_hi(x, y, c1);
}

uint32_t addWW(uint32_t x, uint32_t y, uint32_t *z0) {
	*z0 = x + y;
	return (*z0 < x) ? 1 : 0;
}

uint32_t addMulVVW(uint32_t *z, uint32_t n, constant uint32_t *x, uint32_t y) {
	uint32_t c = 0;
	for (int i = 0; i < n; i++) {
		uint32_t z0;
		uint32_t z1 = mulAddWWW(x[i], y, z[i], &z0);
		c = addWW(z0, c, &z[i]);
		c += z1;
	}
	return c;
}

uint32_t addMulVVWZero(uint32_t *z, uint32_t n, constant uint32_t *x, uint32_t y) {
	uint32_t c = 0;
	for (int i = 0; i < n; i++) {
		uint32_t z0;
		uint32_t z1 = mulAddWWW(x[i], y, 0, &z0);
		c = addWW(z0, c, &z[i]);
		c += z1;
	}
	return c;
}


//p has len MUL_CONSTANT_SIZE words
//x has len MUL_CONSTANT_SIZE words
//y has len 256 bits (8 words)
//p = x * y
void basicMul(uint32_t *p, constant uint32_t *x, uint32_t *y) {
    uint32_t d = y[7];
    addMulVVWZero(&p[0], MUL_CONSTANT_SIZE-0, x, d);

	for (int i = 1; i < 8; i++) {
		uint32_t d = y[7-i];
        addMulVVW(&p[i], MUL_CONSTANT_SIZE-i, x, d);
	}
}

int divisible(uint32_t *x, constant uint32_t *c, constant uint32_t *cSub1) {
	uint32_t p[16];
	basicMul(p, c, x);
	int i = MUL_CONSTANT_SIZE-1;
	while(i > 0 && p[i] == cSub1[i]) {
		i--;
	}
	return p[i] <= cSub1[i];
}

