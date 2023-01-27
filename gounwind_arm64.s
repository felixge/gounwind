TEXT ·regfp(SB),$8-0
	MOVD (R29), R0
  MOVD R0, ret+0(FP)
	RET
