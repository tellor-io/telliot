

#define _W  (32U) // word size in bits
#define _M (0xffffffffU)        // digit mask

#define _W2 (16U)   // half word size in bits
#define _B2 (1U << _W2) // half digit base
#define _M2 (_B2 - 1U)  // half digit mask


// z1<<_W + z0 = x*y
// Adapted from Warren, Hacker's Delight, p. 132.
uint32_t mulWW(uint32_t x, uint32_t y, uint32_t *z0) {
	uint32_t x0 = x & _M2;
	uint32_t x1 = x >> _W2;
	uint32_t y0 = y & _M2;
	uint32_t y1 = y >> _W2;
	uint32_t w0 = x0 * y0;
	uint32_t t = x1*y0 + (w0>>_W2);
	uint32_t w1 = t & _M2;
	uint32_t w2 = t >> _W2;
	w1 += x0 * y1;
	*z0 = x * y;
	return x1*y1 + w2 + (w1>>_W2);
}

// z1<<_W + z0 = x*y + c
uint32_t mulAddWWW(uint32_t x,uint32_t y, uint32_t c, uint32_t *z0) {
	uint32_t zz0;
	uint32_t z1 = mulWW(x, y, &zz0);
	*z0 = zz0 + c;
	if (*z0 < zz0) {
		z1++;
	}
	return z1;
}

uint32_t addWW(uint32_t x, uint32_t y, uint32_t c, uint32_t *z0) {
	uint32_t z1 = 0;
	uint32_t yc = y + c;
	*z0 = x + yc;
	if (*z0 < x || yc < y) {
		z1 = 1;
	}
	return z1;
}


uint32_t addMulVVW(uint32_t *z, uint32_t n, constant uint32_t *x, uint32_t y) {
	uint32_t c = 0;
	for (int i = 0; i < n; i++) {
		uint32_t z0;
		uint32_t z1 = mulAddWWW(x[i], y, z[i], &z0);
		c = addWW(z0, c, 0, &z[i]);
		c += z1;
	}
	return c;
}

uint32_t addMulVVWZero(uint32_t *z, uint32_t n, constant uint32_t *x, uint32_t y) {
	uint32_t c = 0;
	for (int i = 0; i < n; i++) {
		uint32_t z0;
		uint32_t z1 = mulAddWWW(x[i], y, 0, &z0);
		c = addWW(z0, c, 0, &z[i]);
		c += z1;
	}
	return c;
}

//p has len 512 bits
//x has len 512 bits (16 words)
//y has len 256 bits (8 words)
//p = x * y
void basicMul(uint32_t *p, constant uint32_t *x, uint32_t *y) {
    uint32_t d = y[7];
    addMulVVWZero(&p[0], 16-0, x, d);

	for (int i = 1; i < 8; i++) {
		uint32_t d = y[7-i];
        addMulVVW(&p[i], 16-i, x, d);
	}
}

int divisible(uint32_t *x, constant uint32_t *c, constant uint32_t *cSub1) {
	uint32_t p[16];
	basicMul(p, c, x);
	int i = 15;
	while(i > 0 && p[i] == cSub1[i]) {
		i--;
	}
	return p[i] <= cSub1[i];
}

