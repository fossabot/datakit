// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.11
// +build !gccgo,!appengine

#include "textflag.h"

#define NUM_ROUNDS 10

// func xorKeyStreamVX(dst, src []byte, key *[8]uint32, nonce *[3]uint32, counter *uint32)
TEXT ·xorKeyStreamVX(SB), NOSPLIT, $0
	MOVD	dst+0(FP), R1
	MOVD	src+24(FP), R2
	MOVD	src_len+32(FP), R3
	MOVD	key+48(FP), R4
	MOVD	nonce+56(FP), R6
	MOVD	counter+64(FP), R7

	MOVD	$·constants(SB), R10
	MOVD	$·incRotMatrix(SB), R11

	MOVW	(R7), R20

	AND	$~255, R3, R13
	ADD	R2, R13, R12 // R12 for block end
	AND	$255, R3, R13
loop:
	MOVD	$NUM_ROUNDS, R21
	VLD1	(R11), [V30.S4, V31.S4]

	// load contants
	// VLD4R (R10), [V0.S4, V1.S4, V2.S4, V3.S4]
	WORD	$0x4D60E940

	// load keys
	// VLD4R 16(R4), [V4.S4, V5.S4, V6.S4, V7.S4]
	WORD	$0x4DFFE884
	// VLD4R 16(R4), [V8.S4, V9.S4, V10.S4, V11.S4]
	WORD	$0x4DFFE888
	SUB	$32, R4

	// load counter + nonce
	// VLD1R (R7), [V12.S4]
	WORD	$0x4D40C8EC

	// VLD3R (R6), [V13.S4, V14.S4, V15.S4]
	WORD	$0x4D40E8CD

	// update counter
	VADD	V30.S4, V12.S4, V12.S4

chacha:
	// V0..V3 += V4..V7
	// V12..V15 <<<= ((V12..V15 XOR V0..V3), 16)
	VADD	V0.S4, V4.S4, V0.S4
	VADD	V1.S4, V5.S4, V1.S4
	VADD	V2.S4, V6.S4, V2.S4
	VADD	V3.S4, V7.S4, V3.S4
	VEOR	V12.B16, V0.B16, V12.B16
	VEOR	V13.B16, V1.B16, V13.B16
	VEOR	V14.B16, V2.B16, V14.B16
	VEOR	V15.B16, V3.B16, V15.B16
	VREV32	V12.H8, V12.H8
	VREV32	V13.H8, V13.H8
	VREV32	V14.H8, V14.H8
	VREV32	V15.H8, V15.H8
	// V8..V11 += V12..V15
	// V4..V7 <<<= ((V4..V7 XOR V8..V11), 12)
	VADD	V8.S4, V12.S4, V8.S4
	VADD	V9.S4, V13.S4, V9.S4
	VADD	V10.S4, V14.S4, V10.S4
	VADD	V11.S4, V15.S4, V11.S4
	VEOR	V8.B16, V4.B16, V16.B16
	VEOR	V9.B16, V5.B16, V17.B16
	VEOR	V10.B16, V6.B16, V18.B16
	VEOR	V11.B16, V7.B16, V19.B16
	VSHL	$12, V16.S4, V4.S4
	VSHL	$12, V17.S4, V5.S4
	VSHL	$12, V18.S4, V6.S4
	VSHL	$12, V19.S4, V7.S4
	VSRI	$20, V16.S4, V4.S4
	VSRI	$20, V17.S4, V5.S4
	VSRI	$20, V18.S4, V6.S4
	VSRI	$20, V19.S4, V7.S4

	// V0..V3 += V4..V7
	// V12..V15 <<<= ((V12..V15 XOR V0..V3), 8)
	VADD	V0.S4, V4.S4, V0.S4
	VADD	V1.S4, V5.S4, V1.S4
	VADD	V2.S4, V6.S4, V2.S4
	VADD	V3.S4, V7.S4, V3.S4
	VEOR	V12.B16, V0.B16, V12.B16
	VEOR	V13.B16, V1.B16, V13.B16
	VEOR	V14.B16, V2.B16, V14.B16
	VEOR	V15.B16, V3.B16, V15.B16
	VTBL	V31.B16, [V12.B16], V12.B16
	VTBL	V31.B16, [V13.B16], V13.B16
	VTBL	V31.B16, [V14.B16], V14.B16
	VTBL	V31.B16, [V15.B16], V15.B16

	// V8..V11 += V12..V15
	// V4..V7 <<<= ((V4..V7 XOR V8..V11), 7)
	VADD	V12.S4, V8.S4, V8.S4
	VADD	V13.S4, V9.S4, V9.S4
	VADD	V14.S4, V10.S4, V10.S4
	VADD	V15.S4, V11.S4, V11.S4
	VEOR	V8.B16, V4.B16, V16.B16
	VEOR	V9.B16, V5.B16, V17.B16
	VEOR	V10.B16, V6.B16, V18.B16
	VEOR	V11.B16, V7.B16, V19.B16
	VSHL	$7, V16.S4, V4.S4
	VSHL	$7, V17.S4, V5.S4
	VSHL	$7, V18.S4, V6.S4
	VSHL	$7, V19.S4, V7.S4
	VSRI	$25, V16.S4, V4.S4
	VSRI	$25, V17.S4, V5.S4
	VSRI	$25, V18.S4, V6.S4
	VSRI	$25, V19.S4, V7.S4

	// V0..V3 += V5..V7, V4
	// V15,V12-V14 <<<= ((V15,V12-V14 XOR V0..V3), 16)
	VADD	V0.S4, V5.S4, V0.S4
	VADD	V1.S4, V6.S4, V1.S4
	VADD	V2.S4, V7.S4, V2.S4
	VADD	V3.S4, V4.S4, V3.S4
	VEOR	V15.B16, V0.B16, V15.B16
	VEOR	V12.B16, V1.B16, V12.B16
	VEOR	V13.B16, V2.B16, V13.B16
	VEOR	V14.B16, V3.B16, V14.B16
	VREV32	V12.H8, V12.H8
	VREV32	V13.H8, V13.H8
	VREV32	V14.H8, V14.H8
	VREV32	V15.H8, V15.H8

	// V10 += V15; V5 <<<= ((V10 XOR V5), 12)
	// ...
	VADD	V15.S4, V10.S4, V10.S4
	VADD	V12.S4, V11.S4, V11.S4
	VADD	V13.S4, V8.S4, V8.S4
	VADD	V14.S4, V9.S4, V9.S4
	VEOR	V10.B16, V5.B16, V16.B16
	VEOR	V11.B16, V6.B16, V17.B16
	VEOR	V8.B16, V7.B16, V18.B16
	VEOR	V9.B16, V4.B16, V19.B16
	VSHL	$12, V16.S4, V5.S4
	VSHL	$12, V17.S4, V6.S4
	VSHL	$12, V18.S4, V7.S4
	VSHL	$12, V19.S4, V4.S4
	VSRI	$20, V16.S4, V5.S4
	VSRI	$20, V17.S4, V6.S4
	VSRI	$20, V18.S4, V7.S4
	VSRI	$20, V19.S4, V4.S4

	// V0 += V5; V15 <<<= ((V0 XOR V15), 8)
	// ...
	VADD	V5.S4, V0.S4, V0.S4
	VADD	V6.S4, V1.S4, V1.S4
	VADD	V7.S4, V2.S4, V2.S4
	VADD	V4.S4, V3.S4, V3.S4
	VEOR	V0.B16, V15.B16, V15.B16
	VEOR	V1.B16, V12.B16, V12.B16
	VEOR	V2.B16, V13.B16, V13.B16
	VEOR	V3.B16, V14.B16, V14.B16
	VTBL	V31.B16, [V12.B16], V12.B16
	VTBL	V31.B16, [V13.B16], V13.B16
	VTBL	V31.B16, [V14.B16], V14.B16
	VTBL	V31.B16, [V15.B16], V15.B16

	// V10 += V15; V5 <<<= ((V10 XOR V5), 7)
	// ...
	VADD	V15.S4, V10.S4, V10.S4
	VADD	V12.S4, V11.S4, V11.S4
	VADD	V13.S4, V8.S4, V8.S4
	VADD	V14.S4, V9.S4, V9.S4
	VEOR	V10.B16, V5.B16, V16.B16
	VEOR	V11.B16, V6.B16, V17.B16
	VEOR	V8.B16, V7.B16, V18.B16
	VEOR	V9.B16, V4.B16, V19.B16
	VSHL	$7, V16.S4, V5.S4
	VSHL	$7, V17.S4, V6.S4
	VSHL	$7, V18.S4, V7.S4
	VSHL	$7, V19.S4, V4.S4
	VSRI	$25, V16.S4, V5.S4
	VSRI	$25, V17.S4, V6.S4
	VSRI	$25, V18.S4, V7.S4
	VSRI	$25, V19.S4, V4.S4

	SUB	$1, R21
	CBNZ	R21, chacha

	// VLD4R (R10), [V16.S4, V17.S4, V18.S4, V19.S4]
	WORD	$0x4D60E950

	// VLD4R 16(R4), [V20.S4, V21.S4, V22.S4, V23.S4]
	WORD	$0x4DFFE894
	VADD	V30.S4, V12.S4, V12.S4
	VADD	V16.S4, V0.S4, V0.S4
	VADD	V17.S4, V1.S4, V1.S4
	VADD	V18.S4, V2.S4, V2.S4
	VADD	V19.S4, V3.S4, V3.S4
	// VLD4R 16(R4), [V24.S4, V25.S4, V26.S4, V27.S4]
	WORD	$0x4DFFE898
	// restore R4
	SUB	$32, R4

	// load counter + nonce
	// VLD1R (R7), [V28.S4]
	WORD	$0x4D40C8FC
	// VLD3R (R6), [V29.S4, V30.S4, V31.S4]
	WORD	$0x4D40E8DD

	VADD	V20.S4, V4.S4, V4.S4
	VADD	V21.S4, V5.S4, V5.S4
	VADD	V22.S4, V6.S4, V6.S4
	VADD	V23.S4, V7.S4, V7.S4
	VADD	V24.S4, V8.S4, V8.S4
	VADD	V25.S4, V9.S4, V9.S4
	VADD	V26.S4, V10.S4, V10.S4
	VADD	V27.S4, V11.S4, V11.S4
	VADD	V28.S4, V12.S4, V12.S4
	VADD	V29.S4, V13.S4, V13.S4
	VADD	V30.S4, V14.S4, V14.S4
	VADD	V31.S4, V15.S4, V15.S4

	VZIP1	V1.S4, V0.S4, V16.S4
	VZIP2	V1.S4, V0.S4, V17.S4
	VZIP1	V3.S4, V2.S4, V18.S4
	VZIP2	V3.S4, V2.S4, V19.S4
	VZIP1	V5.S4, V4.S4, V20.S4
	VZIP2	V5.S4, V4.S4, V21.S4
	VZIP1	V7.S4, V6.S4, V22.S4
	VZIP2	V7.S4, V6.S4, V23.S4
	VZIP1	V9.S4, V8.S4, V24.S4
	VZIP2	V9.S4, V8.S4, V25.S4
	VZIP1	V11.S4, V10.S4, V26.S4
	VZIP2	V11.S4, V10.S4, V27.S4
	VZIP1	V13.S4, V12.S4, V28.S4
	VZIP2	V13.S4, V12.S4, V29.S4
	VZIP1	V15.S4, V14.S4, V30.S4
	VZIP2	V15.S4, V14.S4, V31.S4
	VZIP1	V18.D2, V16.D2, V0.D2
	VZIP2	V18.D2, V16.D2, V4.D2
	VZIP1	V19.D2, V17.D2, V8.D2
	VZIP2	V19.D2, V17.D2, V12.D2
	VLD1.P	64(R2), [V16.B16, V17.B16, V18.B16, V19.B16]

	VZIP1	V22.D2, V20.D2, V1.D2
	VZIP2	V22.D2, V20.D2, V5.D2
	VZIP1	V23.D2, V21.D2, V9.D2
	VZIP2	V23.D2, V21.D2, V13.D2
	VLD1.P	64(R2), [V20.B16, V21.B16, V22.B16, V23.B16]
	VZIP1	V26.D2, V24.D2, V2.D2
	VZIP2	V26.D2, V24.D2, V6.D2
	VZIP1	V27.D2, V25.D2, V10.D2
	VZIP2	V27.D2, V25.D2, V14.D2
	VLD1.P	64(R2), [V24.B16, V25.B16, V26.B16, V27.B16]
	VZIP1	V30.D2, V28.D2, V3.D2
	VZIP2	V30.D2, V28.D2, V7.D2
	VZIP1	V31.D2, V29.D2, V11.D2
	VZIP2	V31.D2, V29.D2, V15.D2
	VLD1.P	64(R2), [V28.B16, V29.B16, V30.B16, V31.B16]
	VEOR	V0.B16, V16.B16, V16.B16
	VEOR	V1.B16, V17.B16, V17.B16
	VEOR	V2.B16, V18.B16, V18.B16
	VEOR	V3.B16, V19.B16, V19.B16
	VST1.P	[V16.B16, V17.B16, V18.B16, V19.B16], 64(R1)
	VEOR	V4.B16, V20.B16, V20.B16
	VEOR	V5.B16, V21.B16, V21.B16
	VEOR	V6.B16, V22.B16, V22.B16
	VEOR	V7.B16, V23.B16, V23.B16
	VST1.P	[V20.B16, V21.B16, V22.B16, V23.B16], 64(R1)
	VEOR	V8.B16, V24.B16, V24.B16
	VEOR	V9.B16, V25.B16, V25.B16
	VEOR	V10.B16, V26.B16, V26.B16
	VEOR	V11.B16, V27.B16, V27.B16
	VST1.P	[V24.B16, V25.B16, V26.B16, V27.B16], 64(R1)
	VEOR	V12.B16, V28.B16, V28.B16
	VEOR	V13.B16, V29.B16, V29.B16
	VEOR	V14.B16, V30.B16, V30.B16
	VEOR	V15.B16, V31.B16, V31.B16
	VST1.P	[V28.B16, V29.B16, V30.B16, V31.B16], 64(R1)

	ADD	$4, R20
	MOVW	R20, (R7) // update counter

	CMP	R2, R12
	BGT	loop

	RET


DATA	·constants+0x00(SB)/4, $0x61707865
DATA	·constants+0x04(SB)/4, $0x3320646e
DATA	·constants+0x08(SB)/4, $0x79622d32
DATA	·constants+0x0c(SB)/4, $0x6b206574
GLOBL	·constants(SB), NOPTR|RODATA, $32

DATA	·incRotMatrix+0x00(SB)/4, $0x00000000
DATA	·incRotMatrix+0x04(SB)/4, $0x00000001
DATA	·incRotMatrix+0x08(SB)/4, $0x00000002
DATA	·incRotMatrix+0x0c(SB)/4, $0x00000003
DATA	·incRotMatrix+0x10(SB)/4, $0x02010003
DATA	·incRotMatrix+0x14(SB)/4, $0x06050407
DATA	·incRotMatrix+0x18(SB)/4, $0x0A09080B
DATA	·incRotMatrix+0x1c(SB)/4, $0x0E0D0C0F
GLOBL	·incRotMatrix(SB), NOPTR|RODATA, $32
